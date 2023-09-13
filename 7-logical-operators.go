package main

import "fmt"

func logic() {
	a := 4
	b := 10


	// невірно, навіть якщо невірно одне
	if a > 1 && b > 5 {
		fmt.Print("And")
	}


	// вірне, якщо вірне хоча б одне
	if a > 1 || b > 5 {
		fmt.Print("Or")
	}


	// вірне, якщо вірне хоча б одне
	if a != 1 {
		fmt.Print("Not")
	}
}
