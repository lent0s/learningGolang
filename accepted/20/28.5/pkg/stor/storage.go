package stor

import (
	"fmt"
	"gitHomeWork/pkg/stud"
)

type Storage struct {
	Class map[string]*stud.Student
}

func NewStorage() *Storage {
	return &Storage{
		Class: make(map[string]*stud.Student),
	}
}

func (db *Storage) Put(student *stud.Student) {
	db.Class[student.Name] = student
}

func (db *Storage) Get(name string) (*stud.Student, error) {
	student, err := db.Class[name]
	if !err {
		return nil, fmt.Errorf("no data available for %v\n", name)
	}
	return student, nil
}
