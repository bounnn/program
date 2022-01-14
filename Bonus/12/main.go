package main

import (
	"fmt"
	"math/big"
)

func main() {
	var fact = new(big.Int)
	var n int
	fmt.Print("Input: ")
	fmt.Scanln(&n)
	fact.MulRange(1, int64(n))
	// fmt.Println(fact)

	str := fact.String()
	sum := 0
	for _, val := range str {
		sum += int(val - '0')
	}
	fmt.Println(sum)
}
