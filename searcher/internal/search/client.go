package search

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"

	"smusmumbr.io/searcher/internal/config"
)

type osIndex struct {
	Name string `json:"_index"`
	Id   int    `json:"_id,omitempty"`
}

type osIndexOperation struct {
	Index *osIndex `json:"index"`
}

type osTextQuery struct {
	Query string `json:"query"`
}
type osMatch struct {
	Text *osTextQuery `json:"text"`
}
type osQuery struct {
	Match *osMatch `json:"match"`
}
type osSearchBody struct {
	Size  int      `json:"size"`
	Query *osQuery `json:"query"`
}

type SearchClient struct {
	osClient *opensearch.Client
}

type osSearchResponse struct {
	Hits *osSearchResponseHits `json:"hits"`
}
type osSearchResponseHits struct {
	Hits []*osSearchSource `json:"hits"`
}
type osSearchSource struct {
	Source map[string]any `json:"_source"`
}

func (sc *SearchClient) init() {
	var err error
	if sc.osClient, err = opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{config.OpensearchURL()},
	}); err != nil {
		log.Fatal(err)
	}
}

func (sc *SearchClient) CreateIndex(name string) error {
	settings := strings.NewReader(`{
	  'settings': {
	      'index': {
	          'number_of_shards': 1,
	          'number_of_replicas': 0
	          }
	      }
	  }`)
	if _, err := sc.osClient.Indices.Create("sentences"); err != nil {
		return err
	}
	_, err := sc.osClient.Indices.PutSettings(settings)
	return err
}

func appendDoc(b *strings.Builder, doc any) error {
	if mDoc, err := json.Marshal(doc); err != nil {
		return err
	} else {
		b.Write(mDoc)
		b.WriteRune('\n')
	}
	return nil
}

func (sc *SearchClient) BulkCreate(index string, documents []any) error {
	b := strings.Builder{}
	for _, doc := range documents {
		if err := appendDoc(&b, &osIndexOperation{&osIndex{Name: index}}); err != nil {
			return err
		}
		if err := appendDoc(&b, &doc); err != nil {
			return err
		}
	}
	_, err := sc.osClient.Bulk(strings.NewReader(b.String()))
	return err
}

func (sc *SearchClient) SearchWord(index, query string, size int) ([]map[string]any, error) {
	doc := osSearchBody{
		Size: size,
		Query: &osQuery{
			&osMatch{
				&osTextQuery{Query: query},
			},
		},
	}
	mDoc, err := json.Marshal(doc)
	if err != nil {
		return nil, err
	}
	content := strings.NewReader(string(mDoc))
	search := opensearchapi.SearchRequest{Index: []string{index}, Body: content}
	rsp, err := search.Do(context.Background(), sc.osClient)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(rsp.Body); err != nil {
		return nil, err
	}

	var hitsRsp osSearchResponse
	if err := json.Unmarshal(buf.Bytes(), &hitsRsp); err != nil {
		return nil, err
	}
	var res []map[string]any
	for _, hit := range hitsRsp.Hits.Hits {
		res = append(res, hit.Source)
	}
	return res, nil
}

func NewSearchClient() *SearchClient {
	sc := SearchClient{}
	sc.init()
	return &sc
}
