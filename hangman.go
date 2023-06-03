package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	word := "hangman" // The word to guess
	guesses := make([]string, 0)
	maxAttempts := 6
	attempts := 0

	fmt.Println("Welcome to Hangman!")
	fmt.Printf("The word has %d letters.\n", len(word))

	for {
		printWord(word, guesses)
		fmt.Println()

		if attempts >= maxAttempts {
			fmt.Println("You have reached the maximum number of attempts. Game over!")
			fmt.Printf("The word was: %s\n", word)
			break
		}

		fmt.Printf("Attempts remaining: %d\n", maxAttempts-attempts)
		fmt.Print("Enter a letter: ")

		reader := bufio.NewReader(os.Stdin)
		letter, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		letter = strings.TrimSpace(letter)

		if len(letter) != 1 {
			fmt.Println("Please enter a single letter.")
			continue
		}

		if contains(guesses, letter) {
			fmt.Printf("You have already guessed the letter '%s'. Please try again.\n", letter)
			continue
		}

		guesses = append(guesses, letter)

		if strings.Contains(word, letter) {
			fmt.Println("Correct guess!")
		} else {
			fmt.Println("Wrong guess!")
			attempts++
		}

		if isWordGuessed(word, guesses) {
			fmt.Println("Congratulations! You guessed the word correctly!")
			break
		}
	}
}

func printWord(word string, guesses []string) {
	for _, letter := range word {
		if contains(guesses, string(letter)) {
			fmt.Printf("%c ", letter)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

func contains(arr []string, val string) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}

func isWordGuessed(word string, guesses []string) bool {
	for _, letter := range word {
		if !contains(guesses, string(letter)) {
			return false
		}
	}
	return true
}
