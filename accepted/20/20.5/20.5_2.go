package main

import (
	"fmt"
)

/*
Напишите функцию, умножающую две матрицы размерами 3 × 5 и 5 × 4
*/

var (
	matrixA = [][]int{
		{1, 17, 24, 44, 42},
		{2, 6, 4, 12, 74},
		{7, 12, 11, 13, 13}}

	matrixB = [][]int{
		{1, 17, 24, 1},
		{2, 6, 4, 1},
		{7, 12, 11, 1},
		{1, 1, 1, 1},
		{2, 4, 6, 8}}
)

func multiMatrix(A [][]int, B [][]int) (C [][]int) {
	if len(A[0]) != len(B) {
		fmt.Println(
			"Матрицы НЕ согласованы!\n" +
				"Операция умножения двух матриц выполнима только в том случае,\n" +
				"если число столбцов в первом сомножителе равно числу строк во втором.")
		return
	}
	C = make([][]int, len(A), len(A))
	for i, _ := range A {
		C[i] = append(multiIJ(A, B, i))
	}
	return
}

func multiIJ(A [][]int, B [][]int, i int) (C []int) {
	C = make([]int, len(B[0]), len(B[0]))
	for j, _ := range B[0] {
		for k, _ := range A[0] {
			C[j] = C[j] + A[i][k]*B[k][j]
		}
	}
	return
}

func main() {
	fmt.Printf("Результатом умножения матрицы А "+
		"[%vx%v]\n\n", len(matrixA), len(matrixA[0]))
	for i, _ := range matrixA {
		fmt.Println(matrixA[i])
	}
	fmt.Printf("\nна матрицу В [%vx%v]\n\n", len(matrixB), len(matrixB[0]))
	for i, _ := range matrixB {
		fmt.Println(matrixB[i])
	}
	fmt.Printf("\nбудет матрица С [%vx%v]\n\n", len(matrixA), len(matrixB[0]))
	AB := multiMatrix(matrixA, matrixB)
	for i, _ := range AB {
		fmt.Println(AB[i])
	}
	fmt.Printf("\nРезультатом обратного умножения ВхА будет C "+
		"[%vx%v]\n\n", len(matrixB), len(matrixA[0]))
	BA := multiMatrix(matrixB, matrixA)
	for i, _ := range BA {
		fmt.Println(BA[i])
	}
}
