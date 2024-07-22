package main

import "fmt"

func main() {
	adminFactory, _ := GetUserFactory("admin")
	customerFactory, _ := GetUserFactory("customer")

	customerConfig := customerFactory.makeConfig()
	customerPermissions := customerFactory.makePermissions()

	adminConfig := adminFactory.makeConfig()
	adminPermissions := adminFactory.makePermissions()

	fmt.Print(adminConfig.getName())
	fmt.Print(adminConfig.getID())
	fmt.Print(adminPermissions.getScope())
	fmt.Print(adminPermissions.getScope())

	fmt.Print(customerConfig.getName())
	fmt.Print(customerConfig.getID())
	fmt.Print(customerPermissions.getScope())
	fmt.Print(customerPermissions.getScope())
}
