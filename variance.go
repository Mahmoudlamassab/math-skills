package main

import "math/big"

func Variation(slice []int) *big.Rat {
	//get the exact average
	avr := AverageExact(slice)
   //rational accumulator initialized in 0
	var sum big.Rat

	for _, n := range slice {
		//calculate the distance from the avr for each num
		dis := new(big.Rat).Sub(
			new(big.Rat).SetInt64(int64(n)),
			avr,
		)
        //squared the value
		square := new(big.Rat).Mul(dis, dis)
        //store the squared vlues
		sum.Add(&sum, square)
	}
    //return the divising of the sum and the number of values
	return new(big.Rat).Quo(
		&sum,
		new(big.Rat).SetInt64(int64(len(slice))),
	)
}
