package main

// Abstract product
type IShoe interface {
	getLogo() string
	getSize() int
	setLogo(logo string)
	setSize(size int)
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) getSize() int {
	return s.size
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}
