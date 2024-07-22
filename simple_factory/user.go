package main

type User interface {
	getName() string
	getID() int
}

type UserImpl struct{}

func (u UserImpl) getName() string {
	return "user"
}

func (u UserImpl) getID() int {
	return 0
}
