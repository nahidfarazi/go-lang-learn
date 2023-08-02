package main

import "fmt"

func main() {
	for i := 0; i <= 10; i+=3 {
		if i % 2 ==0 {
			fmt.Printf("%v \n",i)
		}
	}
	
}