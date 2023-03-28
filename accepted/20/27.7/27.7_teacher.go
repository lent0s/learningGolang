package main

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

type Student struct {
	Name  string
	Age   int
	Grade int
}

func (s Student) String() string {
	return s.Name + " " + strconv.Itoa(s.Age) + " " + strconv.Itoa(s.Grade)
}

type Storage struct {
	database map[string]*Student
}

func (s *Storage) Get(name string) *Student {
	return s.database[name]
}

func (s *Storage) GetAll() []*Student {
	res := []*Student{}
	for _, v := range s.database {
		res = append(res, v)
	}
	return res
}

func (s *Storage) Put(student *Student) error {
	if student == nil {
		return fmt.Errorf("invalid student: struct is nil")
	}
	if student.Name == "" {
		return fmt.Errorf("invalid student: name is empty")
	}
	s.database[student.Name] = student
	return nil
}

func NewStorage() *Storage {
	return &Storage{
		database: make(map[string]*Student),
	}
}

func main() {
	storage := NewStorage()
	for {
		s := Student{}
		fmt.Println("Введите имя, возраст и оценку студента")
		_, err := fmt.Scan(&s.Name, &s.Age, &s.Grade)
		if errors.Is(err, io.EOF) {
			break
		}
		if err := storage.Put(&s); err != nil {
			fmt.Println("error: ", err.Error())
			return
		}
	}

	fmt.Println("Студенты из хранилища:")
	for _, student := range storage.GetAll() {
		fmt.Println(student.Name, student.Age, student.Grade)
	}
	fmt.Println()
	str := ""
	for {
		fmt.Println("Введите имя студента:")
		fmt.Scan(&str)
		fmt.Println(storage.Get(str))
		fmt.Println(storage.Get(str).String())
	}

}
