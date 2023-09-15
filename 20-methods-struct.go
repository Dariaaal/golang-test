package main

import "fmt"

type User2 struct {
	name     string
	age      int64
	password string
	score []int
}

// метод діставання імені
func (u User2) getName() string {
	return u.name
}

// метод зміни імені
func (u *User2) setName(name1 string) {
	u.name = name1
}


// метод перевірки чи більше 18 років
func (u User2) isElder() bool {
	a := u.age
	isTrue := false
	if a >= 18 {
		isTrue = true
	} else if a < 18 {
		isTrue = false
	}
	return isTrue
}


// метод знаходить найвищий бал юзера
func (u User2) getHighScore() int {
	high := 0
	for _, sc := range u.score {
		if high < sc {
			high = sc
		}
	}
	return high
}

func structur() {
	user := User2{"John", 23, "pass", []int{10, 6, 8, 11, 20, 5}}
	fmt.Println(user.getName())
	// отпримаємо ім'я


	user.setName("Den")
	fmt.Println(user)


	fmt.Println(user.isElder())
	if user.isElder() {
		fmt.Println("Open")
	} else {
		fmt.Println("Closed")
	}


	fmt.Println(user.getHighScore())
}