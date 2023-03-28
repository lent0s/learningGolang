package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
Научиться работать с композитными типами данных: структурами и картами
Напишите программу, которая считывает ввод с stdin, создаёт структуру student
и записывает указатель на структуру в хранилище map[studentName] *Student.

type Student struct {name string	age int		grade int}

Программа должна получать строки в бесконечном цикле, создать структуру
Student через функцию newStudent, далее сохранить указатель на эту структуру
в map, а после получения EOF (ctrl + d) вывести на экран имена всех студентов
из хранилища. Также необходимо реализовать методы put, get.

# Input			go run main.go
# Строки		# Вася 24 1			# Семен 32 2
# EOF
# Output		Студенты из хранилища:		# Вася 24 1			# Семен 32 2

Критерии оценки			Зачёт:
при получении одной строки (например, «имяСтудента 24 1») программа создаёт
студента и сохраняет его,далее ожидает следующую строку или сигнал EOF (Сtrl+Z);
при получении сигнала EOF программа должна вывести имена всех студентов из map.
*/

type Student struct {
	name  string
	age   int
	grade int
}

var class = map[string]*Student{}

/*
Проверяет строку s на корректность введённых данных и заносит их в структуру
*/
func checkInput(s *string, person *Student) {
	inputS := strings.Split((*s)[:len(*s)-2], " ")
	if len(*s) < 5 || len(inputS) < 3 {
		fmt.Println("Wrong input: too few arguments")
		return
	}
	num, err := strconv.ParseInt(inputS[len(inputS)-1], 10, 0)
	if err != nil {
		fmt.Println("Wrong grade: you must input a real number")
		return
	}
	person.grade = int(num)
	num, err = strconv.ParseInt(inputS[len(inputS)-2], 10, 0)
	if err != nil {
		fmt.Println("Wrong age: you must input a real number")
		return
	}
	person.age = int(num)
	person.name = strings.Title(strings.Join(inputS[:len(inputS)-2], " "))
}

/*
Отслеживает сигнал Ctrl+D в строке s
*/
func ctrlD(s string) bool {
	for _, r := range []byte(s) {
		if r == 4 {
			return true
		}
	}
	return false
}

/*
Заносит данные структуры в карту
*/
func (data Student) newStudent() {
	class[data.name] = &data
	fmt.Printf("\t\t\t\t      [accepted]\n")
}

func main() {
	student := Student{}
	for {
		fmt.Printf("Input a new student (name age grade): ")
		input, err := bufio.NewReader(os.Stdin).ReadString(10)
		if err == io.EOF || ctrlD(input) {
			fmt.Println("Input finished")
			break
		}
		checkInput(&input, &student)
		student.newStudent()
	}
	fmt.Printf("Total students [%v]:\n", len(class))
	for _, k := range class {
		fmt.Printf("Name: %v\t\t\tAge: %v\tGrade: %v"+
			"\n", k.name, k.age, k.grade)
	}
}
