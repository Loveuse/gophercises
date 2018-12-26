package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Loveuse/gophercises/exercise1/quiz"
	"github.com/Loveuse/gophercises/exercise1/quiz/extractor/csv"
)

const (
	csvPath      = "quiz/extractor/csv/problems.csv"
	defaultTimer = 30 * time.Second
)

func main() {
	quizPath := flag.String("csv", csvPath, "quiz CSV file with format: question,answer")
	timer := flag.Duration("t", defaultTimer, "time to complete the quiz, format: 1s, 1m, 1h")
	flag.Parse()

	filePtr, err := os.Open(*quizPath)
	if err != nil {
		log.Fatalf("could not open the quiz CSV file %s:", *quizPath)
	}
	defer filePtr.Close()

	store := &csv.Store{}
	QAs, err := store.Extract(filePtr)
	if err != nil {
		log.Fatalf("could not load the questions/answer file: %v", err)
	}
	store.Setup(QAs)

	quiz := quiz.New(store)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, *timer)
	defer cancel()

	quizNotifyChan := make(chan struct{})
	go func() {
		err = quiz.Run(ctx, quizNotifyChan, timer)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("\nTime Expired.")
	case <-quizNotifyChan:
		// user has finished the quiz
	}

	fmt.Printf("Number of questions %d. Correct answers: %d\n", quiz.Store.QAsLen(), quiz.NumCorrectAnswers)
}
