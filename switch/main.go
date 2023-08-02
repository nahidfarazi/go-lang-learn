package main

import "fmt"

func main() {
	var digit int
	fmt.Printf("Enter a digit 0-9: ")
	fmt.Scan(&digit)
	switch digit{
	case 0:
		fmt.Printf("zero \n")
	case 1:
		fmt.Printf("one \n")
	case 2:
		fmt.Printf("tow \n")
	case 3:
		fmt.Printf("three \n")
	case 4:
		fmt.Printf("four \n")
	case 5:
		fmt.Printf("five \n")
	case 6:
		fmt.Printf("six \n")
	case 7:
		fmt.Printf("seven \n")
	case 8:
		fmt.Printf("eight \n")
	case 9:
		fmt.Printf("nine \n")
	default:
		fmt.Printf("Not a valid digit")
	}
}