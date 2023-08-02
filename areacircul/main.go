package main

import "fmt"

func main() {
	var radius,area float32
	fmt.Printf("enter radius: ")
	fmt.Scan(&radius)
	area = 3.1416 * radius *radius
	fmt.Printf(" Area of circle: %v",area)
}