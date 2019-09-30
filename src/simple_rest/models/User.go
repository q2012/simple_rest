package models

var curUserId int32

type User struct {
	Id int32
	Username string
	Password string
}

func NewUser(uname string, pass string) *User {
	return &User{getUserId(), uname, pass}
}

func getUserId() int32 {
	curUserId += 1
	return curUserId
}