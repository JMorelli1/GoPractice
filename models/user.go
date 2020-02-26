package models

import (
	"errors"
	"fmt"
)

//User struct
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(user User) (User, error) {

	if user.ID != 0 {
		return User{}, errors.New("New User must not include id")
	}
	user.ID += nextID
	users = append(users, &user)
	return user, nil
}

func GetUserByID(id int) (User, error) {
	for _, v := range users {
		if v.ID == id {
			return *v, nil
		}
	}
	return User{}, fmt.Errorf("No user found with ID '%v'", id)
}

func UpdateUser(user User) (User, error) {
	for i, value := range users {
		if value.ID == user.ID {
			users[i] = &user
			return user, nil
		}
	}
	return User{}, fmt.Errorf("No user found at '%v'", user.ID)
}

func RemoveUser(id int) error {
	for i, value := range users {
		if value.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("No user found at '%v'", id)
}
