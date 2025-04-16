package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	hobbies := [3]string{"Travelling", "Gaming", "Coding"}
	fmt.Println(hobbies)
	// 2) Also output more data about that array:
	//		- The first element (standalone)
	fmt.Println(hobbies[0])
	//		- The second and third element combined as a new list
	fmt.Println(hobbies[1], hobbies[2])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	s := hobbies[:2]
	fmt.Print(s)
	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.

	s = hobbies[1:]
	fmt.Println(s)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"learn_go", "practice_go"}
	fmt.Println(goals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "make_projects"
	goals = append(goals, "make_tools")
	fmt.Println(goals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.

	p1 := Product{
		title: "Mobile",
		id:    1,
		price: 299,
	}

	p2 := Product{
		title: "Watch",
		id:    2,
		price: 59,
	}

	products := []Product{p1, p2}
	fmt.Println(products)

	products = append(products, Product{
		title: "Laptop",
		id:    3,
		price: 599,
	})

	fmt.Println(products)
}
