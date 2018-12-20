package csv

import (
	"encoding/csv"
	"errors"
	"io"
	"log"

	"github.com/Loveuse/gophercises/exercise1/quiz"
)

// Extractor is a placeholder type that implements the interface Extractor
type Extractor struct{}

// New returns a new empty extractor
func New(r io.Reader) *Extractor {
	return &Extractor{}
}

// Extract reads from a CSV file with format "question,answer" and
// returns questions and answers
func (e *Extractor) Extract(r io.Reader) ([]quiz.QA, error) {
	quizReader := csv.NewReader(r)
	QAs := []quiz.QA{}

	for {
		row, err := quizReader.Read()
		if err == io.EOF {
			break
		}
		// just logging a row not readable
		if err != nil {
			log.Printf("could not read a row: %v", err)
			continue
		}
		QAs = append(QAs, *quiz.NewQA(row[0], row[1]))
	}

	if len(QAs) < 1 {
		return nil, errors.New("could not read any question")
	}

	return QAs, nil
}
