package main

func main() {
	programmer, _ := getEmployee("programmer")
	accountant, _ := getEmployee("accountant")

	programmer.setID(3)
	accountant.setName("New name")
}
