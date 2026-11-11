package main

import "math/big"

func SortInt(table []int) []int {
	for i := 0; i < len(table); i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
	return table
}
func Median(s []int) *big.Rat {
	sorted := SortInt(append([]int(nil), s...))

	middle := len(sorted) / 2

	if len(sorted)%2 != 0 {
		return new(big.Rat).SetInt64(int64(sorted[middle]))
	}

	sum := new(big.Int).Add(
		big.NewInt(int64(sorted[middle])),
		big.NewInt(int64(sorted[middle-1])),
	)

	return new(big.Rat).SetFrac(sum, big.NewInt(2))
}
