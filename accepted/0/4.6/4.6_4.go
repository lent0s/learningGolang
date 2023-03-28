package main

import "fmt"

/*
Напишите программу, которая запрашивает у пользователя три
числа и выводит количество чисел, которые больше или равны пяти.

Рекомендация	Пример работы программы:

Три числа.
Введите первое число:	3
Введите второе число:	5
Введите третье число:	7
Среди введённых чисел 2 больше или равны 5.
*/
func main() {

	number := 0
	countOfNum := 0
	var equals string

	fmt.Println("\nТри числа (2).\n")

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&number)
	if number >= 5 {
		countOfNum += 1
	}

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&number)
	if number >= 5 {
		countOfNum += 1
	}

	fmt.Print("Введите третье число: ")
	fmt.Scanln(&number)
	if number >= 5 {
		countOfNum += 1
	}

	if countOfNum < 2 {
		equals = "равно"
	} else {
		equals = "равны"
	}
	fmt.Println("\nСреди введённых чисел"+
		"", countOfNum, "больше или", equals, "5.")

}
