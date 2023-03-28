package database

import (
	"fmt"
)

type User struct {
	UserId  int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

func (u *User) Stringer() string {

	return fmt.Sprintf("id%d\tName: %-15s\tAge: %4d\tFriends: %v",
		u.UserId, u.Name, u.Age, u.Friends)
}

// get new unique User_ID
func newUserID(s *Storage) (int, error) {

	if err := readDb(s); err != nil {
		return 0, err
	}
	return len(s.users) + 1, nil
}

func (s *Storage) CreateUser(u *User) (int, error) {

	var err error
	u.UserId, err = newUserID(s)
	if err != nil {
		return 0, err
	}

	data, err := makeData(u)
	if err != nil {
		return 0, err
	}
	if err = appendDb(s.fileDB, data); err != nil {
		return 0, err
	}
	return u.UserId, nil
}

func (s *Storage) ReadUser(id int) (*User, error) {

	if err := readDb(s); err != nil {
		return nil, err
	}
	if s.users[id] == nil {
		return nil, fmt.Errorf("user id%d does not exist", id)
	}
	return s.users[id], nil
}

func (s *Storage) UpdateUser(u *User) error {

	if _, err := s.ReadUser(u.UserId); err != nil {
		return err
	}
	data, err := makeData(u)
	if err != nil {
		return err
	}
	if err := updateDb(s.fileDB, data, u.UserId); err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteUser(id int) error {

	user, err := s.ReadUser(id)
	if err != nil {
		return err
	}

	user.Name = ""
	user.Age = 0
	user.Friends = nil

	if s.UpdateUser(user) != nil {
		return err
	}
	return nil
}

func (s *Storage) GetAllUsers() string {

	err := readDb(s)
	if err != nil {
		return fmt.Sprintf("internal error, try later: %v", err)
	}
	response := ""

	for id := 1; id < len(s.users)+1; id++ {
		if s.users[id] == nil {
			continue
		}
		response += s.users[id].Stringer() + "\n"
	}

	if response == "" {
		return "no users yet"
	}
	return response
}
