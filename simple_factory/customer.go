package main

type Customer struct{}

func (c Customer) getName() string {
	return "customer"
}

func (c Customer) getID() int {
	return 1
}
