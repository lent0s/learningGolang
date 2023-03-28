package main

import "fmt"

/*
Пользователь вводит будний день недели в сокращённой форме
(пн, вт, ср, чт, пт) и получает развёрнутый список всех
последующих рабочих дней, включая пятницу.

Рекомендация	Пример работы программы:

Дни недели.
Введите будний день недели: пн, вт, ср, чт, пт:
вт
вторник
среда
четверг
пятница
*/
func main() {

	var answer string

	fmt.Println("\n\n   -=[ Дни недели ]=-   \n")

	fmt.Print("\nВведите будний день недели: " +
		"пн, вт, ср, чт, пт: ")
	fmt.Scan(&answer)

	switch answer {
	case "пн":
		fmt.Println("понедельник")
		fallthrough
	case "вт":
		fmt.Println("вторник")
		fallthrough
	case "ср":
		fmt.Println("среда")
		fallthrough
	case "чт":
		fmt.Println("четверг")
		fallthrough
	case "пт":
		fmt.Println("пятница")
	default:
		fmt.Println("Не знаю такого дня")
	}

}
