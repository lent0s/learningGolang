package main

import (
	"fmt"
	"math"
)

/*
Напишите функцию, производящую следующие вычисления.
S = 2 × x + y ^ 2 − 3/z, где x — int16, y — uint8, a z — float32.
Тип S должен быть во float32
*/

var (
	x   int16   = 2
	y   uint8   = 4
	z   float32 = 17.41
	acc int
)

func main() {
	newZ := incrZ(z)
	S := 2*float32(x) + float32(y)*float32(y)
	S -= 3 * float32(math.Pow10(acc)) / float32(newZ)
	fmt.Printf("[x= %v]\t[y= %v]\t[z= %.5f]\n[S(%T)= %.5f]", x, y, z, S, S)
}

func incrZ(z float32) (newZ int32) {
	if z != float32(int(z)) {
		for z != float32(int(z)) {
			z *= 10
			acc += 1
		}
	}
	newZ = int32(z)
	return
}
