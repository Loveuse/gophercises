package quiz

import (
	"context"
	"fmt"
	"time"

	store "github.com/Loveuse/gophercises/exercise1/quiz/extractor"
)

// Runner interface models an activity that can be executed:
// Run takes as input:
// - context: e.g. used for cancelling the activity
// - chan: signals when the activity ends
// - timer: time available to complete the activity
type Runner interface {
	Run(context.Context, chan struct{}, *time.Duration) error
}

// Quiz holds questions, answers and the number of corrected answers
type Quiz struct {
	Store             store.Store
	NumCorrectAnswers int
}

// New creates a new quiz with store as dependency
func New(store store.Store) *Quiz {
	quiz := &Quiz{
		Store: store,
	}
	return quiz
}

// Run scans from standard input the answers for every question and updates
// the number of corrected answers
func (q *Quiz) Run(ctx context.Context, ch chan struct{}, timer *time.Duration) error {
	var answer string
	fmt.Println("Time available for the quiz: ", timer)
	for _, qa := range q.Store.GetQAs() {
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
