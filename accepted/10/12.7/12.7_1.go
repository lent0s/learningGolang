package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

/*
Напишите программу, которая на вход получала бы строку, введённую
пользователем, а в файл писала № строки, дату и сообщение в формате:

2020-02-10 15:00:00 продам гараж.

При вводе слова exit программа завершает работу.
*/

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Невозможно создать файл", err)
	}
	defer file.Close()

	for i := 0; ; i++ {
		fmt.Println("Введите сообщение (для выхода из программы " +
			"[exit]):")
		userEnter, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Println()
		if userEnter == "exit\n" {
			fmt.Println("\n==========================" +
				"\nРабота программы завершена\n")
			break
		}

		file.WriteString(fmt.Sprintf("%v. [%v] %v",
			strconv.FormatInt(int64(i), 10),
			time.Now().Format("2006-01-02 15:04:05"), userEnter))
	}

	file, err = os.Open("log.txt")
	if err != nil {
		fmt.Println("Невозможно открыть файл", err)
	}
	defer file.Close()
	buf, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Невозможно загрузить данные из файла", err)
		return
	}

	fmt.Println(string(buf))

}
