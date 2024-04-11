package main

import (
	"fmt"
	"math"
)

// Interest calculation
func main() {
	// Constant value
	const inflationRate = 2.5

	// different ways of declaring variables
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	futureRealValue := futureValue * math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
