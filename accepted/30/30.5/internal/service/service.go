package service

import (
	"fmt"
	repo "sb30_5/internal/repository/datebase"
)

var db = repo.NewDatabase()

type Service struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	SourceId int    `json:"source_id"`
	TargetId int    `json:"target_id"`
	NewAge   int    `json:"new age"`
}

func CreateUser(req *Service) (int, string, error) {

	user := repo.User{
		Name: req.Name,
		Age:  req.Age,
	}
	if user.Age < 0 {
		return 0, "", fmt.Errorf("user age is incorrect")
	}
	var err error
	user.Name, err = nameCheck(user.Name)
	if err != nil {
		return 0, "", err
	}
	id := db.CreateUser(&user)
	return id, user.Name, nil
}

func MakeFriends(req *Service) (user1, user2 *repo.User, err error) {

	if req.SourceId == req.TargetId {
		return nil, nil, fmt.Errorf("ids match")
	}
	user1, err = db.ReadUser(req.SourceId)
	if err != nil {
		return nil, nil, err
	}
	user2, err = db.ReadUser(req.TargetId)
	if err != nil {
		return nil, nil, err
	}
	if alreadyFriends(user1.Friends, user2.UserId) {
		return nil, nil, fmt.Errorf("%s and %s are "+
			"already friends %s", user1.Name, user2.Name, string('\U0001f37b'))
	}
	user1.Friends = append(user1.Friends, user2.UserId)
	user2.Friends = append(user2.Friends, user1.UserId)
	return user1, user2, nil
}

func DeleteUser(req *Service) error {

	user, err := db.ReadUser(req.TargetId)
	if err != nil {
		return err
	}
	for _, friend := range user.Friends {
		deleteFriend(&db.UserId[friend].Friends, user.UserId)
	}
	if err = db.DeleteUser(user.UserId); err != nil {
		return err
	}
	return nil
}

func GetFriends(id int) (name, list string, err error) {

	user, err := db.ReadUser(id)
	if err != nil {
		return "", "", err
	}
	list = ""
	for _, friend := range user.Friends {
		list += db.UserId[friend].Stringer() + "\n"
	}
	if len(user.Friends) == 0 {
		list += "has no friends"
	}
	return user.Name, list, nil
}

func UpdateUser(req *Service) error {

	user, err := db.ReadUser(req.TargetId)
	if err != nil {
		return err
	}
	if req.Age < 0 {
		return fmt.Errorf("user age is incorrect")
	}
	user.Age = req.NewAge
	if err = db.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func GetAllUsers() string {

	response := ""
	for _, user := range db.UserId {
		response += user.Stringer() + "\n"
	}
	if response == "" {
		response = "no users yet"
	}
	return response
}

/*
Уровень бизнес-логики (service) отвечает за реализацию бизнес логики приложения,
работу с данными, полученными с уровня репозитория, и отдачу полученного
результата на уровень api или на тот же уровень сервиса. В нем содержится
вся логика работы сервиса: проверки; обработка ошибок; валидация данных,
полученных из хранилищ; походы в другие сервисы и т.д.
*/
