package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Input: ")
	fmt.Scanln(&a)
	fmt.Print("Input: ")
	fmt.Scanln(&b)
	fmt.Print("Input: ")
	fmt.Scanln(&c)
	arr := []int{a, b, c}
	// result := []int{}
	var (
		x1 = arr[0] + arr[1]
		x2 = arr[0] + arr[2]
		x3 = arr[1] + arr[2]
	)
	if x1 > x2 && x1 > x3 {
		fmt.Println(arr[0], " + ", arr[1], " = ", x1, " is Max")
	} else if x2 > x1 && x2 > x3 {
		fmt.Println(arr[0], " + ", arr[2], " = ", x2, " is Max")
	} else {
		fmt.Println(arr[1], " + ", arr[2], " = ", x3, " is Max")
	}
	// for i := 0; i < len(arr); i++ {
	// 	for j := i + 1; j < len(arr); j++ {
	// 		x := arr[i] + arr[j]
	// 		fmt.Println(x)
	// 	}
	// }
}
