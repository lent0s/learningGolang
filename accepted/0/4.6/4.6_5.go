package main

import "fmt"

/*
Напишите программу, которая запрашивает день недели, число гостей
и сумму чека и рассчитывает сумму к оплате. В ресторане действуют
следующие правила:

  - по понедельникам должна применяться скидка 10% на всё меню,
    потому что понедельник — день тяжёлый;
  - по пятницам, если сумма чека превышает 10 000 рублей, включается
    дополнительная скидка в размере 5%;
  - если число гостей в одной компании превышает пять человек,
    автоматически включается надбавка на обслуживание 10%.

Рекомендация	Пример работы программы:

Введите день недели:		5
Введите число гостей:		7
Введите сумму чека:			12000
Скидка по пятницам: 		600
Надбавка на обслуживание:	1200
Сумма к оплате: 			12600
*/
func main() {

	dayOfWeek := 0
	countOfGuests := 0
	bill := 0
	discount := 0
	extraCharge := 0

	fmt.Println("\n***** ООО Ресторан *****\n")

LWrongDay:
	fmt.Print("Введите номер дня недели\n" +
		"(первый день недели - понедельник): ")
	fmt.Scanln(&dayOfWeek)
	if dayOfWeek < 1 || dayOfWeek > 7 {
		fmt.Println("Введите номер дня недели цифрой от 1 до 7.\n")
		goto LWrongDay
	}

LWrongGuest:
	fmt.Print("\nВведите число гостей: ")
	fmt.Scanln(&countOfGuests)
	if countOfGuests < 1 {
		fmt.Println("Кто-то, всё же, должен оплатить чек.\n" +
			"Введите целое число больше нуля.")
		goto LWrongGuest
	}

LWrongBill:
	fmt.Print("\nВведите сумму чека: ")
	fmt.Scanln(&bill)
	if bill < 1 {
		fmt.Println("Чек пуст. Рассчитывать нечего")
		goto LWrongBill
	}
	fmt.Println()

	if dayOfWeek == 1 {
		discount = bill / 10
		fmt.Println("Скидка по понедельникам (10%):", discount)
	}

	if dayOfWeek == 5 {
		discount = bill / 20
		fmt.Println("Скидка по пятницам (5%):", discount)
	}

	if countOfGuests > 5 {
		extraCharge = bill / 10
		fmt.Println("Надбавка на обслуживание (10%):", extraCharge)
	}
	fmt.Println("\nСумма к оплате: ", bill-discount+extraCharge)

}
