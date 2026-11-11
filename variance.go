package main

import "math/big"

func Variation(slice []int) *big.Rat {
	avr := AverageExact(slice)

	var sum big.Rat

	for _, n := range slice {
		dis := new(big.Rat).Sub(
			new(big.Rat).SetInt64(int64(n)),
			avr,
		)

		square := new(big.Rat).Mul(dis, dis)

		sum.Add(&sum, square)
	}

	return new(big.Rat).Quo(
		&sum,
		new(big.Rat).SetInt64(int64(len(slice))),
	)
}
