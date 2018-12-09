package main

import (
	"context"
	"fmt"
	"strings"
)

// Quiz struct that contains the questions, the private answers and the number
// of corrected answers
type Quiz struct {
	Questions         []string
	answers           []string
	NumCorrectAnswers int
}

// StartQuiz scans the input answer for the relative question and update
// the number of corrected answers
func (q *Quiz) StartQuiz(ctx context.Context, ch chan struct{}) error {
	var answer string
	for i, question := range q.Questions {
		ctx.Done()
		fmt.Print(question + ": ")
		fmt.Scanf("%s\n", &answer)
		if answer == strings.TrimSpace(q.answers[i]) {
			q.NumCorrectAnswers++
		}
	}
	ch <- struct{}{}
	return nil
}
