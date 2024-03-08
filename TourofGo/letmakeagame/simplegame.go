package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	answer string
)

func start() {

	for {
		fmt.Println("Do you wanna play this game?")
		fmt.Scan(&answer)
		if answer == "yes" || answer == "y" {
			fmt.Println("Let's start this game")
			break
		} else {
			fmt.Println("Are you sure?, pls answer again")
			fmt.Scan(&answer)
			if answer == "no" || answer == "n" {
				break
			} else {
				fmt.Println("Let's start this game")
				break
			}
		}
	}
}

func instruct() {
	fmt.Println("Welcome, I appriciate to tell you rule in this game ")
	time.Sleep(time.Second)
	fmt.Println("You'll get 5 attempt to make correct answer")
	time.Sleep(time.Second)
	fmt.Println("It's very simply , just guess correct number")
	time.Sleep(time.Second)
	fmt.Println("Result number 'll create in 1 second, Please stand by")
	time.Sleep(time.Second)
	Cnum = rand.Intn(80)

}

var (
	Cnum    int
	Gnum    int
	Attempt int = 5
)

func logic() {

	for {
		if Attempt <= 0 {
			fmt.Print("\033[2J")
			fmt.Print("\033[H")
			fmt.Println("You lose the pancake from Rem. Correct Answer is: ", Cnum)
			break
		} else {
			fmt.Println("Please guess your number: ")
			fmt.Scan(&Gnum)

			if Cnum > Gnum {
				Attempt--
				fmt.Println("Your guess number lower than Correct Number!")
				fmt.Println("Your attempt: ", Attempt)
				continue
			} else if Cnum < Gnum {
				Attempt--
				fmt.Println("Your guess number higher than Correct Number! ")
				fmt.Println("Your attempt: ", Attempt)
				continue
			} else {
				fmt.Println("Your guess correct number ! Congratulations")
			}
		}
	}

}

func main() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	start()
	instruct()
	logic()
}
