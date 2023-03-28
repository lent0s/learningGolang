package main

import (
	"fmt"
	"math"
)

/*
Напишите функцию, которая производит слияние двух отсортированных
массивов длиной четыре и пять в один массив длиной девять
*/
func uniteArray(array1 []int, array2 []int) (array []int) {
	array = make([]int, len(array1)+len(array2),
		len(array1)+len(array2))
	for i, r := range array1 {
		array[i] = r
	}
	for i, r := range array2 {
		array[i+len(array1)] = r
	}
	//array = append(array1, array2...) // медленнее
	return
}

func sortArray(array []int) (newArray []int) {
	// нужно использовать sort.Ints() - быстрее, и работает
	//в исходном массиве, но хотелось попробовать самому
	newArray = make([]int, len(array), len(array))
	minCurrentNum := math.MaxInt
	minGlobalNum := math.MinInt
	for i := 0; i < len(array); {
		for j := 0; j < len(array); j++ {
			if array[j] < minCurrentNum && array[j] > minGlobalNum {
				minCurrentNum = array[j]
			}
		}
		for j := 0; j < len(array); j++ {
			if minCurrentNum == array[j] {
				newArray[i] = minCurrentNum
				i++
			}
		}
		minGlobalNum = minCurrentNum
		minCurrentNum = math.MaxInt
	}
	return
}

var (
	inputArray4 = []int{17, 7, 24, 6}
	inputArray5 = []int{4, 4, 77, 2, 7}
)

func main() {
	outputArray9 := uniteArray(sortArray(inputArray4),
		sortArray(inputArray5))
	fmt.Println(sortArray(outputArray9))
}
