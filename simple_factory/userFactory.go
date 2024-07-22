package main

import "errors"

type UserFactory struct{}

func (uf UserFactory) Create(userType string) (IUser, error) {
	switch userType {
	case "admin":
		return Admin{}, nil
	case "customer":
		return Customer{}, nil
	case "user":
		return User{}, nil
	default:
		return nil, errors.New("wrong user type passed")
	}
}
