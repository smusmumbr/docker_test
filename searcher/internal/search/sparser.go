package search

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type SentenceMeta struct {
	Work   string `json:"work"   binding:"required"`
	Author string `json:"author" binding:"required"`
}

type Sentence struct {
	SentenceMeta
	Text string `json:"text"`
}

func (s *Sentence) String() string {
	return fmt.Sprintf("[Work: %s, Author: %s, Text: %s]", s.Work, s.Author, s.Text)
}

func GetSentenceReader(
	filepath string,
	work string,
	author string,
	chunkSize uint8,
) (func() ([]any, error), error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	return func() ([]any, error) {
		res := make([]any, 0, chunkSize)
		var err error
		for i := 0; i < int(chunkSize); i++ {
			if r, err := reader.Read(); err == io.EOF {
				return res, io.EOF
			} else if err != nil {
				return nil, err
			} else {
				res = append(res, &Sentence{SentenceMeta: SentenceMeta{Author: author, Work: work}, Text: r[0]})
			}
		}
		return res, err
	}, nil
}
