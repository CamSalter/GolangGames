package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	words := []string{"apple", "banana", "cherry", "orange", "mango"} // List of words for the game
	rand.Seed(time.Now().UnixNano())
	word := words[rand.Intn(len(words))] // Randomly select a word from the list

	scrambledWord := scrambleWord(word)

	fmt.Println("Welcome to Word Scramble!")
	fmt.Println("Unscramble the letters to form a word.")
	fmt.Println("Scrambled word:", scrambledWord)

	var guess string
	fmt.Print("Enter your guess: ")
	fmt.Scanln(&guess)

	if strings.ToLower(guess) == word {
		fmt.Println("Congratulations! You guessed the word correctly!")
	} else {
		fmt.Println("Oops! Your guess is incorrect.")
		fmt.Println("The correct word is:", word)
	}
}

func scrambleWord(word string) string {
	runes := []rune(word)
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}
