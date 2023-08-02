package main

import "fmt"

func main() {
	numbers := -10
	
	if numbers > 0 {
		fmt.Printf("Positive \n")
	} else if numbers < 0 {
		fmt.Printf("Negative \n")
	} else {
		fmt.Printf("Zero \n")
	}
}