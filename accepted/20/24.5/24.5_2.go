package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/*
Поиск символов в нескольких строках
Что нужно сделать
Напишите функцию, которая на вход принимает массив предложений
(длинных строк) и массив символов типа rune, а возвращает 2D-массив,
где на позиции [i][j] стоит индекс вхождения символа j из chars в
последнее слово в предложении i (строку надо разбить на слова и взять
последнее). То есть сигнатура следующая:

func parseTest(sentences []string, chars []rune)

Пример входных данных
sentences := [4]string{"Hello world", "Hello Skillbox", "Привет Мир",
"Привет Skillbox"}
chars := [5]rune{'H','E','L','П','М'}
Пример вывода результата в первом элементе массива
'H' position 0
'E' position 1
'L' position 9
*/

/*
функция поиска вхождений рун в ПОСЛЕДНЕЕ слово предложения.
для учёта регистра рун установить registr=1
*/
func containLetters(text []string, letters []rune, registr int8) (target [][]int) {
	if registr < 0 || registr > 1 {
		fmt.Printf("(Значение учёта регистра букв было изменено на " +
			"стандартное: 0 (не учитывать)\n")
		registr = 0
	}
	target = make([][]int, len(text), len(text))
	for i, r := range text {
		target[i] = make([]int, len(letters), len(letters))
		for strings.LastIndex(r, " ") == len(r)-1 && len(r) > 1 {
			r = r[:len(r)-1]
		}
		div := strings.LastIndex(r, " ") + 1
		tempString := r[div:]
		for ii, rr := range letters {
			x, y := -1, -1
			if registr == 0 {
				x = strings.IndexRune(strings.ToLower(tempString), rr)
				y = strings.IndexRune(strings.ToUpper(tempString), rr)
				if x == -1 && y > -1 || y != -1 && y < x {
					x = y
				}
			} else {
				x = strings.IndexRune(tempString, rr)
			}
			if x > -1 {
				x += div
			}
			target[i][ii] = x
		}
	}
	return
}

/*
функция создания массива из случайного количества предложений не
более maxSentences со случайным количеством слов не более
maxCountWords случайной длины от minLenWord до maxLenWord
*/
func generateAbracadabra(maxSentences, maxCountWords, minLenWord,
	maxLenWord int) (text []string) {
	defaultNum, tempWs, tempLs := 4, 1, 1
	rand.Seed(time.Now().UnixNano())
	abc := usedLetters()
	if maxSentences < 1 {
		fmt.Printf("(Значение предложений было изменено"+
			" на стандартное: %v)\n", defaultNum)
		maxSentences = defaultNum
	} else if maxSentences != 1 {
		maxSentences = rand.Intn(maxSentences-1) + 1
	}
	if maxCountWords < 1 {
		fmt.Printf("(Значение количества слов было изменено"+
			" на стандартное: %v)\n", defaultNum)
		maxCountWords = defaultNum
	} else if maxCountWords == 1 {
		tempWs = 0
	}
	if maxLenWord < 2 {
		fmt.Printf("(Значение длины слов было изменено на"+
			" стандартное: %v)\n", defaultNum)
		maxLenWord = defaultNum
	} else if maxLenWord == 2 {
		tempLs = 0
	}
	if minLenWord < 1 || minLenWord >= maxLenWord {
		fmt.Printf("(Значение минимальной длины слов было"+
			" изменено на стандартное: %v)\n", maxLenWord/2)
		minLenWord = maxLenWord / 2
	}
	text = make([]string, maxSentences, maxSentences)
	for i, sentence := range text {
		for word := 0; word < rand.Intn(maxCountWords)*tempWs+1; word++ {
			for letter := 0; letter < rand.Intn(maxLenWord)*tempLs+minLenWord; letter++ {
				sentence += string(abc[rand.Intn(118)])
			}
			sentence += " "
		}
		text[i] = sentence[:len(sentence)-1]
	}
	return
}

func usedLetters() (abc [118]rune) {
	i := 65
	for r, _ := range abc {
		abc[r] = rune(i)
		i++
		switch i {
		case 91:
			i = 97
		case 123:
			i = 1025
		case 1026:
			i = 1040
		case 1104:
			i = 1105
		}
	}
	return
}

/*
функция создания массива случайного числа рун до maxLetters
латиницей или с добавлением кириллицы при addCyrillic=1
*/
func generateLetters(maxLetters, addCyrillic int) (letters []rune) {
	defaultNum := 4
	if maxLetters < 1 {
		fmt.Printf("(Значение количества искомых символов "+
			"было изменено на стандартное: %v)\n", defaultNum)
		maxLetters = defaultNum
	} else if maxLetters > 2 {
		maxLetters = rand.Intn(maxLetters-2) + 2
	}
	if addCyrillic != 1 && addCyrillic != 0 {
		fmt.Printf("(Значение использования кириллицы было" +
			" изменено на стандартное: 1)\n")
		addCyrillic = 1
	}
	letters = make([]rune, maxLetters, maxLetters)
	abc := usedLetters()
	for i, _ := range letters {
		letters[i] = abc[rand.Intn(52+addCyrillic*66)]
	}
	return
}

func main() {
	text := generateAbracadabra(4, 4, 4, 8)
	letters := generateLetters(4, 1)
	//text = []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет SkilLbox"}
	//letters = []rune{'H', 'E', 'L', 'П', 'М'}
	for i, _ := range text {
		fmt.Printf("[%v] ", text[i])
	}
	fmt.Println()
	for i, _ := range letters {
		fmt.Printf("[%v] ", string(letters[i]))
	}
	fmt.Println()
	targetArray := containLetters(text, letters, 0)
	for i, _ := range targetArray {
		fmt.Printf("В последнем слове %vго предложения [%v] первое "+
			"вхождение\n", i+1, text[i])
		for j, r := range targetArray[i] {
			if r > -1 {
				fmt.Printf("[%v] на [%v] позиции\n", string(letters[j]), r)
			} else {
				fmt.Printf("[%v] отсутствует\n", string(letters[j]))
			}
		}
	}
}
