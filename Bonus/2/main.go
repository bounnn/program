package main

import (
	"fmt"
)

func reverse(s string) string {
	rns := []rune(s) //rune = int32 data type
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
func main() {
	var str string
	fmt.Print("Please enter your string: ")
	fmt.Scanln(&str)
	strRev := reverse(str)
	fmt.Println("result: ", strRev)
}
