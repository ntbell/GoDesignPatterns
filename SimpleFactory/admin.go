package main

// Admin struct
type Admin struct{}

func (a Admin) GetType() string {
	return "admin"
}
