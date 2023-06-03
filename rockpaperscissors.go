package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Welcome to Rock-Paper-Scissors!")
	fmt.Println("Enter your choice (rock, paper, scissors):")

	playerChoice := getPlayerChoice()
	computerChoice := getComputerChoice()

	fmt.Println("Player choice:", playerChoice)
	fmt.Println("Computer choice:", computerChoice)

	result := determineWinner(playerChoice, computerChoice)
	fmt.Println("Result:", result)
}

func getPlayerChoice() string {
	var choice string
	fmt.Scanln(&choice)
	return choice
}

func getComputerChoice() string {
	choices := []string{"rock", "paper", "scissors"}
	return choices[rand.Intn(len(choices))]
}

func determineWinner(playerChoice, computerChoice string) string {
	if playerChoice == computerChoice {
		return "It's a tie!"
	} else if playerChoice == "rock" && computerChoice == "scissors" ||
		playerChoice == "paper" && computerChoice == "rock" ||
		playerChoice == "scissors" && computerChoice == "paper" {
		return "Player wins!"
	} else {
		return "Computer wins!"
	}
}
