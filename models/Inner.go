package models

type Inner struct {
	Id int32
}

var curInnerId int32

func getInnerId() int32 {
	curInnerId += 1
	return curInnerId
}

func NewInner() *Inner {
	return &Inner{getInnerId()}
}