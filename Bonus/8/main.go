package main

import "fmt"

func main() {
	var amout, total int
	fmt.Print("Input amout of products: ")
	fmt.Scanln(&amout)
	if amout < 10 {
		amout *= 20
		total = amout
	} else {
		amout *= 8
		total = amout
	}
	fmt.Println("total = ", total)
}
