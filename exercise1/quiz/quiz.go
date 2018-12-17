package quiz

import (
	"context"
	"fmt"
	"io"
	"strings"
)

// Extractor interface that models teh actvity to extract/retrieve data from a reader.
type Extractor interface {
	Extract(r io.Reader) (Questions, Answers, error)
}

// Runnable interface models an activity that can be executed:
// Run takes as input:
// - context: e.g. used for cancelling the activity
// - chan: signals when the activity ends
type Runnable interface {
	Run(context.Context, chan struct{}) error
}

// Questions defines questions as slice of strings
type Questions []string

// Answers defines answers as slice of strings
type Answers []string

// Quiz holds questions, answers and the number of corrected answers
type Quiz struct {
	Questions         Questions
	Answers           Answers
	NumCorrectAnswers int
}

// New creates and returns a new quiz from questions and answers
func New(questions Questions, answers Answers) *Quiz {
	quiz := &Quiz{
		Questions: questions,
		Answers:   answers,
	}
	return quiz
}

// Run scans from standard input the answers for every question and updates
// the number of corrected answers
func (q *Quiz) Run(ctx context.Context, ch chan struct{}) error {
	var answer string
	for i, question := range q.Questions {
		fmt.Print(question + ": ")
		if _, err := fmt.Scanf("%s\n", &answer); err != nil {
			return fmt.Errorf("could not read the answer for %s: ", question)
		}

		if answer == strings.TrimSpace(q.Answers[i]) {
			q.NumCorrectAnswers++
		}
	}
	ch <- struct{}{}
	return nil
}
