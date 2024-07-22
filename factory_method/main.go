package main

func main() {
	admin, _ := getUser("admin")
	customer, _ := getUser("customer")

	admin.setID(3)
	customer.setID(4)

	admin.setName("Administrator")
	customer.setName("New Customer Name")
}
