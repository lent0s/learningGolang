package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
Реализуйте паттерн-конвейер:
Программа принимает числа из стандартного ввода в бесконечном цикле и передаёт число в горутину.
Квадрат: горутина высчитывает квадрат этого числа и передаёт в следующую горутину.
Произведение: следующая горутина умножает квадрат числа на 2.
При вводе «стоп» выполнение программы останавливается.
Советы и рекомендации
Воспользуйтесь небуферизированными каналами и waitgroup.

Что оценивается
Ввод : 3
Квадрат : 9
Произведение : 18
*/

func main() {
	wg := sync.WaitGroup{}
	chan1, chan2 := make(chan int64), make(chan int64)
	defer func() {
		close(chan1)
		close(chan2)
	}()
	userEnter := ""
	for {
		fmt.Print("Ввод (для выхода \"стоп\"): ")
		fmt.Scan(&userEnter)
		if userEnter == "стоп" {
			break
		}
		inputNum, err := strconv.ParseInt(userEnter, 10, 0)
		if err != nil {
			fmt.Println("Введите число")
			continue
		}
		wg.Add(1)
		go square(chan1, chan2)
		go func() {
			defer wg.Done()
			multiply(chan2)
		}()
		chan1 <- inputNum
		wg.Wait()
		fmt.Println()
	}
}

func square(chanIn, chanOut chan int64) {
	x := <-chanIn
	x *= x
	chanOut <- x
	fmt.Println("Квадрат:", x)
}

func multiply(chanIn chan int64) {
	x := <-chanIn
	x *= 2
	fmt.Println("Произведение:", x)
}
