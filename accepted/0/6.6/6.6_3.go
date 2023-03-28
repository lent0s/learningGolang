package main

import "fmt"

/*
Напишите программу, которая принимает на вход цену товара и скидку.
Посчитайте и верните на экран сумму скидки. Скидка должна быть не
больше 30% от цены товара и не больше 2000 рублей.
*/
func main() {

	cost := 0
	discount := 1

	fmt.Println("\n\n   -=[ Расчёт суммы скидки ]=-   ")

	for discount != 0 {
		fmt.Print("\nВведите цену товара: ")
		fmt.Scanln(&cost)
		fmt.Print("Введите размер скидки в % (для выхода 0): ")
		fmt.Scanln(&discount)

		if discount > 30 {
			fmt.Println("Скидка не может быть больше 30%!")
		} else if cost*discount/100 > 2000 {
			fmt.Println("Скидка не может быть больше 2000 руб.!")
		} else {
			fmt.Println("Скидка", cost*discount/100, "руб.")
		}
	}

	fmt.Println("\nРабота программы завершена")

}
