package models

import (
	"api-server/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// GetAllUsers fetches all user data.
func GetAllUsers(user *[]User) (err error) {
	err = config.DB.Find(user).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateUser inserts a new user.
func CreateUser(user *User) (err error) {
	err = config.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUserByID fetches a user by ID.
func GetUserByID(user *User, id string) (err error) {
	err = config.DB.Where("id = ?", id).First(user).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser updates a user
func UpdateUser(user *User, id string) (err error) {
	config.DB.Save(user)
	fmt.Println(user)
	return nil
}

// DeleteUser deletes a user
func DeleteUser(user *User, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}
