package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Анонимные функции
Что нужно сделать
Напишите анонимную функцию, которая на вход получает массив типа integer,
сортирует его пузырьком и переворачивает (либо сразу сортирует в обратном
порядке, как посчитаете нужным).
*/
func makeArray() (array []int) {
	rand.Seed(time.Now().UnixNano())
	n := 5 + rand.Intn(15)
	array = make([]int, n, n)
	for i := 0; i < n; i++ {
		array[i] = rand.Intn(n * n)
		fmt.Printf("%v\t", array[i])
	}
	fmt.Println()
	return
}

func main() {
	fmt.Println("Исходный массив:")
	inputArray := makeArray()
	fmt.Println("Отсортированный на уменьшение массив:")
	//	inputArray = []int{0, 1, 2, 2, 4, 5, 1, 0, 0, 9}
	sortArray := func(array []int) (newArray []int) {
		min := len(array) - 1
		max := 0
		for min > max {
			for i := max; i < min; i++ {
				if array[i] < array[i+1] {
					array[i], array[i+1] = array[i+1], array[i]
				}
			}
			min -= 1

			//	цикл проходки в обратную сторону для ускорения
			for i := min; i > max; i-- {
				if array[i] > array[i-1] {
					array[i], array[i-1] = array[i-1], array[i]
				}
			}
			max += 1

		}
		return array
	}
	sortArray(inputArray)
	for i, _ := range inputArray {
		fmt.Printf("%v\t", inputArray[i])
		i += 0
	}
}
