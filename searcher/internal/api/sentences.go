package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"smusmumbr.io/searcher/internal/search"
)

type Server struct {
	searchClient *search.SearchClient
}

func (srv *Server) SearchSentences(c *gin.Context) {
	term := c.Query("term")
	size := c.DefaultQuery("size", "10")
	if s, err := strconv.Atoi(size); err != nil {
		setStatus(c, http.StatusBadRequest, fmt.Sprintf("Invalid \"size\" query param: %s", size))
	} else if sentences, err := srv.search(term, s); err != nil {
		setStatus(c, http.StatusInternalServerError, fmt.Sprintf("Unknown error: %s", err))
	} else {
		c.JSON(http.StatusOK, sentences)
	}
}

func (s *Server) search(term string, size int) ([]*search.Sentence, error) {
	sentences := make([]*search.Sentence, 0)
	sources, err := s.searchClient.SearchWord("sentences", term, size)
	if err != nil {
		return nil, err
	}
	for _, source := range sources {
		sentences = append(
			sentences,
			&search.Sentence{
				SentenceMeta: search.SentenceMeta{
					Author: source["author"].(string),
					Work:   source["work"].(string),
				},
				Text: source["text"].(string),
			},
		)
	}
	return sentences, nil
}

func setStatus(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"message": message})
}

type uploadSentParams struct {
	search.SentenceMeta
	FileUrl string `json:"file_url" binding:"required"`
}

func (srv *Server) AddSentences(c *gin.Context) {
	var uparams uploadSentParams
	if err := c.ShouldBindJSON(&uparams); err != nil {
		setStatus(c, http.StatusBadRequest, fmt.Sprintf("binding error: %s", err))
		return
	}
	resp, err := http.Get(uparams.FileUrl)
	if err != nil {
		setStatus(c, http.StatusBadRequest, err.Error())
	}

	tmpFile, err := os.CreateTemp(os.TempDir(), "sentencer")
	if err != nil {
		setStatus(
			c,
			http.StatusInternalServerError,
			fmt.Sprintf("Error while creating temp file: %s", err),
		)
		return
	}
	defer os.Remove(tmpFile.Name())

	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		setStatus(
			c,
			http.StatusInternalServerError,
			fmt.Sprintf("Error while saving file: %s", err),
		)
		return
	}
	reader, err := search.GetSentenceReader(tmpFile.Name(), uparams.Work, uparams.Author, 2)
	if err != nil {
		setStatus(c, http.StatusInternalServerError, fmt.Sprintf("CSV reader error: %s", err))
	}
	for {
		r, err := reader()
		if len(r) > 0 {
			if err := srv.searchClient.BulkCreate("sentences", r); err != nil {
				setStatus(
					c,
					http.StatusInternalServerError,
					fmt.Sprintf("OS saving error: %s", err),
				)
			}
		}
		if err == io.EOF {
			break
		}
	}
	setStatus(c, http.StatusOK, "file uploaded successfully")
}

func NewServer() *Server {
	return &Server{searchClient: search.NewSearchClient()}
}
