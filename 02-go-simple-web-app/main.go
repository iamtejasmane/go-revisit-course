package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the main home page function.
func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to my home page!")
	if err != nil {
		fmt.Println(err)
	}
}

// About is the about page of the app.
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValue(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is an about page and 2 + 2 = %d", sum))
}

// addValue function is to add two value and return the sum.
func addValue(num1, num2 int) int {
	sum := num1 + num2
	return sum
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Server starting on port %s", portNumber)
	_ = http.ListenAndServe(":8080", nil)
}
