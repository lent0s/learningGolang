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

func main() {
	inputArray := [10]int{}
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Исходный массив:")
	for i, _ := range inputArray {
		inputArray[i] = rand.Intn(100)
		fmt.Printf("%v\t", inputArray[i])
	}
	fmt.Println()
	fmt.Println("Отсортированный на уменьшение массив:")
	//	inputArray = [10]int{0, 1, 2, 2, 4, 5, 1, 0, 0, 9}
	sortArray := func(array [10]int) (newArray [10]int) {
		min := len(array) - 1
		max := 0
		for min-max > 0 {
			count := 1
			for i := max; i < min; i++ {
				if array[i] < array[i+1] {
					array[i], array[i+1] = array[i+1], array[i]
					count = 1
				} else {
					count++
				}
			}
			min -= count
			count = 1
			for i := min; i > max; i-- {
				if array[i] > array[i-1] {
					array[i], array[i-1] = array[i-1], array[i]
					count = 1
				} else {
					count++
				}
			}
			max += count
		}
		return array
	}
	backwardSortArray := sortArray(inputArray)
	for i, _ := range backwardSortArray {
		fmt.Printf("%v\t", backwardSortArray[i])
	}
}
