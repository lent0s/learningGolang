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

/*
Загружает данные из файлов f1 и f2 в строку
*/
func result(f1, f2 string) string {
	var s1, s2 []string
	if len(f1) != 0 {
		s1 = loadFile(f1)
	}
	sRes := strings.Join(s1, "\n")
	if len(f2) != 0 {
		s2 = loadFile(f2)
		if len(sRes) > 0 {
			sRes += "\n"
		}
		sRes += strings.Join(s2, "\n")
	}
	return sRes
}

/*
Загружает данные из файла f в массив строк
*/
func loadFile(f string) []string {
	data, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln("Невозможно прочитать файл", f)
		return nil
	}
	return strings.Split(string(data), "\n")
}

/*
Сохраняет строку f в файл str
*/
func saveFile(f, str string) {
	data, err := os.Create(f)
	defer data.Close()
	if err != nil {
		log.Fatalln("Ошибка создания/перезаписи файла", f)
		return
	}
	data.WriteString(str)
	fmt.Println("Произведена запись в файл:", f)
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
