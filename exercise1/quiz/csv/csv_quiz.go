package fromcsv

import (
	"encoding/csv"
	"errors"
	"io"
	"log"

	"github.com/loveuse/gophercises/exercise1/quiz"
)

// Extractor is a placeholder type that will implement the interface Extractor
type Extractor struct{}

// New returns a new empty extractor
func New(r io.Reader) *Extractor {
	return &Extractor{}
}

// Extract reads from a CSV file with format question,answer and map those records to
// questions and answers of the Store
func (e *Extractor) Extract(r io.Reader) (quiz.Questions, quiz.Answers, error) {
	quizReader := csv.NewReader(r)

	questions := quiz.Questions{}
	answers := quiz.Answers{}

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

		questions = append(questions, row[0])
		answers = append(answers, row[1])
	}

	if len(questions) < 1 || len(answers) < 1 {
		return nil, nil, errors.New("could not read any question")
	}

	return questions, answers, nil
}
