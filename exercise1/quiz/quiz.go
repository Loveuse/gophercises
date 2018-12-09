package quiz

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

type Loader interface {
	load(r io.Reader) error
}

type Store struct {
	Questions []string
	answers   []string
}

func NewStore(r io.Reader) (*Store, error) {
	store := &Store{Questions: []string{},
		answers: []string{},
	}
	if err := store.load(r); err != nil {
		return nil, err
	}
	return store, nil
}

// load reads from a CSV file with format question,answer and map those records to
// questions and answers of the Store
func (s *Store) load(r io.Reader) error {
	quizReader := csv.NewReader(r)
	s.Questions = []string{}
	s.answers = []string{}

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

		s.Questions = append(s.Questions, row[0])
		s.answers = append(s.answers, row[1])
	}

	if len(s.Questions) < 1 || len(s.answers) < 1 {
		return errors.New("could not read any question")
	}

	return nil
}

// Quiz struct that contains the questions, the private answers and the number
// of corrected answers
type Quiz struct {
	Store             *Store
	NumCorrectAnswers int
}

func New(store *Store) *Quiz {
	return &Quiz{
		Store: store,
	}
}

// StartQuiz scans the input answer for every question and updates
// the number of corrected answers
func (q *Quiz) StartQuiz(ctx context.Context, ch chan struct{}) error {
	var answer string
	for i, question := range q.Store.Questions {
		fmt.Print(question + ": ")
		if _, err := fmt.Scanf("%s\n", &answer); err != nil {
			return fmt.Errorf("could not read the answer for %s: ", question)
		}

		if answer == strings.TrimSpace(q.Store.answers[i]) {
			q.NumCorrectAnswers++
		}
	}
	ch <- struct{}{}
	return nil
}
