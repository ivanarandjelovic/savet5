package user

import (
	//	"fmt"
	. "aivan/savet5/db"
	"log"
	"errors"
)

// User object
type User struct {
	Id          uint64 `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Description string `json:"description"`
	Role        string `json:"role"`
}

func (u User) TableName() string {
	return "Users"
}

type UserError struct {
}

//Try to find the user based on email and pass, return the user if found, nil otherwise
func Login(email string, password string) (User, error) {
	var user User
	var totalCount uint32
	//DB.Where("email = ? and password = ?", email, password).First(&user)
	if DB.First(&user).RecordNotFound() {
		return user, errors.New("User not found!")
	}
	//	if user == nil {
	log.Println("user not found! user: ", user, user.Email, user.Password, totalCount)
	//} else {
	//	log.Println("user FOUND! user: ", user)
	//}
	return user, nil
}
