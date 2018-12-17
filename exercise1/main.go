package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/loveuse/gophercises/exercise1/quiz"
	"github.com/loveuse/gophercises/exercise1/quiz/csv"
)

func main() {
	quizPath := flag.String("csv", "quiz/csv/problems.csv", "quiz CSV file with format: question,answer")
	timer := flag.Duration("t", 30*time.Second, "timer for complete the quiz, format: 1s, 1m, 1h")
	flag.Parse()

	filePtr, err := os.Open(*quizPath)
	if err != nil {
		log.Fatalf("could not open the quiz CSV file %s:", *quizPath)
	}
	defer filePtr.Close()

	qaExtractor := fromcsv.Extractor{}
	questions, answers, err := qaExtractor.Extract(filePtr)
	if err != nil {
		log.Fatalf("could not load the questions/answer file: %v", err)
	}

	quiz := quiz.New(questions, answers)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, *timer)
	defer cancel()

	quizNotifyChan := make(chan struct{})
	go func() {
		err = quiz.Run(ctx, quizNotifyChan)
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

	fmt.Printf("Number of questions %d. Correct answers: %d\n", len(quiz.Questions), quiz.NumCorrectAnswers)
}
