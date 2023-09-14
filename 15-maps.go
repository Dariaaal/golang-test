package main

import "fmt"

func map1() {
	//  в картах значення зберігається під ключем, доступ за ключем, а не за порядковим номером

	var money map[string]int = map[string]int{
		"euros":   2000,
		"dollars": 1000,
		"hryvnas": 200000,
	}

	// простіший запис
	// money := map[string]int{
	// 	"euros":   2000,
	// 	"dollars": 1000,
	//  "hryvnas": 200000,
	// }

	fmt.Println(money)
	// виведе map[dollars: 1000, euros: 2000, hryvnas: 200000,]
	// також карти завжди сортують в алфавітному порядку

	fmt.Println(money["dollars"])
	//  виведе 1000

	// змінюємо
	money["dollars"] = 5000
	fmt.Println(money["dollars"])
	//  виведе 5000

	// видалення
	delete(money, "euros")

	el, status := money["dollars"]
	// в el присвоюємо значення по ключу dollars, а статус може бути true(якщо такий ключ знайдено і значення змінено) або false
	fmt.Println(el, status)
	// виведе 1000 true
}