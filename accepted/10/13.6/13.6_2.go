package main

import (
	"fmt"
)

/*
Напишите функцию, которая принимает в качестве аргументов
указатели на два типа int и меняет их значения местами.

Рекомендация	В методе main создайте и присвойте значения
двум переменным типа int, выведите значения этих переменных
в консоль до вызова функции и после.
*/

func mirror(a, b *int) {
	*a += *b
	*b = *a - *b
	*a -= *b
}

func main() {
	var a, b int
	fmt.Print("\nПервое число: ")
	fmt.Scanln(&a)
	fmt.Print("Второе число: ")
	fmt.Scanln(&b)
	mirror(&a, &b)
	fmt.Println("\nОтныне первое число:", a,
		"\nА второе:", b)
}