package main

import "fmt"

func range2() {
	// продвинутий спосіб перебору масивів та зрізів
	slice := []int{1, 4, 5, 2, 6}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }

	for i, el := range slice {
	// el == slice[i]
	    fmt.Printf("%d: %d\n", i, el)
	}

	for _, el := range slice {
		fmt.Printf("%d\n", el)
	}
}