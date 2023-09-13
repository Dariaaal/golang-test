package main

import (
	"fmt"
	"math"
)

func mathf() {
	a := 2
	b := 4
	sum := a + b
	fmt.Println(sum)


	// int ціле число, float дробове
	var c float64 = 3
	var d float64 = 5
	answer := c / d
	fmt.Println(answer)


	// ділення з поверненням остачі %


	// math
	var e float64 = 144
	result := math.Sqrt(e)
	// корінь з 144(12)
	fmt.Println(result)


	// округляємо
	var f float64 = 25.0465656
	// вверх 26
	res := math.Ceil(f)
	fmt.Println(res)
	// вниз 25
	resp := math.Floor(f)
	fmt.Println(resp)
	// округляє
	respn := math.Round(f)
	fmt.Println(respn)
}