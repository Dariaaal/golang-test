package main

import "fmt"

// структура це кастомний тип даних
type User struct {
	name string
	age int64
	password string
}

func changeu(u *User) {
	u.name = "Kate"
}

func structure() {
    // var user User = User{name: "John", age: 27, password: "pass"}
	user := User{"John", 23, "pass"}
	fmt.Println(user)
	// виведе {John, 27, pass}
	user.name = "Den"
	// змінюємо ім'я
	fmt.Println(user.name)


	changeu(&user)
	fmt.Println(user)
	// виведе {Kate, 27, pass}
}