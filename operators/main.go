package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Printf("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Printf("Enter last number: ")
	fmt.Scan(&num2)
	result := num1+num2
	fmt.Printf("%v + %v = %v ",num1,num2,result)

	

}