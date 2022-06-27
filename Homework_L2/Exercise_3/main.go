package main

import (
	"fmt"
	//	"math"
)

func main() {
	var a int
	fmt.Print("Введите число от 0 до 999 : ")
	fmt.Scanln(&a) 
	wordEnd(a)

}

func wordEnd(a int) {
	var a100 = (a) % 100
	var a101 = a100 / 10
	var a10 = a100 % 10
	var a1 = (a) / 100
	var b int
	var c int
	var d int
	//Проверяем и выставляем индекс для сотен
	if a1 >= 5 && a1 <= 20 {
		b = 2 // 5
	} else if a1 > 1 && a1 < 5 {
		b = 1 // 2
	} else if a1 == 1 {
		b = 0 // 1
	} else {
		b = 2 // 5}
	}
	//Проверяем и выставляем индекс для десятков
	if a101 >= 5 && a101 <= 20 {
		c = 2 // 5
	} else if a101 > 1 && a101 < 5 {
		c = 1 // 2
	} else if a101 == 1 {
		c = 0 // 1
	} else {
		c = 2 // 5}
	}
	//Проверяем и выставляем индекс для единиц
	if a10 >= 5 && a10 <= 20 {
		d = 2 // 5
	} else if a10 > 1 && a10 < 5 {
		d = 1 // 2
	} else if a10 == 1 {
		d = 0 // 1
	} else {
		d = 2 // 5}
	}

	arrHuns := [3]string{"сотня", "сотни", "сотен"}
	arrTens := [3]string{"десяток", "десятка", "десятков"}
	arrUnits := [3]string{"единица", "единицы", "единиц"}

	// Производим проверку что введено число
	if a < 0 || a > 999 || (a1 == 0 && a101 == 0 && a10 == 0) {
		fmt.Println("Ошибка ввода, данное значение не является числом или не находися в интермале [0, 999].")
	} else {
		fmt.Println(a1, arrHuns[b], a101, arrTens[c], a10, arrUnits[d])
	}
}