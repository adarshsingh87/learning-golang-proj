package main

import(
	"fmt"
)

func main(){
	fmt.Println("Welcome to the quiz game!")

	var name string
	var age uint // uint so no negative values
	score :=0

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Hi %v, welcome to the game! \n", name)

	fmt.Printf("%v, how old are you? ", name)

	fmt.Scan(&age)

	if age < 15 {
		fmt.Println("Sorry, you are too young to play our game")
		return
	}

	var cores int
	fmt.Print("How many cores does Ryzen 9 3900x have: ")
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct! ")
		score++
	} else {
		fmt.Printf("Incorrect, score %v", score)
		return
	}

	var fastestGpu1 string
	var fastestGpu2 string

	fmt.Print("What is the fastest 40Series Nvidia GPU: ")
	fmt.Scan(&fastestGpu1, &fastestGpu2)

	if fastestGpu1 + " " + fastestGpu2 == "RTX 4090" {
		fmt.Println("Correct! ")
		score++
	} else if fastestGpu1 + " " + fastestGpu2 == "rtx 4090" {
		fmt.Println("Correct! ")
		score++
	}	else {
		fmt.Printf("Incorrect, score %v", score)
		return
	}

	fmt.Printf("score: %v", score)

}
