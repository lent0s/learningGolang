package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Сортировка вставками
Что нужно сделать
Напишите функцию, сортирующую массив длины 10 вставками.
*/

func sortPasted(array [10]int) (sortedArray [10]int) {
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			for j := i; array[j] < array[j-1]; j-- {
				array[j], array[j-1] = array[j-1], array[j]
				if j == 1 {
					break
				}
			}
		}
	}
	for _, v := range array {
		fmt.Printf("%v\t", v)
	}
	return array
}

func main() {
	inputArray := [10]int{}
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Исходный массив:")
	for i, _ := range inputArray {
		inputArray[i] = rand.Intn(100)
		fmt.Printf("%v\t", inputArray[i])
	}
	fmt.Println()
	fmt.Println("Отсортированный массив:")
	sortPasted(inputArray)
}
