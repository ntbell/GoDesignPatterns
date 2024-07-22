package main

// Product interface
type IUser interface {
	getName() string
	getID() int
	setName(name string)
	setID(id int)
}
