package main

import (
	"fmt"
	"math"
)

/*
(A-B)V=361
BV+A=9
V+A-B=38

(1-B)V=X
*/

func resultB(B float64) (x float64) {
	x = math.Pow(B, 9) + 2*math.Pow(B, 8) - 17*math.Pow(B, 7) - 401*math.Pow(B, 6) - 651*math.Pow(B, 5) + 435*math.Pow(B, 4) + 1355*math.Pow(B, 3) + 271*math.Pow(B, 2) - 740*B - 361
	return x
}

func main() {

	var A, B, C, r361, r9, r38, r float64
	const (
		cB = -.785
		//cB = -.965

	)

	C = resultB(float64(cB))
	A = (-math.Pow(C, 5) + 19*C*C*C - 9*C*C - 18*C + 9) / (-C*C*C*C - 2*C*C*C + 2*C + 1)
	B = (C*C - 27) / (1 - C*C)
	r361 = ((A-B)*C+A-B)*C + C - C
	r9 = (B*C+A-B)*C + A
	r38 = (C+A-B)*C + A - B
	r = C*C - B*C*C + A*C - B*C

	fmt.Printf("A=%v\tB=%v\tC=%v\n361=%v\t9=%v\t38=%v\nX=%v", A, B, C, r361, r9, r38, r)

}
