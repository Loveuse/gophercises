package store

import (
	"io"
	"strings"
)

// Store interface that models the operation on a store:
// - Extract: retrieve QAs from a reader
// - Setup: loads the QAs in the store
// - GetQAs: returns QAs
// - LenQAs: returns the number of QAs
type Store interface {
	Extract(r io.Reader) ([]QA, error)
	Setup([]QA)
	GetQAs() []QA
	QAsLen() int
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
