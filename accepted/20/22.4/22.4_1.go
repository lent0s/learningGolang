package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*	Подсчёт чисел в массиве
Заполните массив неупорядоченными числами на основе генератора случайных чисел.
Введите число. Программа должна найти это число в массиве и вывести,
сколько чисел находится в массиве после введённого.
При отсутствии введённого числа в массиве — вывести 0.
Для удобства проверки реализуйте вывод массива на экран.
*/

func main() {
	array := makeArray()
	//	array = []int{0, 5, 2, 17, 17, 4, 1, 17, 22, 55}
	fmt.Println("Введите искомое число: ")
	findNum := 0
	fmt.Scanln(&findNum)
	indexNum, contains := finding(array, findNum)
	printArray(array)
	if !contains {
		fmt.Println("Искомого числа в массиве не найдено. [0] вхождений.")
		return
	}
	i := 0
	for j := 0; j < len(indexNum); j++ {
		for i < indexNum[j] {
			fmt.Print("\t")
			i++
		}
		fmt.Printf("[%v]%v\t", indexNum[j], array[indexNum[j]])
		i++
	}
	fmt.Printf("\nПосле первого вхождения в массиве ещё чисел: [%v].",
		len(array)-1-indexNum[0])
}

func finding(array []int, f int) (contain []int, sign bool) {
	for i, r := range array {
		if r == f {
			contain = append(contain, i)
			sign = true
		}
	}
	return
}

func makeArray() (array []int) {
	rand.Seed(time.Now().UnixNano())
	n := 5 + rand.Intn(15)
	array = make([]int, n, n)
	for i := 0; i < n; i++ {
		array[i] = rand.Intn(n * n)
	}
	printArray(array)
	return
}

func printArray(array []int) {
	for i := 0; i < len(array); i++ {
		fmt.Printf("[%v]%v\t", i, array[i])
	}
	fmt.Println()
}
