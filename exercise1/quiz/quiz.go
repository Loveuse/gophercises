package quiz

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"

	store "github.com/Loveuse/gophercises/exercise1/quiz/store"
	qa "github.com/Loveuse/gophercises/exercise1/quiz/store/qa"
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
// the number of corrected answers after retrieving the QAs from the store
func (q *Quiz) Run(ctx context.Context, ch chan struct{}, timer *time.Duration) error {
	var answer string
	fmt.Println("Time available for the quiz: ", timer)

	qas, err := q.getQAs()
	if err != nil {
		return errors.Errorf("could not retrieve the QAs from the store: %v", err)
	}

	for _, qa := range qas {
		fmt.Print(qa.Question + ": ")
		if _, err := fmt.Scanf("%s\n", &answer); err != nil {
			return errors.Errorf("could not read the answer for %s: ", qa.Question)
		}

		if qa.CheckAnswer(answer) {
			q.NumCorrectAnswers++
		}
	}
	ch <- struct{}{}
	return nil
}

func (q *Quiz) getQAs() ([]qa.QA, error) {
	byteQAs, err := q.Store.Get()
	if err != nil {
		return nil, errors.Errorf("could not retrieve the QAs to run the quiz: %v", err)
	}
	qas := []qa.QA{}
	if err := json.Unmarshal(byteQAs, &qas); err != nil {
		return nil, errors.Errorf("could not retrieve the QAs from the store: %v", err)
	}
	return qas, nil
}
