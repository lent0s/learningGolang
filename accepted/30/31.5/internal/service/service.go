package service

import (
	"fmt"
	repo "sb30_5/internal/repository/database"
)

type Service struct {
	Name     string
	Age      int
	SourceId int
	TargetId int
	NewAge   int
}

type Api struct {
	list map[string]interface{}
	db   *repo.Storage
}

func CreateUser(req map[string]interface{}, db *repo.Storage) (int,
	string, error) {

	srv, err := convReqToService(req)
	if err != nil {
		return 0, "", err
	}

	user := repo.User{}
	user.Name = srv.Name
	user.Age = srv.Age
	if user.Name, err = nameCheck(user.Name); err != nil {
		return 0, "", err
	}

	id, err := db.CreateUser(&user)
	if err != nil {
		return 0, "", err
	}
	return id, user.Name, nil
}

func MakeFriends(req map[string]interface{}, db *repo.Storage) (user1,
	user2 *repo.User, err error) {

	srv, err := convReqToService(req)
	if err != nil {
		return nil, nil, err
	}
	if srv.SourceId == srv.TargetId {
		return nil, nil, fmt.Errorf("ids match")
	}

	user1, err = db.ReadUser(srv.SourceId)
	if err != nil {
		return nil, nil, err
	}
	user2, err = db.ReadUser(srv.TargetId)
	if err != nil {
		return nil, nil, err
	}
	if alreadyFriends(user1.Friends, user2.UserId) {
		return nil, nil, fmt.Errorf("%s and %s are "+
			"already friends %s", user1.Name, user2.Name, string('\U0001f37b'))
	}

	user1.Friends = appendFriend(user1.Friends, user2.UserId)
	user2.Friends = appendFriend(user2.Friends, user1.UserId)
	if db.UpdateUser(user1) != nil {
		return nil, nil, err
	}
	if db.UpdateUser(user2) != nil {
		return nil, nil, err
	}
	return user1, user2, nil
}

func DeleteUser(req map[string]interface{}, db *repo.Storage) error {

	srv, err := convReqToService(req)
	if err != nil {
		return err
	}
	user, err := db.ReadUser(srv.TargetId)
	if err != nil {
		return err
	}

	for _, friendID := range user.Friends {
		friend, _ := db.ReadUser(friendID)
		deleteFriend(&friend.Friends, user.UserId)
		if err = db.UpdateUser(friend); err != nil {
			return err
		}
	}

	if err = db.DeleteUser(user.UserId); err != nil {
		return err
	}
	return nil
}

func GetFriends(id int, db *repo.Storage) (name, list string, err error) {

	user, err := db.ReadUser(id)
	if err != nil {
		return "", "", err
	}

	if len(user.Friends) == 0 {
		return user.Name, "has no friends", nil
	}

	list = ""
	for _, friendID := range user.Friends {
		friend, _ := db.ReadUser(friendID)
		list += friend.Stringer() + "\n"
	}

	return user.Name, list, nil
}

func UpdateUser(req map[string]interface{}, db *repo.Storage) error {

	srv, err := convReqToService(req)
	if err != nil {
		return err
	}
	user, err := db.ReadUser(srv.TargetId)
	if err != nil {
		return err
	}

	if user.Age == srv.NewAge {
		return nil
	}
	user.Age = srv.NewAge

	if err = db.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func GetAllUsers(db *repo.Storage) string {

	return db.GetAllUsers()
}
