package main

import (
	"fmt"
)

func check(arr []int, n int) bool {
	for _, v := range arr {
		if v == n {
			return true
		}
	}
	return false
}
func main() {
	var n, x int
	fmt.Print("input n: ")
	fmt.Scanln(&n)

	arr := make([]int, n)

	fmt.Println("input number: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Print("input number do u wanna find: ")
	fmt.Scan(&x)
	result := check(arr, x)
	if result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
