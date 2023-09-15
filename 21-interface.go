package main

import "fmt"

type Numbers struct {
	num1 int
	num2 int
}

// інтерфейс містить методи
// з ними можна наслідувати методи у декількох структурах
type NumInterface interface {
	Sum() int
	Multiply() int
	Division() float64
	Substract() int
}

func (n Numbers) Sum() int {
	return n.num1 + n.num2
}

func (n Numbers) Multiply() int {
	return n.num1 * n.num2
}

func (n Numbers) Division() float64 {
	return float64(n.num1) / float64(n.num2)
}

func (n Numbers) Substract() int {
	return n.num1 - n.num2
}

func interf() {
	var i NumInterface
	numbers := Numbers{2, 8}
	i = numbers
	fmt.Printf("Сума: %d\n", i.Sum())
	fmt.Printf("Множення: %d\n", i.Multiply())
	fmt.Printf("Ділення: %f\n", i.Division())
	fmt.Printf("Віднімання: %d\n", i.Substract())
}