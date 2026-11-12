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
	quotient := new(big.Int)
	remainder := new(big.Int)
    //store the division of the numerator and the Denominator of value in quotient and the remainded
	quotient.QuoRem(value.Num(), value.Denom(), remainder)

	absRemainder := new(big.Int).Abs(remainder)//gets the absolute value of the remainder
	//multiply the absRemainder two times to know wether it bigger than the the  Denominator
	twiceRemainder := new(big.Int).Mul(absRemainder, big.NewInt(2))
   //compare twiceRemainder with the denominator 
	if twiceRemainder.Cmp(value.Denom()) >= 0 {
		//if it's negative we substract one else add one
		if value.Sign() < 0 {
			quotient.Sub(quotient, big.NewInt(1))
		} else {
			quotient.Add(quotient, big.NewInt(1))
		}
	}

	return quotient
}
