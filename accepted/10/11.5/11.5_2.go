package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Напишите программу, которая выведет все части строки
a10 10 20b 20 30c30 30 dd,
которые можно привести к числу в десятичном формате.

Рекомендация	Пример работы программы:

Исходная строка:	a10 10 20b 20 30c30 30 dd
В строке содержатся числа в десятичном формате: 10 20 30
*/

func main() {

	firstStr := "a10 10 20b 20 30c30 30 dd"
	fmt.Print("Исходная строка:\n", firstStr, "\n\n")
	firstStr += " "
	newStr := ""

	for strings.Index(firstStr, " ") > 0 {
		tempStr := firstStr[:(strings.Index(firstStr, " ") + 1)]
		firstStr = strings.Replace(firstStr, tempStr, "", 1)
		tempStr = strings.Trim(tempStr, " ")
		_, err := strconv.ParseInt(tempStr, 10, 0)
		if err == nil {
			newStr += tempStr + " "
		}
	}

	fmt.Print("В строке содержатся числа в десятичном формате:"+
		"\n", newStr)

}
