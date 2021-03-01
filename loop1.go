package main

import "fmt"

func main() {

	var name string
	var age int
	 {

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	break 
}

	fmt.Print("Enter your age: ")
	fmt.Scanf("%d", &age)
	fmt.Printf("%s is %d years old\n", name, age)
}
