package main

import "fmt"

type Students struct {
	id      int
	address string
	age     int
}

func main() {
	Nahid := Students{101, "bangladesh", 19}
	fmt.Println("id:", Nahid.id)
	fmt.Println("address:", Nahid.address)
	fmt.Println("age:", Nahid.age)
}