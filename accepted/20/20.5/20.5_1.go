package main

import (
	"fmt"
	"math"
)

/*
Напишите функцию, вычисляющую определитель матрицы размером 3 × 3
*/

var (
	inputMatrix = [][]int{ /*
			{1, 17, 24, 1, 0, 0, 0},
			{2, 6, 	4, 	1, 0, 0, 0},
			{7, 12, 11, 1, 0, 0, 0},
			{1, 1, 	1, 	1, 0, 0, 0},
			{0, 0, 	0, 	0, 1, 0, 0},
			{0, 0, 	0, 	0, 0, 1, 0},
			{0, 0, 	0, 	0, 0, 0, 7}}	*/

		{1, 17, 24},
		{2, 6, 4},
		{7, 12, 11}}
)

func makeSmallerMatrix(oldMatrix [][]int, i int) (smallerMatrix [][]int) {
	cols := len(oldMatrix)
	smallerMatrix = make([][]int, cols-1, cols-1)
	for row := 0; row < cols-1; row++ {
		if i == 0 {
			smallerMatrix[row] = oldMatrix[row+1][i+1:]
		} else if i == cols-1 {
			smallerMatrix[row] = oldMatrix[row+1][:i]
		} else {
			for col := 0; col <= cols-1; col++ {
				if col == i {
					col++
				}
				smallerMatrix[row] = append(smallerMatrix[row], oldMatrix[row+1][col])
			}
		}
	}
	return
}

func determinant(matrix [][]int) (result int) {
	rows := len(matrix)
	cols := len(matrix[0])
	if rows != cols {
		fmt.Println("Детерминант имеет смысл только для квадратных матриц.")
		return
	}
	if rows > 2 {
		for i := 0; i < cols; i++ {
			result = result + matrix[0][i]*int(math.Pow(-1, float64(i)))*
				determinant(makeSmallerMatrix(matrix, i))
		}
	} else {
		result = matrix[0][0]*matrix[1][1] - matrix[1][0]*matrix[0][1]
	}
	return result
}

func main() {
	det := determinant(inputMatrix)
	fmt.Println("Для матрицы:")
	for i, _ := range inputMatrix {
		fmt.Println(inputMatrix[i])
	}
	fmt.Printf("Детерминант (определитель) равен: [%v]", det)
}
