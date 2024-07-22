package main

// Concrete product
type User struct {
	ID   int
	Name string
}

func (u *User) setName(name string) {
	u.Name = name
}

func (u *User) getName() string {
	return u.Name
}

func (u *User) setID(id int) {
	u.ID = id
}

func (u *User) getID() int {
	return u.ID
}
