package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	fmt.Println(prices[0])
	fmt.Println(&prices[0])
	fmt.Println(&prices[1])
	prices[1] = 9.99

	prices = append(prices, 5.99)
	fmt.Println(prices)
	fmt.Println(&prices[0])
	fmt.Println(&prices[1])
	fmt.Println(&prices[2])

}

// func main() {
// 	// declare and initialize a array
// 	prices := [4]float64{10, 20.3, 4.6, 50.6}
// 	fmt.Println(prices)

// 	// declare a array
// 	var productNames [10]string
// 	fmt.Println(productNames)

// 	productNames = [10]string{"First_value"}
// 	productNames[5] = "Fifth_value"

// 	fmt.Println(productNames)
// 	fmt.Println(productNames[5])

// 	featurePrices := prices[1:]
// 	fmt.Println(featurePrices)

// 	// slice of slice
// 	featurePrices[0] = 199.99
// 	highlightedPrices := featurePrices[:1]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// }
