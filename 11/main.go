package main

import "fmt"

func findNumber(num int) {
	var (
		temp int
	)
	for i := 1; i < num; i++ {
		if num%i == 0 {
			temp = i
		}
	}
	x := num / temp
	fmt.Print(x, " ")
	if temp > 25 {
		findNumber(temp)
	} else if temp == 1 {
		fmt.Println("")
	} else {
		fmt.Print(temp, " ")
	}
}
func main() {
	var n int
	fmt.Print("Input number: ")
	fmt.Scanln(&n)
	findNumber(n)
}
