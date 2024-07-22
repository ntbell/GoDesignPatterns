package main

type User struct{}

func (u User) getName() string {
	return "user"
}

func (u User) getID() int {
	return 0
}
