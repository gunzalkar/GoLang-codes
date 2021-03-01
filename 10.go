package main

import "fmt"

func main() {
	foo()

}
func foo() {
	var syear, eyear, i int
	fmt.Println("Enter starting year")
	fmt.Scanln(&syear)
	fmt.Println("Enter end year")
	fmt.Scanln(&eyear)
	for i = syear; i <= eyear; i++ {
		if i%4 == 0 {
			fmt.Println("The leap years are ", i)
		}
	}
	fmt.Println("The totoal year count is", eyear-syear)
}
