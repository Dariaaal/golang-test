package main

import (
	"fmt"
	"io/ioutil"
	// "os"
)

func file() {
	// читаємо файл text.txt
	file_data, err := ioutil.ReadFile("text.txt")
	// якщо помилка не дорівнює пустоті, тобто існує
	if err != nil {
		fmt.Println("Can't read this file \n", err)
	}
	// треба конвертувати в строку, бо тип в file_data це []byte
	fmt.Println(string(file_data))



	// створюємо файл
	data := []byte("Text to file")
    error := ioutil.WriteFile("text2.txt", data, 0600) //останнє це розширення файлу
	if error != nil {
		fmt.Println("Can't create file \n", error)
	}
	// Потім читаємо його
	file_data2, err := ioutil.ReadFile("text2.txt")
	if err != nil {
		fmt.Println("Can't read this file \n", err)
	}
	fmt.Println(string(file_data2))



	// видалення файлу
	// os.Remove("text2.txt")
}