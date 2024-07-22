package main

// Concrete product
type Person struct {
	ID   int
	Name string
}

func (p *Person) setName(name string) {
	p.Name = name
}

func (p *Person) getName() string {
	return p.Name
}

func (p *Person) setID(id int) {
	p.ID = id
}

func (p *Person) getID() int {
	return p.ID
}
