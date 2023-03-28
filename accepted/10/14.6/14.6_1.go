package main

import "fmt"

/*
Напишите функцию, которая на вход получает число и возвращает
true, если число четное, и false, если нечётное.

Рекомендация	Программа запрашивает у пользователя или
генерирует случайное число, передает в функцию в качестве
аргумента и выводит в консоль результат её работы.
*/

func evenNum(enterNum int) bool {
	return enterNum%2 == 0
}

func main() {

	fmt.Print("Введите число для проверки чётности: ")
	userEnter := 0
	fmt.Scanln(&userEnter)
	if evenNum(userEnter) {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечётное")
	}

}
