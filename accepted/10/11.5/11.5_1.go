package main

import (
	"fmt"
	"strings"
)

/*
Напишите программу, которая выведет количество слов, начинающихся
с большой буквы в строке: Go is an Open source programming Language
that makes it Easy to build simple, reliable, and efficient Software.

Рекомендация	Пример работы программы:

Определение количества слов, начинающихся с большой буквы в строке:
Go is an Open source programming Language that makes it Easy
to build simple, reliable, and efficient Software
Строка содержит 5 слов с большой буквы.
*/

func main() {
	firstStr := "Go is an Open source programming Language that" +
		" makes it Easy to build simple, reliable, and efficient" +
		" Software."
	fmt.Print("Определение количества слов, начинающихся с"+
		" большой буквы, в строке:\n", firstStr, "\n\n")
	countUpper := 0
	firstStr += " "
	upperABC := "QWERTYUIOPASDFGHJKLZXCVBNM"

	for strings.Index(firstStr, " ") > 0 {
		tempStr := firstStr[:(strings.Index(firstStr, " ") + 1)]
		firstStr = strings.Replace(firstStr, tempStr, "", 1)
		tempStr = tempStr[:1]
		if strings.Trim(tempStr, upperABC) < tempStr {
			countUpper++
		}
	}
	fmt.Printf("Строка содержит %v слов с большой буквы.", countUpper)
}
