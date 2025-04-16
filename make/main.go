package main

import "fmt"

func main() {
	names := make([]string, 3, 6)

	names[0] = "Julie"

	fmt.Println(names)

	names = append(names, "Max")
	names = append(names, "Manuel")
	names = append(names, "Tejas")

	fmt.Println(names)

	for n, i := range names {
		fmt.Println(n, i)
	}
}
