package main

import (
	"example.com/price-calculator/prices"
)

func main() {
	taxReates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxReates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}

}
