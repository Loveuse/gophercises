package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Loveuse/gophercises/exercise1/quiz"
	csv "github.com/Loveuse/gophercises/exercise1/quiz/extract/csv"
	store "github.com/Loveuse/gophercises/exercise1/quiz/store/qa"
)

const (
	csvPath      = "quiz/extract/csv/problems.csv"
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

	// extract data from a CSV file
	extractor := csv.Extractor{}
	byteQAs, err := extractor.Extract(filePtr)
	if err != nil {
		log.Fatalf("could not load the questions/answer file: %v", err)
	}

	// create a store to save the QAs
	qastore := store.New()
	if err := qastore.Set(byteQAs); err != nil {
		log.Fatalf("could not save QAs into the store: %v", err)
	}

	quiz := quiz.New(qastore)

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

	lenQAs, err := qastore.LenQAs()
	if err != nil {
		log.Fatalf("could not retrieve the number of QAs: %v", err)
	}

	fmt.Printf("Number of questions %d. Correct answers: %d\n", lenQAs, quiz.NumCorrectAnswers)
}
