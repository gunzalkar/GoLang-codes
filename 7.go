package main

import (
	"fmt"
	"os"
)

func main() {
	var n, m, input int
	var name string
	fmt.Println("Hello! Please enter your name")
	fmt.Scanln(&name)
	fmt.Println("1.Print tables from 1 to your choice\n2.Table of your choice")
	fmt.Scanln(&input)
	if input > 2 {
		fmt.Println("Wrong input", name, "!\n ")
		os.Exit(input)
	}
	fmt.Println("Enter the Multiplicant")
	fmt.Scanln(&n)
	fmt.Println("Enter the range of multiplier")
	fmt.Scanln(&m)
	if input == 2 {
		fmt.Println("The Table of", n, "is\n ")
		for i := 1; i <= m; i++ {
			fmt.Println(n, " * ", i, " = ", n*i, "\t\n ")
		}
	} else if input == 1 {
		fmt.Println("The table from 1 to", n, "is \n ")
		for n := n; n >= 1; n-- {
			fmt.Println("-----------------------\n ")
			fmt.Println("The table of", n, "is\t")
			for i := 1; i <= m; i++ {
				fmt.Println(n, " * ", i, " = ", n*i, "\t\n ")
			}
		}
	}
}
