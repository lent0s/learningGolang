package main

import "fmt"

/*
Напишите программу, которая на вход получает число, затем с
помощью двух функций преобразует его. Первая умножает, а
вторая прибавляет число, используя именованные возвращаемые
значения.
*/

func multiplicate(x int) (y int) {
	y = x * x
	return
}

func addition(x int) (y int) {
	y = x + x
	return
}

func main() {

	fmt.Print("Введите число для преобразования: ")
	userEnter := 0
	fmt.Scanln(&userEnter)
	fmt.Printf("x = %v\nx * x = %v\nx + x = %v",
		userEnter, multiplicate(userEnter), addition(userEnter))

}
