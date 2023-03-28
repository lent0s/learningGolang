package main

import (
	"fmt"
	"math"
)

/*
Достаточно часто при передаче по Сети или сохранении больших
объёмов данных приходится выбирать тип с минимальным размером
памяти, чтобы экономить трафик или место на диске. Напишите
программу, в которую пользователь вводит два числа (int16), а
программа выводит, в какой минимальный тип данных можно
сохранить результат умножения этих чисел.

Советы и рекомендации
Обратите внимание, что положительный результат можно сохранить
в меньшем типе за счёт использования uint8, uint16. Чтобы не
возникло проблем с переполнением в процессе умножения, числа
считывайте в int16, а перед умножением переведите их в int32.

Проверить на примерах:

1 1 результат uint8
1 −1 результат int8
640 100 результат uint16
−640 100 результат int32
3000 3000 результат uint32
−3000 3000 результат int32
*/
func main() {

	var firstNum, secondNum int16
	var answer string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&firstNum)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&secondNum)

	multi := int32(firstNum) * int32(secondNum)
	if multi < 0 {
		switch {
		case multi > math.MinInt8:
			answer = "int8"
		case multi > math.MinInt16:
			answer = "int16"
		default:
			answer = "int32"
		}
	} else {
		switch {
		case multi < math.MaxUint8:
			answer = "uint8"
		case multi < math.MaxUint16:
			answer = "uint16"
		default:
			answer = "uint32"
		}
	}

	fmt.Printf("%v * %v = %v, тип данных: %v \n",
		firstNum, secondNum, multi, answer)
}
