package main

import (
	"fmt"
)

func main() {
	score := 70
	if score >= 80 && score <= 100 {
		fmt.Println("Grade A")
	} else if score >= 70 && score <= 79 {
		fmt.Println("Grade B")
	} else if score >= 60 && score <= 69 {
		fmt.Println("Grade C")
	} else if score >= 50 && score <= 59 {
		fmt.Println("Grade D")
	} else if score < 50 {
		fmt.Println("Grade F")
	} else {
		fmt.Println("Your score incorrect!!!!")
	}
}
