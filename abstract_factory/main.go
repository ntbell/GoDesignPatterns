package main

import "fmt"

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	fmt.Print(adidasShoe.getLogo())
	fmt.Print(adidasShirt.getLogo())
	fmt.Printf("%d", nikeShoe.getSize())
	fmt.Printf("%d", nikeShirt.getSize())
}
