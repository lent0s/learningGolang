package main

import (
	"fmt"
)

/*
Напишите функцию, принимающую на вход массив и возвращающую массив,
в котором элементы идут в обратном порядке по сравнению с исходным.
Напишите программу, демонстрирующую работу этого метода.
Что оценивается		При вводе 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
программа должна выводить при помощи дополнительной функции,
реверсировав массив: 10, 9, 8, 7, 6, 5, 4, 3, 2, 1.
*/

func mirrorM(arrayIn []int) (arrayOut []int) {
	fmt.Println("Массив\t\tМассив\nисходный:\tновый:")
	for i, i2 := range arrayIn {
		arrayOut = append(arrayOut, arrayIn[len(arrayIn)-1-i])
		fmt.Printf("%v\t\t%v\n", i2, arrayOut[i])
	}
	return
}

func main() {
	numbers := []int{1, 22, 333, 4444, 55555}
	mirrorM(numbers)
}
