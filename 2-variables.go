package main

import "fmt"

func variables() {
    var number1 int8 = 5
	fmt.Println(number1)
	number2 := 5
	fmt.Println(number2)

    var name string
	var age uint8
	fmt.Println("What is your name?")
	fmt.Scan(&name)
	fmt.Println("Hello " + name)
	fmt.Println("How old are you?")
	fmt.Scan(&age)
	fmt.Println("Your age is " + fmt.Sprint(age))

	var h int8 = 10
	var g float32 = float32(h)
	fmt.Println(g)
}