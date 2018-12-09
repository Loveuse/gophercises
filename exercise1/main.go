package main

import (
	"context"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	quizPath := flag.String("csv", "problems.csv", "quiz CSV file with format: question,answer")
	timer := flag.Duration("t", 3*time.Second, "timer for complete the quiz, format: 1s, 1m, 1h")
	flag.Parse()

	filePtr, err := os.Open(*quizPath)
	if err != nil {
		log.Fatalf("could not open the quiz CSV file %s:", *quizPath)
	}
	defer filePtr.Close()

	questions, answers, err := loadQuiz(filePtr)

	quiz := &Quiz{Questions: questions, answers: answers}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, *timer)
	defer cancel()

	ch := make(chan struct{})
	go func() {
		err = quiz.StartQuiz(ctx, ch)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("\nTime Expired.")
	case <-ch:
	}
	fmt.Printf("Number of questions %d. Correct answers: %d\n", len(quiz.Questions), quiz.NumCorrectAnswers)
}

// loadQuiz reads from the CSV and returns the questions and answers slices
func loadQuiz(filePtr *os.File) ([]string, []string, error) {
	quizReader := csv.NewReader(filePtr)
	questions := []string{}
	answers := []string{}

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
