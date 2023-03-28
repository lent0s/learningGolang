package main

import "fmt"

/*
Программное обеспечение банкоматов постоянно решает задачу, как
имеющимися купюрами сформировать сумму, введённую пользователем.
Попробуйте решить похожую задачу и определить, сможет ли пользователь
заплатить за товар без сдачи или нет. Для этого он будет вводить
стоимость товара и номиналы трёх монет.

Рекомендация	Проверьте все сценарии, когда сумма может быть
сформирована одной монетой, двумя или всеми тремя.
*/
func main() {

	var cost,
		value1, value2, value3,
		minVal, midVal, maxVal,
		count1, count2, count3 int

	fmt.Println("\nСумма без сдачи\n")

	fmt.Print("Введите стоимость товара: ")
	fmt.Scan(&cost)
	fmt.Print("\nВведите номинал первой монеты:  ")
	fmt.Scan(&value1)
	fmt.Print("Введите номинал второй монеты:  ")
	fmt.Scan(&value2)
	fmt.Print("Введите номинал третьей монеты: ")
	fmt.Scan(&value3)

	// можно было написать код только для чисел больше нуля,
	// но задача не в этом, хотя, если необходимо - могу исправить.
	// не стал тратить Ваше время на проверку этого

	if value1 < value2 && value1 < value3 {
		minVal = value1
		if value2 < value3 {
			midVal = value2
			maxVal = value3
		} else {
			midVal = value3
			maxVal = value2
		}
	}
	if value2 < value1 && value2 < value3 {
		minVal = value2
		if value1 < value3 {
			midVal = value1
			maxVal = value3
		} else {
			midVal = value3
			maxVal = value1
		}
	}
	if value3 < value1 && value3 < value2 {
		minVal = value3
		if value1 < value2 {
			midVal = value1
			maxVal = value2
		} else {
			midVal = value2
			maxVal = value1
		}
	}

	count3 = cost / maxVal
	if cost%maxVal == 0 {
		goto LEnd
	} else {
	L1: // пока ещё не изучил циклы, return и т.п. О_О
		count2 = (cost - count3*maxVal) / midVal
	L2:
		if (cost-count3*maxVal)%midVal == 0 {
			goto LEnd
		} else {
			if (cost-count3*maxVal-count2*midVal)%minVal == 0 {
				count1 = (cost - count3*maxVal - count2*midVal) / minVal
				goto LEnd
			} else if count3 != 0 {
				count3 -= 1
				goto L1
			} else if count2 != 0 {
				count2 -= 1
				goto L2
			} else {
				fmt.Println("\nНе получается оплатить без сдачи " +
					"имеющимися монетами")
				goto L3
			}
		}
	}

LEnd:
	fmt.Println("\nДля оплаты необходимо:"+
		"\n", count3, "монет номиналом", maxVal, "руб. и"+
		"\n", count2, "монет номиналом", midVal, "руб. и"+
		"\n", count1, "монет номиналом", minVal, "руб.")
L3:
}
