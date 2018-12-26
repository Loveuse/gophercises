package csv

import (
	"encoding/csv"
	"errors"
	"io"
	"log"

	store "github.com/Loveuse/gophercises/exercise1/quiz/extractor"
)

// Store holds the questions and answers and implements extract.Store
type Store struct {
	QAs []store.QA
}

// Extract reads from a CSV file with format "question,answer" and
// returns questions and answers
func (s *Store) Extract(r io.Reader) ([]store.QA, error) {
	quizReader := csv.NewReader(r)
	QAs := []store.QA{}

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
		QAs = append(QAs, *store.NewQA(row[0], row[1]))
	}

	if len(QAs) < 1 {
		return nil, errors.New("could not read any question")
	}

	return QAs, nil
}

// Setup assigns the questions and answers for the store
func (s *Store) Setup(qas []store.QA) {
	s.QAs = qas
}

// GetQAs returns the questions and answers held by the store
func (s *Store) GetQAs() []store.QA {
	return s.QAs
}

// QAsLen returns the number of questions and answers
func (s *Store) QAsLen() int {
	return len(s.QAs)
}
