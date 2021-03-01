package main

import (
	"fmt"
)

func main() {
	fact := 1
	var sum, i, input, n int
	fmt.Println("Enter your name")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Enter a number")
	fmt.Scanln(&n)
	fmt.Println("Enter 1 for sum enter 2 for product")
	fmt.Scanln(&input)

	if input == 1 {
		for i = 1; i <= n; i++ {
			{
				sum += i
			}
		}
		fmt.Println("The sum from 1 to", n, "is", sum)
	}

	if input == 2 {
		for i = 1; i <= n; i++ {
			fact = fact * i
		}
		fmt.Println("The factorial of", n, "is", fact)
	}
	if input != 2 && input != 1 {
		fmt.Println("Wrong input", name)
	}
}
