package main

import "errors"

type UserFactory struct{}

func (uf UserFactory) Create(userType string) (User, error) {
	switch userType {
	case "user":
		return UserImpl{}, nil
	case "customer":
		return Customer{}, nil
	case "admin":
		return Admin{}, nil
	default:
		return nil, errors.New("wrong user type passed")
	}
}
