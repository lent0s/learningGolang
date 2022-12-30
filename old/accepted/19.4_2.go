package main

import (
	"fmt"
)

/*
Отсортируйте массив длиной шесть пузырьком
*/

var (
	inputArray6 = []int{1, 2, 24, 6, 9, 6}
)

func bubbleSort(array []int) {
	for j := 0; j < len(array)-1; j++ {
		f := true
		for i := 0; i < len(array)-j-1; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				f = false
			}
		}
		if f {
			return
		}
	}
	return
}

func main() {
	bubbleSort(inputArray6)
	fmt.Println(inputArray6)
}
