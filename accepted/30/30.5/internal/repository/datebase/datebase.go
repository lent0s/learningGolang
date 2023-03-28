package datebase

import (
	"fmt"
)

var userId = 0

type User struct {
	UserId  int
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

func (u *User) Stringer() string {

	return fmt.Sprintf("id%d\tName: %-15s\tAge: %4d\tFriends: %v",
		u.UserId, u.Name, u.Age, u.Friends)
}

type Storage struct {
	UserId map[int]*User
}

func NewDatabase() *Storage {

	return &Storage{
		make(map[int]*User),
	}
}

func (s *Storage) CreateUser(u *User) int {

	userId++
	u.UserId = userId
	s.UserId[u.UserId] = u
	return userId
}

func (s *Storage) ReadUser(id int) (*User, error) {

	if s.UserId[id] == nil {
		return nil, fmt.Errorf("user id%d does not exist", id)
	}
	return s.UserId[id], nil
}

func (s *Storage) UpdateUser(u *User) error {

	if _, err := s.ReadUser(u.UserId); err != nil {
		return err
	}
	s.UserId[u.UserId] = u
	return nil
}

func (s *Storage) DeleteUser(id int) error {

	if _, err := s.ReadUser(id); err != nil {
		return err
	}
	delete(s.UserId, id)
	return nil
}

/*
Уровень репозитория (repository) в основном отвечает за работу с сервисами
хранения данных - БД и кэшом. Функции этого уровня ответственны за
создание/обновление/удаление/получение данных из хранилищ, обработку ошибок,
получаемых от хранилищ и отдачу содержимого хранилища наверх - в сервис уровень
(в основном - в виде структуры). Уровень репозитория должен быть максимально
тупым. Все что он делает - это манипулирует данными в хранилище.
*/
