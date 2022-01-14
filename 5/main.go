package main

import "fmt"

func main() {
	var A, B int = 17, 9
	var (
		value1 = A / B
		value2 = A % B
	)
	fmt.Println("A Division B = ", value1)
	fmt.Println("A Modulus B = ", value2)
}
