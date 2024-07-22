package main

type Admin struct{}

func (a Admin) GetType() string {
	return "admin"
}
