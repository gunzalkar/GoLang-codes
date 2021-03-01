package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter you first name")
	var first string
	fmt.Scanln(&first)
	fmt.Println("Enter your last name")
	var last string
	fmt.Scanln(&last)
	fmt.Println("Your name is " + first + " " + last)
}
