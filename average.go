package main

import "math/big"

func AverageExact(slice []int) *big.Rat {
	//create a pointer to big.Int initialzed in 0
	sum := new(big.Int)

	for _, n := range slice {
		//add all numbers
		sum.Add(sum, big.NewInt(int64(n)))
	}
	//get number of elements(numbers)
	count := big.NewInt(int64(len(slice)))
	//return their frac: sum/count
	return new(big.Rat).SetFrac(sum, count)
}
