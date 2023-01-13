package main

import "fmt"

func main() {

	usedLetters := [118]rune{}
	i := 65
	for r, _ := range usedLetters {
		usedLetters[r] = rune(i)
		i++
		if i == 91 {
			i = 97
		}
		if i == 123 {
			i = 1025
		}
		if i == 1026 {
			i = 1040
		}
		if i == 1104 {
			i = 1105
		}
	}

	letters := []rune{'a'}
	for _, r := range usedLetters {
		letters[0] = r
		//	fmt.Printf("%c ", letters)
	}

	fmt.Printf("%c", usedLetters[:52])
}

//65-91		97-123		1025		1040-1104		1105
