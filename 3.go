package main

import (
	"fmt"
)

func main() {
	fmt.Println("User 1 name?")
	var user1 string
	fmt.Scanln(&user1)
	fmt.Println("User 2 name?")
	var user2 string
	fmt.Scanln(&user2)

	fmt.Println("Hello " + user1 + " and " + user2)

}
