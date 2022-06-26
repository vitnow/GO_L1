package main

import (
	"fmt"
	"math"
	)

/*  Доработать калькулятор из методички: больше операций и валидации данных
 (проверка деления на 0, возведение в степень, факториал)
+ форматирование строки при выводе дробного числа ( 2 знака после точки)
*/
func factorial(n uint) uint {

	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var a, b, rest float64
	var op string
	fmt.Print("Введите число а: ")
	fmt.Scanln(&a)
	fmt.Print("Введите число b: ")
	fmt.Scanln(&b)
	fmt.Print("Введите арифметическую операцию (+, -, *, /, x2, x3, x^n, n!): ")
	fmt.Scanln(&op)
	if op == "+" || op == "-" || op == "*" {
		if op == "+" {
			rest = a + b
		} else if op == "-" {
			rest = a - b
		} else {
			rest = a * b
		}
		fmt.Printf("Результат выполнения операции: %.2f\n ", rest)     //%d\n
	} else if op == "/" && b != 0 {
		rest = a / b
		fmt.Printf("Результат выполнения операции: %.2f\n", rest)     //"%.2f\n
	} else if op == "n!" {
			var n uint64 = uint64(a)
			var m uint64 = uint64(b)
		fmt.Println("Результат выполнения операции n! -> числа a: ", factorial(uint(n)))
		fmt.Println("Результат выполнения операции n! -> числа b: ", factorial(uint(m)))
	} else if op == "x2" {
		rest = a * a
		restS := b * b
		fmt.Printf("Результат выполнения операции x2 -> числа a: %.0f\n", rest)
		fmt.Printf("Результат выполнения операции x2 -> числа b: %.0f\n", restS)
	} else if op == "x3" {
		rest = a * a * a
		restS := b * b * b
		fmt.Printf("Результат выполнения операции x3 -> числа a: %.0f\n", rest)
		fmt.Printf("Результат выполнения операции x3 -> числа b: %.0f\n", restS)
	} else if op == "x^n" {
		rest = math.Pow(a, b)
		fmt.Println("Результат выполнения операции x^n:", rest)
	
	} else {
		fmt.Println("Операция выбрана неверно")
	}
}