package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedreturnRate := 5.5

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedreturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedreturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Real Value: %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}

func outputText(text string) {
	fmt.Println(text)
}

func calculateFutureValue(investmentAmount, expectedreturnRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1+expectedreturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)

	return fv, frv

}
