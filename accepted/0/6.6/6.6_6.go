package main

import "fmt"

/*
В доме — 24 этажа. Лифт должен ходить вверх-вниз, пока не доставит
всех пассажиров на первый этаж. Три пассажира ждут на четвёртом,
седьмом и десятом этажах. При движении вверх лифт не должен
останавливаться, при движении вниз — должен собирать всех, но не
более двух человек в лифте. При этом лифт каждый раз доезжает до
самого верхнего этажа и только после этого начинает движение вниз.
Напишите программу, которая доставит всех пассажиров на первый этаж.
*/
func main() {

	var passenger1, passenger2, passenger3, floor int
	slot1 := "пусто"
	slot2 := "пусто"
	height := -1

	fmt.Println("\n\n   -=[ Движение лифта ]=-   \n")

	for height < passenger1 || height < passenger2 ||
		height < passenger3 || height < floor || passenger1 < 1 ||
		passenger2 < 1 || passenger3 < 1 || floor < 1 {
		fmt.Print("\nСколько этажей в доме? Введите: ")
		fmt.Scan(&height)
		fmt.Print("С какого этажа вызываем лифт? Введите: ")
		fmt.Scan(&floor)
		fmt.Print("На каком этаже Пассажир1? Введите: ")
		fmt.Scan(&passenger1)
		fmt.Print("На каком этаже Пассажир2? Введите: ")
		fmt.Scan(&passenger2)
		fmt.Print("На каком этаже Пассажир3? Введите: ")
		fmt.Scan(&passenger3)
		fmt.Println()
	}

	for passenger1 > 1 || passenger2 > 1 || passenger3 > 1 {
		if slot1 == "пусто" && slot2 == "пусто" {
			for floor < height {
				fmt.Println("Лифт на [", floor, "] этаже."+
					" В кабине [", slot1, "] и [", slot2, "]. "+
					"Движение вверх")
				floor += 1
			}
		}

		for floor > 1 {
			if slot1 == "пусто" {
				if floor == passenger1 {
					slot1 = "Пассажир1"
				} else if floor == passenger2 {
					slot1 = "Пассажир2"
				} else if floor == passenger3 {
					slot1 = "Пассажир3"
				}
			}

			if slot1 != "пусто" && slot2 == "пусто" {
				if floor == passenger1 && slot1 != "Пассажир1" {
					slot2 = "Пассажир1"
				} else if floor == passenger2 && slot1 != "Пассажир2" {
					slot2 = "Пассажир2"
				} else if floor == passenger3 && slot1 != "Пассажир3" {
					slot2 = "Пассажир3"
				}
			}

			fmt.Println("Лифт на [", floor, "] этаже. В кабине [",
				slot1, "] и [", slot2, "]. Движение вниз")
			floor -= 1
			if slot1 == "Пассажир1" || slot2 == "Пассажир1" {
				passenger1 -= 1
			}
			if slot1 == "Пассажир2" || slot2 == "Пассажир2" {
				passenger2 -= 1
			}
			if slot1 == "Пассажир3" || slot2 == "Пассажир3" {
				passenger3 -= 1
			}
			if floor == 1 {
				slot1 = "пусто"
				slot2 = "пусто"
			}

		}

	}
	fmt.Println("Лифт на [", floor, "] этаже. В кабине [", slot1,
		"] и [", slot2, "]. Ожидание команд")
	fmt.Println("\nВсе пассажиры доставлены")

}
