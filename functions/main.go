package main

import "fmt"

// custom type for func

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	doubled := transformNumbers(&numbers, double)
	trippled := transformNumbers(&numbers, triple)

	sq := transformNumbers(&numbers, getSquareFn())

	fmt.Println(doubled)
	fmt.Println(trippled)
	fmt.Println(sq)

}

func transformNumbers(numbers *[]int, fn transformFn) []int {
	res := []int{}
	for _, n := range *numbers {
		res = append(res, fn(n))
	}

	return res
}

func double(n int) int {
	return n * 2
}
func triple(n int) int {
	return n * 3
}

// returning function as a value
func getSquareFn() transformFn {
	return square
}

func square(n int) int {
	return n * n
}
