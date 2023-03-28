package main

import "fmt"

/*
Помогите человеку, желающему ездить только со счастливым
билетом (сумма первых трёх цифр должна быть равна сумме
последних трёх цифр). Для этого он готов покупать любое
количество билетов, но, с одной стороны, он не хочет
переплачивать, а с другой — он хочет быть твёрдо уверен,
что среди купленных билетов будет счастливый. Необходимо
написать программу, которая для билетов в диапазоне от
100000 до 999999 выведет, сколько минимум нужно купить
билетов, чтобы среди них оказался счастливый, если учесть,
что номер текущего билета мы не знаем. Иными словами, нам
надо найти максимальное расстояние между счастливыми
билетами. Так, после 100001 до следующего счастливого
числа — 100010 — придётся купить 9 билетов, а между
400220 и следующим счастливым
билетом — 400301 — расстояние будет уже в 81 билет.

Советы и рекомендации:	Необходимо запоминать предыдущий
счастливый билет и максимальное расстояние, найденное
ранее. При нахождении очередного счастливого билета
необходимо находить расстояние до предыдущего и сравнивать
его с максимальным. Если новое больше, то запоминать его.
В любом случае текущее счастливое число нужно запоминать,
как и предыдущее, и проверять числа дальше, до следующего
счастливого.

Пример работы программы:

Минимальное количество билетов, которое нужно купить,
чтобы среди них оказался счастливый: 1001
*/
func main() {

	minNum := 100000
	maxNum := 999999
	prevLucky := minNum
	var maxInterval, lPart, rPart int

	fmt.Println("\n\n   -=[ Счастливые билеты ]=-   \n")

	for minNum <= maxNum {
		lPart = minNum/100000 + minNum/10000%10 + minNum/1000%10
		rPart = minNum%10 + minNum/10%10 + minNum/100%10
		if lPart == rPart && minNum-prevLucky > maxInterval {
			maxInterval = minNum - prevLucky
		}
		if lPart == rPart {
			prevLucky = minNum
		}
		minNum++
	}

	fmt.Println("Минимальное количество билетов, которое "+
		"нужно купить, чтобы среди них оказался счастливый:",
		maxInterval)

}
