package main

import (
	"fmt"
	"strconv"
)

func main() {
	var binary string
	fmt.Print("Input: ")
	fmt.Scanln(&binary)
	output, err :=
		strconv.ParseInt(binary, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Output = ", output)
}
