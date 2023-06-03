package main

import (
	"fmt"
	"math/rand"
	"time"
)

// seed the random number generator with the current time
func main() {
	rand.Seed(time.Now().UnixNano())
	//generate a random integer between 1 and 100 and assign it to the variable secretNumber
	secretNumber := rand.Intn(100) + 1
	//create a variable named guess
	var guess int

	//loop infinitely
	for {
		//print the message
		fmt.Println("Guess the number (between 1 and 100): ")
		//take the user input for guess using the scan function
		fmt.Scan(&guess)

		//if guess is less than the secretNumber, print "Too low!"
		if guess < secretNumber {
			fmt.Println("Too low!")
			//otherwise, print too high!
		} else if guess > secretNumber {
			fmt.Println("Too high!")
		} else {
			//if the player guesses the correct number, print "you got it!"
			fmt.Println("You got it!")
			//break the loop
			break
		}
	}
}
