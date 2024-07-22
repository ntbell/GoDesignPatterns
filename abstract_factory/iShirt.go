package main

// Abstract product
type IShirt interface {
	getLogo() string
	getSize() int
	setLogo(logo string)
	setSize(size int)
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) getSize() int {
	return s.size
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) setSize(size int) {
	s.size = size
}