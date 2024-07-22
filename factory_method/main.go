package main

import "fmt"

func main() {
	admin, _ := GetUser("admin")
	customer, _ := GetUser("customer")
	user, _ := GetUser("user")

	admin.setID(3)
	customer.setID(4)
	user.setID(5)

	admin.setName("New Administrator Name")
	customer.setName("New Customer Name")
	user.setName("New User Name")

	fmt.Println(admin.getName())
	fmt.Println(admin.getID())
	fmt.Println(customer.getName())
	fmt.Println(customer.getID())
	fmt.Println(user.getName())
	fmt.Println(user.getID())
}
