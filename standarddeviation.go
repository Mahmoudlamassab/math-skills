package main

import "math/big"

func StandardDeviation(variance *big.Rat) float64 {
	// create a new big.Float, set 256 bits(how many binary digits are available)
	// convert variance big.Float approximation
	f := new(big.Float).SetPrec(256).SetRat(variance)
	// calculate the square root of f and store it in result
	result := new(big.Float).SetPrec(256).Sqrt(f)
	//convert the value to float64
	value, _ := result.Float64()

	return value
}
