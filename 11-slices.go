package main

import (
	"fmt"
	"sort"
)

// Зріз це динамічний масив, можна додавати, видаляти елементи
func slice() {
  slice := []int{3, 1, 2, 5, 7, 4}
//   додаємо
  slice = append(slice, 0)
// змінюємо
  slice[0] = 11
//   сортуємо по порядку
  sort.Ints(slice)
//   в алфавітному порядку
  slice2 := []string{"d", "e", "v", "b", "a"}

  fmt.Println(slice)
  fmt.Println(slice2)
}