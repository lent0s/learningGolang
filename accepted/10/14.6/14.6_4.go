package main

import "fmt"

/*
Напишите программу, в которой будет три функции, попарно
использующие разные глобальные переменные. Функции должны
прибавлять к поданному на вход числу глобальную переменную
и возвращать результат. Затем вызовите по очереди три
функции, передавая результат из одной в другую.
*/

var (
	tau  = 7
	piu  = 11
	lyau = 4
)

func add1(x int) int {
	x += tau
	fmt.Printf("+ %v = %v\n", tau, x)
	return x
}

func add2(x int) int {
	x += piu
	fmt.Printf("+ %v = %v\n", piu, x)
	return x
}

func add3(x int) int {
	x += lyau
	fmt.Printf("+ %v = %v\n", lyau, x)
	return x
}

func main() {

	fmt.Print("Введите число для преобразования: ")
	userEnter := 0
	fmt.Scanln(&userEnter)
	temp := userEnter
	add1(temp)
	temp = userEnter
	add2(temp)
	temp = userEnter
	add3(temp)
	fmt.Println()
	fmt.Println("\nИтог:", add3(add2(add1(userEnter))))

}
