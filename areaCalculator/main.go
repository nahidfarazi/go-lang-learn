package main

import (
	"fmt"
)

func main() {
	var base, height ,area float32
	fmt.Printf("base: ")
	fmt.Scan(&base)
	fmt.Printf("height: ")
	fmt.Scan(&height)
	area = 0.5 * base * height
	fmt.Printf("Area of triangle = %v", area)
}