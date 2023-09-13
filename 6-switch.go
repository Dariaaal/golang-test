package main

import "fmt"

func switchf() {
	// switch це альтернатива if, коли багато умов
	name := "Andrea"

	switch name {
	case "Jordan":
		fmt.Println("Hello Jordan")

	case "Kate":
		fmt.Println("Hello Kate")

	case "John":
		fmt.Println("Hello John")

	default:
		fmt.Println("I don't know you")
	}


	// з числами можна так:
	number := 8

	switch {
	case number > 1:
		fmt.Println("Number > 1")

	case number < 10:
		fmt.Println("Number < 10")

	default:
		fmt.Println("Another value")
	}


	// якщо задовільняються декілька умов, то можно додати fallthrough, щоб не спрацьовував break і ми дійшли до наступної умови
	num := 10

	switch {
	case num > 1:
		fmt.Println("Number > 1")
        fallthrough
	case num < 11:
		fmt.Println("Number < 11")
	}
}