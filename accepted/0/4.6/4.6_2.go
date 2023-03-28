package main

import "fmt"

/*
Напишите программу, которая запрашивает у пользователя три
числа и сообщает, есть ли среди них число больше пяти.

Рекомендация	Пример работы программы:

Три числа.
Введите первое число:	3
Введите второе число:	6
Введите третье число:	2
Среди введённых чисел есть число больше 5.
*/
func main() {

	firstNum := 0
	secondNum := 0
	thirdNum := 0

	fmt.Println("\nТри числа.\n")

	fmt.Print("Введите первое число: ")
	fmt.Scan(&firstNum)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&secondNum)
	fmt.Print("Введите третье число: ")
	fmt.Scan(&thirdNum)
	fmt.Println()

	if firstNum > 5 || secondNum > 5 || thirdNum > 5 {
		fmt.Println("Среди введённых чисел есть число больше 5")
	} else {
		fmt.Println("Среди введённых чисел отсутствует число больше 5.")
	}
}
