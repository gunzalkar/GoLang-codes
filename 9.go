package main

import (
	"fmt"
	"os"
)

var n, input int
var name string

func square() {
	fmt.Println("Square of", n, "is", n*n)
}
func squarelist() {
	fmt.Println("The list of Squared numbers are \n ")
	for n := n; n >= 1; n-- {
		fmt.Println("Square of", n, "is", n*n)
	}
}
func main() {
	fmt.Println("Enter your name")
	fmt.Scanln(&name)
	fmt.Println("Enter the number you want to square")
	fmt.Scanln(&n)
	fmt.Println("1.Square a Number\n 2.Square numbers till", n)
	fmt.Scanln(&input)
	if input == 1 {
		square()
	}
	if input == 2 {
		squarelist()
	}
	if input > 2 {
		fmt.Println("Wrong input", name, "!\n ")
		os.Exit(input)
	}
}
