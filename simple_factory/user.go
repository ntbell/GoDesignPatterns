package main

type User interface {
	GetType() string
}

type UserImpl struct{}

func (u UserImpl) GetType() string {
	return "user"
}
