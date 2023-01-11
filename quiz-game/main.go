package main

import(
	"fmt"
)

func main(){
	fmt.Println("Welcome to the quiz game!")

	var name string
	var age uint // uint so no negative values

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Hi %v, welcome to the game! \n", name)

	fmt.Printf("%v, how old are you? ", name)

	fmt.Scan(&age)

	if age < 15 {
		fmt.Println("Sorry, you are too young to play our game")
		return
	}
		fmt.Println("Test")
}
