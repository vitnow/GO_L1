package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Print("Введите длину и ширину прямоугольника -  a, b: ")
	fmt.Scanln(&a, &b)

	if a > 0 && b > 0 {
		fmt.Println("Площадь прямоугольника", (a * b), "см")
	} else {
		fmt.Println("ВЫ ввели  некорректные данные. Попробуйте еще раз.")
	}
}