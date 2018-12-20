package quiz

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// Runner interface models an activity that can be executed:
// Run takes as input:
// - context: e.g. used for cancelling the activity
// - chan: signals when the activity ends
// - timer: time available to complete the activity
type Runner interface {
	Run(context.Context, chan struct{}, *time.Duration) error
}

// QA holds a question with its answer
type QA struct {
	Question string
	answer   string
}

// NewQA constructor for question answer holder
func NewQA(question, answer string) *QA {
	qa := &QA{
		Question: question,
		answer:   answer,
	}
	return qa
}

// CheckAnswer check whether the answer provided is correct
func (qa *QA) CheckAnswer(answer string) bool {
	return strings.TrimSpace(qa.answer) == strings.TrimSpace(answer)
}

// Quiz holds questions, answers and the number of corrected answers
type Quiz struct {
	QAs               []QA
	NumCorrectAnswers int
}

// New creates and returns a new quiz from questions and answers
func New(questionsAnswers []QA) *Quiz {
	quiz := &Quiz{
		QAs: questionsAnswers,
	}
	return quiz
}

// Run scans from standard input the answers for every question and updates
// the number of corrected answers
func (q *Quiz) Run(ctx context.Context, ch chan struct{}, timer *time.Duration) error {
	var answer string
	fmt.Println("Time available for the quiz: ", timer)
	for _, qa := range q.QAs {
		fmt.Print(qa.Question + ": ")
		if _, err := fmt.Scanf("%s\n", &answer); err != nil {
			return fmt.Errorf("could not read the answer for %s: ", qa.Question)
		}

		if qa.CheckAnswer(answer) {
			q.NumCorrectAnswers++
		}
	}
	ch <- struct{}{}
	return nil
}
