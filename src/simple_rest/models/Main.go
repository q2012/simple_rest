package models

var curMainId int32

type Main struct {
		Id int32
		In *Inner
}

func NewMain(in *Inner) *Main {
	return &Main{getMainId(), in}
}

func getMainId() int32 {
	curMainId += 1
	return curMainId
}