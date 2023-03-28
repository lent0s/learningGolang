package main

import "fmt"

/*
Представьте, что у вас есть три корзины разной ёмкости. Пользователю
предлагается ввести, какое количество яблок помещается в каждую
корзину. После этого программа должна заполнить все корзины по
очереди, учитывая, какие корзины уже заполнены, строго соблюдая
очерёдность заполнения и добавляя по одному яблоку в каждой итерации.

Советы и рекомендации	Используйте if и continue и break.

Пример: пользователь решил, что у корзин будет ёмкость на шесть,
четыре и девять яблок. Программа должна заполнить корзину 1 и в
следующей итерации перейти к корзине 2, далее в следующей итерации
перейти к корзине 3. Если очередная корзина уже заполнена, программа
должна переходить к следующей по очереди, и так по кругу, пока не
заполнит все.
*/
func main() {

	var basket1, basket2, basket3 int
	maxBasket1 := 1000
	maxBasket2 := 1000
	maxBasket3 := 1000
	var shiftL, shiftR string
	shiftL = "       "
	shiftR = "     "

	fmt.Println("\n\n   -=[ Корзины ]=-   ")

	for maxBasket1 > 999 || maxBasket2 > 999 || maxBasket3 > 999 ||
		maxBasket1 < 0 || maxBasket2 < 0 || maxBasket3 < 0 {
		fmt.Print("\n[Вместимость корзины не может быть " +
			"меньше нуля и больше 999]\n[Для выхода введите нули]" +
			"\nВведите вместимость первой корзины:  ")
		fmt.Scanln(&maxBasket1)
		fmt.Print("Введите вместимость второй корзины:  ")
		fmt.Scanln(&maxBasket2)
		fmt.Print("Введите вместимость третьей корзины: ")
		fmt.Scanln(&maxBasket3)
		if maxBasket1 == 0 && maxBasket2 == 0 && maxBasket3 == 0 {
			fmt.Println("Работа программы завершена")
			return
		}
	}

	fmt.Println("\n\nПЕРВАЯ  ВТОРАЯ  ТРЕТЬЯ\n0       0       0")

	for maxBasket1+maxBasket2+maxBasket3 > basket1+basket2+basket3 {
		if basket1 < maxBasket1 {
			basket1 += 1
			if basket1 < 10 {
				shiftL = "     "
			} else if basket1 < 100 {
				shiftL = "    "
			} else {
				shiftL = "   "
			}
			fmt.Println(basket1, shiftL, basket2, shiftR, basket3)
		}
		if basket2 < maxBasket2 {
			basket2 += 1
			if basket2 < 10 {
				shiftR = "     "
			} else if basket2 < 100 {
				shiftR = "    "
			} else {
				shiftR = "   "
			}
			fmt.Println(basket1, shiftL, basket2, shiftR, basket3)
		}
		if basket3 < maxBasket3 {
			basket3 += 1
			fmt.Println(basket1, shiftL, basket2, shiftR, basket3)
		}
	}

	fmt.Println("\nКорзины заполнены доверху, Хозяин")

}
