package main

import "math/big"

func StandardDeviation(variance *big.Rat) float64 {
	f := new(big.Float).SetPrec(256).SetRat(variance)

	result := new(big.Float).SetPrec(256).Sqrt(f)

	value, _ := result.Float64()

	return value
}
