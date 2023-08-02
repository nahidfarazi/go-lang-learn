package main

import "fmt"

func main() {
	var numbers int
	fmt.Printf("Enter your int number: ")
	fmt.Scan(&numbers)
	if numbers % 2 == 0 {
		fmt.Printf("Even")
	} else {
		fmt.Printf("Odd")
	}
}