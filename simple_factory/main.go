package main

import "fmt"

func main() {
	factory := UserFactory{}

	user, err := factory.Create("user")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Created:", user.GetType())
	}

	customer, err := factory.Create("customer")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Created:", customer.GetType())
	}

	admin, err := factory.Create("admin")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Created:", admin.GetType())
	}
}
