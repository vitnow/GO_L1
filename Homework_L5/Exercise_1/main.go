package main

import "fmt"

func fibonachi(i int) (int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonachi(i-1) + fibonachi(i-2)
}
func main() {
	var a int      					    
	fmt.Print("Введите количество элементов последовательности от 0 до 50: ")
	fmt.Scanln(&a)
	if a > 0 && a <= 50 {
	var i int
	for i = 0; i < a; i++ {
				fmt.Printf("%d ", fibonachi(i))
	} 
	} else {fmt.Println("Введенное значение некорректно. Попробуйте еще раз.")}
}
