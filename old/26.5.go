package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
Написать программу аналог cat.
Программа должна получать на вход имена двух файлов, необходимо
конкатенировать их содержимое, используя strings.Join.
При получении одного файла на входе программа должна печатать его
содержимое на экран.
При получении двух файлов на входе программа соединяет их и
печатает содержимое обоих файлов на экран.
Если программа запущена командой go run firstFile.txt
secondFile.txt resultFile.txt, то она должна написать два
соединённых файла в результирующий.

first.txt		контент первого файла
second.txt		контент второго файла
result .txt		контент первого файла		контент второго файла
Input			go run first.txt second.txt result.txt

first.txt		контент первого файла
second.txt		контент второго файла
Input			go run first.txt second.txt
Output			контент первого файла		контент второго файла
*/

func result(firstF, secondF string) string {
	var firstStr, secondStr []string
	if len(firstF) != 0 {
		firstStr = loadFile(firstF)
	}
	resStr := strings.Join(firstStr, "\n")
	if len(secondF) != 0 {
		secondStr = loadFile(secondF)
		if len(resStr) > 0 {
			resStr += "\n"
		}
		resStr += strings.Join(secondStr, "\n")
	}
	return resStr
}

func loadFile(fileName string) []string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Невозможно прочитать файл", fileName)
		return nil
	}
	dataStr := strings.Split(string(data), "\n")
	return dataStr
}

func saveFile(fileName, inputStr string) {
	data, err := os.Create(fileName)
	defer data.Close()
	if err != nil {
		log.Fatalln("Ошибка создания/перезаписи файла", fileName)
		return
	}
	data.WriteString(inputStr)
	return
}

func main() {
	var firstFile, secondFile, resFile, resStr string
	flag.StringVar(&firstFile, "1st", "", "name of first file")
	flag.StringVar(&secondFile, "2nd", "", "name of second file")
	flag.StringVar(&resFile, "res", "", "name of resultant file")
	flag.Parse()
	if len(firstFile) == 0 && len(secondFile) == 0 && len(resFile) == 0 {
		log.Fatal("Работа программы завершена. Не введены имена файлов")
	}
	resStr = result(firstFile, secondFile)
	if len(resFile) != 0 {
		saveFile(resFile, resStr)
		return
	}
	fmt.Println(resStr)
}
