package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func getNumber(numbers []string) []int {
	slice := []int{}
	for _, num := range numbers {
		n, err := strconv.Atoi(num)
		if err != nil {
			fmt.Printf("Invalid value: %q", num)
			return nil
		}
		slice = append(slice, n)
	}
	return slice
}
func RoundRat(value *big.Rat) *big.Int {
	num := new(big.Int).Set(value.Num())
	den := value.Denom()

	quotient := new(big.Int)
	remainder := new(big.Int)

	quotient.QuoRem(num, den, remainder)

	twiceRemainder := new(big.Int).Mul(remainder, big.NewInt(2))

	if new(big.Int).Abs(twiceRemainder).Cmp(den) >= 0 {
		if value.Sign() < 0 {
			quotient.Sub(quotient, big.NewInt(1))
		} else {
			quotient.Add(quotient, big.NewInt(1))
		}
	}

	return quotient
}
