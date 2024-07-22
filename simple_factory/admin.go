package main

type Admin struct{}

func (a Admin) getName() string {
	return "admin"
}

func (a Admin) getID() int {
	return 2
}
