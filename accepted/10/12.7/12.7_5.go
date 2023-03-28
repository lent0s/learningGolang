package main

import (
	"fmt"
	"strings"
)

/*
Напишите программу, которая на вход принимала бы интовое
число и для него генерировала бы все возможные
комбинации круглых скобок.

Рекомендация	Первый пример вывода программы:
Введите количество пар скобок:	3
["((()))","(()())","(())()","()(())","()()()"]

Второй пример вывода программы:
Введите количество пар скобок:	1
["()"]
*/

func main() {

	fmt.Print("\nВведите количество пар скобок: ")
	bracketPairs := 0
	fmt.Scanln(&bracketPairs)
	brackets := ""
	x := "("
	for i := 0; i != bracketPairs; i++ {
		brackets += x
		if i == bracketPairs-1 && x == "(" {
			i = -1
			x = ")"
		}
	}
	result := "\"" + brackets + "\""
	count := 1

	for strings.Count(brackets, "()") != bracketPairs {
		count++
		switch {
		case strings.LastIndex(brackets, "(") <
			strings.LastIndex(brackets, "))"):
			brackets = brackets[:strings.LastIndex(brackets,
				"(")] + ")(" +
				brackets[strings.LastIndex(brackets, "(")+2:]
		case strings.LastIndex(brackets, "())") != -1:
			mark := strings.LastIndex(brackets, "())")
			brackets = brackets[:mark] + ")()" + brackets[mark+3:]
			mark += 2
			openBra := strings.Count(brackets[mark:], "(")
			if openBra >= 1 {
				closedBra := strings.Count(brackets[mark:], ")")
				for i := 0; i < openBra; i++ {
					brackets = brackets[:mark] + "("
					mark += 1
				}
				for i := 0; i < closedBra; i++ {
					brackets = brackets[:mark] + ")"
					mark += 1
				}
			}
		}
		result += ",\"" + brackets + "\""
	}

	fmt.Print("Получилось комбинаций: ", count,
		"\n", "[", result, "]")
}
