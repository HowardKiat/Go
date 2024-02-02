package main

import "fmt"

func main() {
	fmt.Println("Welcome To The Quiz")
	fmt.Println("-----------------------")

	fmt.Printf("Enter Your Name: ")

	var name string
	fmt.Scan(&name)
	// Auto determine the type same thing as var
	// name := "Joe"
	fmt.Printf("Wassup, %v, Welcome to the game!", name)
}
