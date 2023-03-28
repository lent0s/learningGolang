package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Напишите программу, которая с помощью функции генерирует три
случайные точки в двумерном пространстве (две координаты), а
затем с помощью другой функции преобразует эти координаты по
формулам: x = 2 × x + 10, y = −3 × y − 5.
*/

func generatePoints() (x, y int) {
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(10)
	y = rand.Intn(10)
	fmt.Println("\n\tx =", x, "\t\ty =", y)
	return x, y
}

func changePoints(x, y int) {
	fmt.Println("New\tx =", 2*x+10, "\t\ty =", -3*y-5)
}

func main() {
	for i := 0; i < 3; i++ {
		changePoints(generatePoints())
	}
}
