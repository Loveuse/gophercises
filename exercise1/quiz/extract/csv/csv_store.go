package csv

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"

	store "github.com/Loveuse/gophercises/exercise1/quiz/store/qa"
	"github.com/pkg/errors"
)

// Extractor implements extract.
type Extractor struct{}

// Extract reads from a CSV file with format "question,answer" and
// returns questions and answers
func (e *Extractor) Extract(r io.Reader) ([]byte, error) {
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

	b, err := json.Marshal(QAs)
	if err != nil {
		return nil, errors.Errorf("could not return QAs retrieved: %v", err)
	}

	return b, nil
}
