package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

/*
Перепишите задачи 1 и 2, используя пакет ioutil.
*/

func main() {

	file, err := os.Create("ioutil.txt")
	if err != nil {
		fmt.Println("Невозможно создать файл.", err)
		return
	}
	defer file.Close()

	i := 0
	text := ""
	for {
		fmt.Println("Введите сообщение " +
			"(для выхода из программы [exit]):")
		userEnter, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Println()
		if userEnter == "exit\n" {
			fmt.Println("\n==========================" +
				"\nРабота программы завершена\n")
			break
		}
		i++
		text = text + fmt.Sprintf("%v. [%v] %v",
			strconv.FormatInt(int64(i), 10),
			time.Now().Format("2006-01-02 15:04:05"), userEnter)
	}

	err = ioutil.WriteFile("ioutil.txt",
		[]byte(text), 0644)
	if err != nil {
		fmt.Println("Невозможно произвести запись в файл.", err)
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Ошибка чтения файла.", err)
		return
	}
	if string(buf) == "" {
		fmt.Println("Файл пуст.")
	} else {
		fmt.Println(string(buf))
	}

}
