package main

import "fmt"

func format() {
	age := 30
	name := "John"
	value := 1000.476
	a := true
	// fmt.Println("My age is " + fmt.Sprint(age))
	// fmt.Println("My age is", age)

	// форматування для цілих чисел(інт) %d, 
	// для строк %s, 
	// для дробових чисел %f
	// для boolean %t
	fmt.Printf("My age is %d. My name is %s. I have %f dollars. %t", age, name, value, a)
}