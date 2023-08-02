package main

import "fmt"

func main() {
	var firstName,lastName string
	 var age int
	 var gpa float32

	//input: scan. scanln, scanf
	fmt.Print("Enter Your Name: ")
	fmt.Scan(&firstName, &lastName)
	fmt.Printf("%v is a student \n",firstName+ " "+lastName)
	fmt.Scanf("%v %v", &age, &gpa)
	fmt.Printf("%v is a %v year old \n",firstName+" "+lastName, age)
	fmt.Printf("%v got gpa %v",firstName+" "+lastName,gpa)
}