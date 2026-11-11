package main

import "math/big"

func Average(slice []int) float64 {
	var sum float64
	for _, n := range slice {
		sum += float64(n)
	}
	return float64(sum) / float64(len(slice))
}
func AverageFloat(slice []float64) float64 {
	var sum float64
	for _, n := range slice {
		sum += n
	}
	return float64(sum) / float64(len(slice))
}
func AverageExact(slice []int) *big.Rat {
	sum := new(big.Int)

	for _, n := range slice {
		sum.Add(sum, big.NewInt(int64(n)))
	}

	count := big.NewInt(int64(len(slice)))

	return new(big.Rat).SetFrac(sum, count)
}
