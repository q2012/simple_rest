package models

type Session struct {
	Key string
	UserId int
}

func NewSession(key string, uid int) *Session {
	return &Session{key, uid}
}
