package main

import (
	"fmt"
	"gitHomeWork/pkg/stor"
	"gitHomeWork/pkg/stud"
	"io"
)

func main() {
	storage := stor.NewStorage()
	for {
		student := stud.NewStudent()
		fmt.Print("\nEnter new information (name age grade)\n" +
			"(press CTRL+D/CTRL+Z to exit entering): ")
		_, err := fmt.Scan(&student.Name, &student.Age, &student.Grade)
		if err == io.EOF {
			break
		}
		storage.Put(student)
	}

	printStudent := func(s *stud.Student, err error) {
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Name: %v\tAge: %v\tGrade: %v\n", s.Name, s.Age, s.Grade)
	}
	for _, student := range storage.Class {
		printStudent(storage.Get(student.Name))
	}
}
