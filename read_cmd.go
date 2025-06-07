package main

import "fmt"

func ReadCmd() (string, error) {
	var input string

	fmt.Println("Пример названия файла: \"image.png\"")
	fmt.Println("Введите название файла: ")
	
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", err
	}

	return input, nil
}
