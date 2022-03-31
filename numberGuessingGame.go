package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var input string

func main() {
	//This code creates random number and place as variable
	rand.Seed(time.Now().UnixNano())
	hiddenNumber := int64(rand.Intn(100))

	for tries := 1; tries <= 5; tries++ {

		fmt.Println("Input guessing number:")
		fmt.Scanln(&input)

		userInput, errorTest := strconv.ParseInt(input, 10, 0)

		if errorTest != nil {
			fmt.Println("Please input a numerical value")
			continue
		}

		if userInput == 101 {
			fmt.Println("You have chosen to end the game")
			break
		} else if userInput > 100 || userInput < 0 {
			fmt.Println("You have entered an incorrect number input:")
			fmt.Println("Pleae enter number within 0-100")
		} else if userInput == hiddenNumber {
			fmt.Println("You have correctly guessed the number:")
			fmt.Println("Game is over")
			break
		} else if userInput > hiddenNumber {
			fmt.Println("The number is too high")
		} else {
			fmt.Println("The number is too low")
		}

		if tries == 5 {
			fmt.Println("Game over, you have failed")
		}

	}

}
