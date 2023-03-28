package main

import (
	"flag"
	"fmt"
	"log"
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
	s1Rune := []rune(str)
	s2Rune := []rune(subStr)
	for i := 0; i < len(s1Rune); i++ {
		for j := 0; j < len(s2Rune); j++ {
			if s1Rune[i+j] == s2Rune[j] {
				if j == len(s2Rune)-1 {
					return true
				}
			} else {
				break
			}
		}
	}
	return false
}

func main() {
	inputStr, subStr, answr := "какой-то текст", "кой", "НЕ СОДЕРЖИТСЯ"
	flag.StringVar(&inputStr, "str", "текст", "введите текст")
	flag.StringVar(&subStr, "substr", "искомый текст", "введите искомый текст")
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
