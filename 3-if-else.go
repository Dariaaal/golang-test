package main

import "fmt"

func ifelse() {
	num := 3
	if num > 0 {
		fmt.Println("Number > 0")
	} else if num < 0 {
		fmt.Println("Number < 0")
	} else if num == 3 {
		fmt.Println("Number = 3")
	}
}