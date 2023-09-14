package main

import "fmt"

func forx(someFunc func(int) int) {
	fmt.Println(someFunc(25))
}

func test(x string) func() {
	return func() {
		fmt.Println(x)
	}
}


// main
func function2() {
	// функція всередині функції
    a := func() {
      fmt.Println("Hello from a")
	}
	a()


	b := func(a int, b int) int {
		result := a + b
		return result
	}
	sum := b(2, 8)
	fmt.Println(sum)


	// посилаємось на функцію forx 
	square := func(x int) int {
		return x * x
	}
    forx(square)



	// присвоюємо змінній "с" функцію test і викликаємо
	c := test("Hello from test")
	c()
	// або
	test("Hello from test")()

}