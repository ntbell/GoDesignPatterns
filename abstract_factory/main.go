package main

import "fmt"

func main() {
	adminFactory, _ := GetUserFactory("admin")
	customerFactory, _ := GetUserFactory("customer")

	customerConfig := customerFactory.makeConfig()
	customerPermissions := customerFactory.makePermissions()

	adminConfig := adminFactory.makeConfig()
	adminPermissions := adminFactory.makePermissions()

	fmt.Println(adminConfig.getName())
	fmt.Println(adminConfig.getID())
	fmt.Println(adminPermissions.getScope())
	fmt.Println(adminPermissions.getScope())

	fmt.Println(customerConfig.getName())
	fmt.Println(customerConfig.getID())
	fmt.Println(customerPermissions.getScope())
	fmt.Println(customerPermissions.getScope())
}
