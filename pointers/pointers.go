package main

import "fmt"

func main() {
	age := 32

	agePtr := &age

	fmt.Println(agePtr)

	add(age)
	fmt.Println(*agePtr)

	// fmt.Println("Age: ", age)
}

func add(age int) int {
	fmt.Println("Address in side fn :", &age)
	return 0
}
