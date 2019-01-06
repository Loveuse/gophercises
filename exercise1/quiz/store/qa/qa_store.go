package qastore

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrNoData     = errors.New("no questions and answers to retrieve from the store")
	ErrRetrieveQA = errors.New("could not retrieve questions and answers")
)

// QA holds a question with his relative answer
type QA struct {
	Question string
	answer   string
}

// NewQA constructor for question and answer holder
func NewQA(question, answer string) *QA {
	qa := &QA{
		Question: question,
		answer:   answer,
	}
	return qa
}

// MarshalJSON implement json.Unmarshaller
func (qa *QA) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Question string
		Answer   string
	}{
		Question: qa.Question,
		Answer:   qa.answer,
	})
}

// UnmarshalJSON implement json.Unmarshaller
func (qa *QA) UnmarshalJSON(b []byte) error {
	temp := struct {
		Question string
		Answer   string
	}{}
	if err := json.Unmarshal(b, &temp); err != nil {
		return errors.Errorf("could not retrieve the QA: %v", err)
	}
	qa.Question = temp.Question
	qa.answer = temp.Answer
	return nil
}

// String omits the answer from the QA
func (qa QA) String() string {
	return qa.Question
}

// Format should be implemented for fmt.Printf("%#v", qas)

// CheckAnswer check whether the answer provided is correct
func (qa *QA) CheckAnswer(answer string) bool {
	return strings.TrimSpace(qa.answer) == strings.TrimSpace(answer)
}

// QAStore holds the questions and answers and implements store.Store
type QAStore struct {
	QAs []QA
}

// New return a new QAStore
func New() *QAStore {
	store := &QAStore{}
	return store
}

// Get returns the []byte representation of the questions and answers
func (qas *QAStore) Get() ([]byte, error) {
	if len(qas.QAs) < 1 {
		return nil, ErrNoData
	}

	b, err := json.Marshal(qas.QAs)
	if err != nil {
		return nil, ErrRetrieveQA
	}
	return b, nil
}

// Set retrieves the []byte representation of []QA and save it in the store
func (qas *QAStore) Set(b []byte) error {
	if err := json.Unmarshal(b, &qas.QAs); err != nil {
		return errors.Errorf("could not save data in the store: %v", err)
	}
	return nil
}

// LenQAs returns the number of QAs
func (qas *QAStore) LenQAs() (int, error) {
	if len(qas.QAs) < 1 {
		return 0, ErrNoData
	}
	return len(qas.QAs), nil
}
