package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var d float64

	fmt.Print("Введите площадь круга -  a: ")
	fmt.Scanln(&a)

	d = math.Sqrt((4 * a) / math.Pi)

	if a > 0 {
		//S = d2 : 4 × π ->
		fmt.Println("Диаметр окружности = ", fmt.Sprintf("%.2f", d), "см")
		//L = D × π ->
		fmt.Println("Длина окружности = ", fmt.Sprintf("%.2f", d*math.Pi), "см")

	} else {
		fmt.Println("ВЫ ввели  некорректные данные. Попробуйте еще раз.")
	}
}