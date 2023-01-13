package main

import (
	"flag"
	"fmt"
	"log"
	"unicode/utf8"
)

/*
Написать программу для нахождения подстроки в кириллической подстроке.
Программа должна запускаться с помощью команды:

	go run main.go --str "строка для поиска" --substr "поиска"

Для реализации такой работы с флагами воспользуйтесь пакетом flags,
а для поиска подстроки в строке вам понадобятся руны.

Что нужно сделать
Спроектировать алгоритм поиска подстроки.
Определить строку и подстроку, используя флаги.
Написать алгоритм реализацию для работы со строками UTF-8
(для этого необходимо воспользоваться рунами).

Что оценивается
Алгоритм может работать с различными символами (кириллица, китайские
иероглифы). Использованы руны.

Input	привет мир!, вет							Output	true
Input	Программирование - это просто, вание		Output	true
Input	Программирование - это просто, корабль		Output	false
*/
func findingPhrase(str, subStr string) bool {
	tempSubStr := subStr
	contain := false
	subLetter, ss := utf8.DecodeRuneInString(subStr)
	for len(str) > 0 {
		letter, size := utf8.DecodeRuneInString(str)
		if subLetter == letter {
			subStr = subStr[ss:]
			if len(subStr) == 0 {
				contain = true
				break
			}
			subLetter, ss = utf8.DecodeRuneInString(subStr)
		} else if len(subStr) < len(tempSubStr) {
			subStr = tempSubStr
			subLetter, ss = utf8.DecodeRuneInString(subStr)
			continue
		}
		str = str[size:]
		if len(str) == 0 {
			break
		}
	}
	return contain
}

func main() {
	inputStr, subStr, answr := "какой-то текст", "кой", "НЕ СОДЕРЖИТСЯ"
	flag.StringVar(&inputStr, "str", "введённый текст", "введите текст")
	flag.StringVar(&subStr, "substr", "введённый искомый текст", "введите искомый текст")
	flag.Parse()
	if len(inputStr) == 0 || len(subStr) == 0 {
		log.Fatalln("Не указаны входные данные")
		return
	}
	if findingPhrase(inputStr, subStr) {
		answr = answr[5:]
	}
	fmt.Printf("В искомом тексте [%v] фраза [%v] - %v", inputStr, subStr, answr)
}
