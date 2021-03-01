package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number")
	var n, sum, i int
	fmt.Scanln(&n)

	for i = 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}

	}

	fmt.Println(" The total form 1 to ", n, "is ", sum)
}
