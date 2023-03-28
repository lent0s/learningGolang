package main

import "fmt"

/*
Напишите функцию, которая на вход принимает функцию вида A func (int, int) int,
а внутри оборачивает и вызывает её при выходе (через defer).

Вызовите эту функцию с тремя разными анонимными функциями A. Тела функций
могут быть любыми, но главное, чтобы все три выполняли разное действие.
*/

func main() {
	withDefer(4, 7, func(a int, b int) int { return a + b })
	withDefer(4, 7, func(a int, b int) int { return a })
	withDefer(4, 7, func(a int, b int) int { return b })
}

func withDefer(a, b int, A func(int, int) int) {
	defer fmt.Println(A(a, b))
}
