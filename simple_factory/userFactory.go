package main

import "errors"

type UserFactory struct{}

func (uf UserFactory) Create(userType string) (User, error) {
	switch userType {
	case "admin":
		return Admin{}, nil
	case "customer":
		return Customer{}, nil
	case "user":
		return UserImpl{}, nil
	default:
		return nil, errors.New("wrong user type passed")
	}
}
