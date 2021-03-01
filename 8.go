package main

import (
	"fmt"
	"os"
)

var n, m, input int

func main() {
	var name string
	fmt.Println("Hello! Please enter your name")
	fmt.Scanln(&name)
	fmt.Println("1.Print tables range of table \n2.Table of your choice\n ")
	fmt.Scanln(&input)
	if input > 2 {
		fmt.Println("Wrong input", name, "!\n ")
		os.Exit(input)
	}
	fmt.Println("Enter the Multiplicant\n ")
	fmt.Scanln(&n)
	fmt.Println("Enter the range of multiplier\n ")
	fmt.Scanln(&m)
	if input == 2 {
		input2()
	} else if input == 1 {
		input1()

	}
}
func input1() {
	fmt.Println("The table from 1 to", n, "is \n ")
	for n := n; n >= 1; n-- {
		fmt.Println("----------------------- \n ")
		fmt.Println("The table of", n, "is\t \n ")
		for i := 1; i <= m; i++ {
			fmt.Println(n, " * ", i, " = ", n*i, "\t\n ")
		}
	}

}
func input2() {
	fmt.Println("The Table of", n, "is \n ")
	for i := 1; i <= m; i++ {
		fmt.Println(n, " * ", i, " = ", n*i, "\t\n ")
	}

}
