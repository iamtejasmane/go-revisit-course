package main

import (
	"errors"
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

func Divide(w http.ResponseWriter, r *http.Request) {
	res, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0!")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf(" %d / %d = %d", 100.0, 10.0, res))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		e := errors.New("Divide by 0 exception")
		return 0, e
	}
	res := x / y
	return res, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Server starting on port %s", portNumber)
	_ = http.ListenAndServe(":8080", nil)
}
