package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Нахождение первого вхождения числа в упорядоченном массиве
(числа могут повторяться)
Что нужно сделать		Заполните упорядоченный массив из 12
элементов и введите число. Необходимо реализовать поиск первого
вхождения заданного числа в массив. Сложность алгоритма должна
быть минимальная.

Что оценивается		Верность индекса.
При вводе массива 1 2 2 2 3 4 5 6 7 8 9 10 и вводе числа 2
программа должна вывести индекс 1.
*/
var index = -1

func main() {
	array := makeArray()
	//	array = []int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print("Введите искомое число: ")
	findNum := 0
	fmt.Scanln(&findNum)
	finding(array, 0, len(array)-1, findNum)
	if index != -1 {
		fmt.Printf("Первое вступление искомого числа в массиве на "+
			"позиции [%v]", index)
	} else {
		fmt.Println("В массиве отсутствует указанное число.")
	}
}

func finding(array []int, min, max, findNum int) {
	if array[min] > findNum || array[max] < findNum {
		return
	}
	if array[min] == findNum {
		index = min
		return
	}
	if array[min+(max-min)/2] < findNum {
		min += (max-min)/2 + 1
	} else {
		max -= (max - min) / 2
	}
	finding(array, min, max, findNum)
	return
}

func makeArray() (array []int) {
	rand.Seed(time.Now().UnixNano())
	n := 12
	array = make([]int, n, n)
	for i := 0; i < n; i++ {
		array[i] = i*n + rand.Intn(n)
	}
	printArray(array)
	return
}

func printArray(array []int) {
	for i := 0; i < len(array)-1; i++ {
		fmt.Printf("[%v]%v\t", i, array[i])
	}
	fmt.Println()
}
