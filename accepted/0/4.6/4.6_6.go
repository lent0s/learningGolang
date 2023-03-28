package main

import "fmt"

/*
Перед старостой группы стоит задача разделить весь курс, состоящий
из N студентов, на K групп. Напишите программу, которая поможет
старосте сделать это: он вводит N, K и порядковый номер студента,
а программа определяет, в какую группу он попадёт.

Рекомендация
В одну группу могут попадать студенты из разных частей списка.
*/
func main() {

	maxStudents := 0
	maxGroups := 0
	method := 0
	studentNum := 0
	groupNum := 0

	fmt.Println("\n-=Е   СТУДЕНТЫ   Ǝ=-\n")

LmaxStudents:
	fmt.Print("Введите количество студентов на потоке: ")
	fmt.Scanln(&maxStudents)
	if maxStudents < 1 {
		fmt.Println("Некому обучаться.\n")
		goto LmaxStudents
	}

LmaxGroups:
	fmt.Print("Введите количество групп студентов на потоке: ")
	fmt.Scanln(&maxGroups)
	if maxGroups < 1 {
		fmt.Println("Должна быть хоть одна группа.\n")
		goto LmaxGroups
	}

Lmethod:
	fmt.Print("Введите метод распределения студентов на потоке" +
		"\n(0 - если по порядку, 1 - если по заполнению групп): ")
	fmt.Scanln(&method)
	if method != 0 && method != 1 {
		fmt.Println("\n[0 - если по порядку]\n" +
			"[1 - если по заполнению групп]\n")
		goto Lmethod
	}

LstudentNum:
	fmt.Print("\nВведите порядковый номер студента\n" +
		"(или 0 для выхода из программы): ")
	fmt.Scanln(&studentNum)
	if studentNum < 0 || studentNum > maxStudents {
		fmt.Println("\nСтудент с таким порядковым номером отсутствует.")
		goto LstudentNum
	}
	if studentNum == 0 {
		goto LEnd
	}

	if method == 0 {
		groupNum = studentNum % maxGroups
		if groupNum == 0 {
			groupNum = maxGroups
		}
	} else {
		maxStudInGroup := maxStudents / maxGroups
		if maxStudents%maxGroups != 0 {
			maxStudInGroup += 1
		}
		temp := maxStudInGroup
		groupNum = 1
	L:
		if studentNum > temp && groupNum < maxGroups {
			groupNum += 1
			temp += maxStudInGroup
			goto L
		}
	}

	fmt.Println("Студент с порядковым номером"+
		"", studentNum, "определён в группу", groupNum)
	goto LstudentNum

LEnd:
	fmt.Println("\n\n\nЗавершение работы программы.")

}
