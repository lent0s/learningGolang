package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Чётные и нечётные
Что нужно сделать
Напишите функцию, которая принимает массив чисел,
а возвращает два массива: один из чётных чисел, второй из нечётных.
*/

func makeArray() (array []int) {
	rand.Seed(time.Now().UnixNano())
	n := 2 + rand.Intn(13)
	array = make([]int, n, n)
	for i, _ := range array {
		array[i] = rand.Intn(99)
	}
	return
}

func printArray(array []int, title string) {
	fmt.Println(title)
	for i, _ := range array {
		fmt.Printf("%v\t", array[i])
	}
	fmt.Printf("\t[len= %v\t cap= %v]\n\n", len(array), cap(array))
}

func divForEvalArray(array []int) (evalArray []int, notEvalArray []int) {
	for _, r := range array {
		if r%2 == 0 {
			evalArray = append(evalArray, r)
		} else {
			notEvalArray = append(notEvalArray, r)
		}
	}
	return
}

func main() {
	inputArray := makeArray()
	//	inputArray = []int{0, 1, 2, 2, 4, 5, 1, 0, 0, 9}
	printArray(inputArray, "Исходный массив:")
	evalArray, notEvalArray := divForEvalArray(inputArray)
	printArray(evalArray, "Чётные числа из массива:")
	printArray(notEvalArray, "Нечётные числа из массива:")
}
