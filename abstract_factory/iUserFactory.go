package main

import "fmt"

// Abstract factory interface
type IUserFactory interface {
	makePermissions() IPermissions
	makeConfig() IConfig
}

func GetUserFactory(userType string) (IUserFactory, error) {
	if userType == "admin" {
		return &Admin{}, nil
	}
	if userType == "customer" {
		return &Customer{}, nil
	}
	return nil, fmt.Errorf("wrong user type passed")
}
