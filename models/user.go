package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func Query(N string, P string) (user User, ok bool) {
	for _, u := range GetAllUsers() {
		if u.Username == N {
			user = u
			ok = user.CheckPassword(P)
			return
		}
	}
	return User{}, false
}

func QueryUserById(id interface{}) (user User) {
	for _, u := range GetAllUsers() {
		if u.ID == id {
			user = u
			return
		}
	}
	return User{}
}
