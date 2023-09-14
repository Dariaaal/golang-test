package main

import "fmt"

func first() {
	fmt.Println("Hello from first func")
}

func sum(a int, b int) {
    fmt.Println(a + b)
}

func sum2(a int, b int) int {
    result := a + b
	return result
}

func mathfunc(a int, b int) (int, int, int, int) {
	sum := a + b
	sub := a - b
	mult := a * b
	div := a / b
	return sum, sub, mult, div
}

// другий варіант запису верхньої функції
// func mathfunc(a int, b int) (sum int, sub int, mult int, div int) {
// 	sum = a + b
// 	sub = a - b
// 	mult = a * b
// 	div = a / b
// 	return 
// }



// main
func function() {
   first()

   sum(5, 9)

   value := sum2(4, 6)
   fmt.Println(value)

   sm, sb, m, d := mathfunc(3, 5)
   fmt.Println(sm, sb, m, d)
//    виведе 8 -2 15 0 
}