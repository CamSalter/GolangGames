package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Adventure Game!")
	fmt.Println("You find yourself standing at the entrance of a mysterious cave.")
	fmt.Println("Do you want to enter the cave? (yes/no)")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice == "yes" {
		fmt.Println("You enter the dark cave and start exploring.")

		fmt.Println("You come across two tunnels. Do you want to go left or right? (left/right)")
		choice, _ = reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "left" {
			fmt.Println("You choose the left tunnel and continue walking.")

			fmt.Println("You encounter a fierce dragon blocking your way!")
			fmt.Println("Do you want to fight the dragon or run away? (fight/run)")
			choice, _ = reader.ReadString('\n')
			choice = strings.TrimSpace(choice)

			if choice == "fight" {
				fmt.Println("You bravely engage in a fight with the dragon.")
				fmt.Println("After a fierce battle, you manage to defeat the dragon!")
				fmt.Println("Congratulations! You have completed the adventure.")
			} else if choice == "run" {
				fmt.Println("You decide not to risk it and run away from the dragon.")
				fmt.Println("You find another path and continue your adventure.")
				fmt.Println("Eventually, you reach the end of the cave and emerge victorious!")
				fmt.Println("Congratulations! You have completed the adventure.")
			} else {
				fmt.Println("Invalid choice. The adventure ends here.")
			}
		} else if choice == "right" {
			fmt.Println("You choose the right tunnel and continue walking.")

			fmt.Println("As you proceed further, you encounter a deep chasm.")
			fmt.Println("Do you want to jump across the chasm or turn back? (jump/turn back)")
			choice, _ = reader.ReadString('\n')
			choice = strings.TrimSpace(choice)

			if choice == "jump" {
				fmt.Println("You take a leap of faith and jump across the chasm!")
				fmt.Println("You successfully land on the other side and continue your adventure.")
				fmt.Println("After facing numerous challenges, you finally emerge victorious!")
				fmt.Println("Congratulations! You have completed the adventure.")
			} else if choice == "turn back" {
				fmt.Println("You decide not to risk it and turn back.")
				fmt.Println("You exit the cave without completing the adventure.")
				fmt.Println("Better luck next time!")
			} else {
				fmt.Println("Invalid choice. The adventure ends here.")
			}
		} else {
			fmt.Println("Invalid choice. The adventure ends here.")
		}
	} else if choice == "no" {
		fmt.Println("You decide not to enter the cave. The adventure ends here.")
	} else {
		fmt.Println("Invalid choice. The adventure ends here.")
	}
}
