package main

import (
	"fmt"
	"math"
)

func arr() {
	// var names [3]string 
	names := [3]string{"John", "Kate", "Mary"}
	// names[0] = "John"
	// names[1] = "Kate"
	// names[2] = "Mary"
	fmt.Println(names, len(names))

	for i := 0; i<len(names); i++ {
		fmt.Println(names[i])
	}


    //  Середнє арифметичне оцінок
	marks := [5]float64{11, 9, 12, 8, 11}
	var sum float64 = 0
	for i := 0; i<len(marks); i++ {
		sum += marks[i]
	}
	var result float64 = sum / float64(len(marks))
	fmt.Println(math.Round(result))
}

// багатовимірні масиви
// array := [3][5]int{{11, 9, 12, 8, 11}, {11, 9, 12, 8, 11}, {11, 9, 12, 8, 11}}