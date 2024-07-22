package main

// User interface
type User interface {
	GetType() string
}

// User struct
type UserImpl struct{}

func (u UserImpl) GetType() string {
	return "user"
}
