package main

import "fmt"

func main() {
	factory := UserFactory{}

	admin, err := factory.Create("admin")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Created:", admin.getName())
	}

	customer, err := factory.Create("customer")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Created:", customer.getName())
	}

	user, err := factory.Create("user")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Created:", user.getName())
	}
}
