package main

import "fmt"

/*
Ну и какой же компьютер без игр? Давайте научим его играть в «Угадай число».
Пользователь загадывает число от 1 до 10 (включительно). Программа пытается
это число угадать, для этого она выводит число, а пользователь должен
ответить: угадала программа, больше загаданное число или меньше.

Рекомендация
Программа не должна делать больше четырёх попыток в процессе угадывания.
*/
func main() {

	minNum := 0
	maxNum := 0
	Num := 0
	var answer string

	fmt.Print("\n\n-= Игра «Угадай число» =-" +
		"\nЗагадайте целое число от (введите) ")
	fmt.Scanln(&minNum)
	Num = minNum
	fmt.Print("до (введите) ")
	fmt.Scanln(&maxNum)
	fmt.Println()

LYes:
	//	if (maxNum-minNum)%2 == 1 {
	//		Num += 1
	//	}
	Num = minNum + (maxNum-minNum)/2

LWrongA:
	fmt.Print("Это число больше или равно ", Num, " (y/n)? ")
	fmt.Scanln(&answer)

	if answer != "y" && answer != "n" {
		fmt.Println("Введите y, если ДА, либо n, если НЕТ")
		goto LWrongA
	}

	if answer == "y" {
		if Num == maxNum {
			goto LEnd
		}
		if maxNum-Num <= 1 {
			fmt.Print("Загаданное число ", Num, "? ")
			fmt.Scanln(&answer)
			if answer == "y" {
				goto LEnd
			} else {
				Num = maxNum
				goto LEnd
			}
		}
		minNum = Num
		goto LYes
	} else {
		if Num-minNum == 1 {
			Num = minNum
			goto LEnd
		}
		maxNum = Num - 1
		if maxNum-minNum == 1 {
			Num = maxNum
			goto LWrongA
		}
		//		Num /= 2
		//		if Num == maxNum {
		//			goto LEnd
		//		}
		goto LYes
	}

LEnd:
	fmt.Println("Мы угадали число! Это", Num)
}
