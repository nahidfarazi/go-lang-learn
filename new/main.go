package main

import "fmt"

func main() {
	var decimalNum int
	fmt.Printf("Enter Decimal Number: ")
	fmt.Scanf("%v",&decimalNum)
	fmt.Printf("Binary Number: %b \n",decimalNum)
	fmt.Printf("Octal Number: %o \n",decimalNum)
	fmt.Printf("HexaDecimal Number: %x \n",decimalNum)
}