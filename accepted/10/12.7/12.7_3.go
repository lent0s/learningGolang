package main

import (
	"fmt"
	"os"
)

/*
Напишите программу, создающую текстовый файл только для
чтения, и проверьте, что в него нельзя записать данные.

Рекомендация	Для проверки создайте файл, установите
режим только для чтения, закройте его, а затем, открыв,
попытайтесь прочесть из него данные.
*/

func main() {
	file, err := os.Create("testAccess.txt")
	if err := os.Chmod("testAccess.txt",
		0444); err != nil {
		fmt.Println("Невозможно создать файл.", err)
		return
	}
	defer file.Close()
	file.Close()

	file, err = os.Open("testAccess.txt")
	if err != nil {
		fmt.Println("Невозможно прочесть файл.", err)
		return
	}

	_, err = file.WriteString("n66w")
	if err != nil {
		fmt.Println("Запись в файл невозможна.", err)
		return
	}

	fmt.Println("оп-па!")

}
