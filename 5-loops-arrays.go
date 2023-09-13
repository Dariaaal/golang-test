package main

import "fmt"

func loopsarrays() {
	// only "for" exist


    for i  := 1; i<=5; i++ {
		fmt.Println("Hello", i)
	}


    for i  := 1; i<=5; {
		fmt.Println("Hello", i)
		i++
	}


	// Незкінченний цикл
	// for {
	// 	fmt.Println("Loop")
	// }


	// Виведе тільки парні числа
    for i :=1; i <=100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
	    }
    }


	// Також тільки парні
	for i :=1; i <=100; i++ {
		if i%2 != 0 {
			continue
	    }
		fmt.Println(i)
    }


	// Переривання циклу
	for i :=1; i <=100; i++ {
		if i == 50 {
			break
	    }
		fmt.Println(i)
    }


	// Типу while
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}


	// range
	nums := []int{1, 2, 3, 4, 5}  // Зріз чисел
	// for i :=0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }
	for index, element := range nums {
		fmt.Printf("Index: %d Element: %d\n", index, element)
	}


	// Двовимірний зріз-матриця
	// 1 2 3
	// 4 5 6
	// 7 8 9 
    matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for _, row := range matrix {
		for _, col := range row {
            fmt.Printf("%d ", col)
		}
		fmt.Println()
	}
}