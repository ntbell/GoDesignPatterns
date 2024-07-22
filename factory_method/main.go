package main

func main() {
	admin, _ := getUser("admin")
	customer, _ := getUser("customer")
	user, _ := getUser("user")

	admin.setID(3)
	customer.setID(4)
	user.setID(5)

	admin.setName("New Administrator Name")
	customer.setName("New Customer Name")
	user.setName("New User Name")
}
