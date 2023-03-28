package main

import (
	"fmt"
)

/*
Напишите программу, в которую пользователь будет вводить
четырёхзначный номер билета, а программа будет выводить,
является ли он зеркальным, счастливым или обычным билетом.

Советы и рекомендации
При решении задачи необходимо применить целочисленное
деление и остаток от деления. Примеры вывода программы:

1234 -> обычный билет
3425 -> счастливый билет
1221 -> зеркальный билет
*/
func main() {

	var ticketNum, dig1, dig2, dig3, dig4 int

LBegin:
	fmt.Println("\n\n-= Счастливый билет =-" +
		"\n(для выхода из программы введите не номер билета)\n")
	// при условии, что "счастливым" считается билет,
	// в котором сумма первой половины чисел равна второй
	fmt.Print("Введите четырёхзначный номер билета: ")
	fmt.Scan(&ticketNum)
	if ticketNum > 9999 || ticketNum < 1000 {
		goto LEnd
	}

	dig1 = ticketNum / 1000
	dig2 = ticketNum / 100 % 10
	dig3 = ticketNum % 100 / 10
	dig4 = ticketNum % 10

	if dig1 == dig4 && dig2 == dig3 {
		fmt.Println("Зеркальный билет \n")
		goto LBegin
	} else if dig1*dig2 == dig3*dig4 {
		fmt.Println("Cчастливый билет \n")
		goto LBegin
	} else {
		fmt.Println("Oбычный билет \n")
		goto LBegin
	}

LEnd:
	fmt.Println("\nРабота программы завершена")

}
