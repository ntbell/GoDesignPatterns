package main

// Product interface
type IEmployee interface {
	getName() string
	getID() int
	setName(name string)
	setID(id int)
}
