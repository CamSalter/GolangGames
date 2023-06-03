package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	score := 0
	totalQuestions := 5
	timeLimit := 10 // Time limit in seconds

	fmt.Println("Welcome to the Math Quiz Game!")
	fmt.Printf("You have %d seconds to answer %d math questions.\n", timeLimit, totalQuestions)
	fmt.Println("Let's begin!")

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	quit := make(chan bool)

	go func() {
		for i := 0; i < totalQuestions; i++ {
			a, b := generateRandomNumbers()
			operator := generateRandomOperator()

			question := generateQuestion(a, b, operator)
			answer := calculateAnswer(a, b, operator)

			fmt.Printf("Question %d: %s\n", i+1, question)

			var userAnswer int
			_, err := fmt.Scanln(&userAnswer)
			if err != nil {
				fmt.Println("Invalid input. Skipping question.")
				continue
			}

			if userAnswer == answer {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Println("Incorrect!")
			}

			fmt.Println()
		}
		quit <- true
	}()

	select {
	case <-timer.C:
		fmt.Println("Time's up!")
	case <-quit:
		fmt.Println("Quiz completed!")
	}

	fmt.Printf("Score: %d/%d\n", score, totalQuestions)
}
