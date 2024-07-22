package main

import "fmt"

func getUser(userType string) (IUser, error) {
	switch userType {
	case "admin":
		return newAdmin(), nil
	case "customer":
		return newCustomer(), nil
	case "user":
		return newUser(), nil
	default:
		return nil, fmt.Errorf("wrong user type passed")
	}
}
