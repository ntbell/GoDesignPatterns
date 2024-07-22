package main

type Customer struct{}

func (c Customer) GetType() string {
	return "customer"
}
