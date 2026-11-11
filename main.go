package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: go run . file.txt\n")
		return
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}
	numbers := strings.Fields(string(file))
	slice := getNumber(numbers)
	if slice == nil {
	}
	if len(slice) == 0 {
		fmt.Printf("Empty file\n")
		return
	}
	average := AverageExact(slice)
	median := Median((slice))
	variance := Variation(slice)
	standardDeviation := StandardDeviation(variance)
	fmt.Println("Average:", RoundRat(average))
	fmt.Println("Median:", RoundRat(median))
	fmt.Printf("Variance: %d\n", RoundRat(variance))
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(standardDeviation)))
}
