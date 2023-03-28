package main

import (
	"fmt"
)

/*
Разработайте программу, позволяющую ввести 10 целых чисел,
а затем вывести из них количество чётных и нечётных чисел.
Для ввода и подсчёта используйте разные циклы.

Что оценивается		Для введённых чисел 1, 1, 1, 2, 2, 2, 3,
3, 3, 4 программа должна вывести: чётных — 4, нечётных — 6.
*/

func main() {
	const len = 10
	var numbers [len]int

	for i, _ := range numbers {
		fmt.Printf("Введите %vе целое число: ", i+1)
		fmt.Scanln(&numbers[i])
	}
	count := 0
	for _, i2 := range numbers {
		if i2%2 == 0 {
			count++
		}
	}
	fmt.Printf("Для введённых чисел %v чётных - %v,"+
		" нечётных - %v", numbers, count, len-count)
}
