package main

// Abstract product
type IConfig interface {
	getName() string
	setName(name string)
	getID() int
	setID(id int)
}

type Config struct {
	Name string
	ID   int
}

func (c *Config) getName() string {
	return c.Name
}

func (c *Config) getID() int {
	return c.ID
}

func (c *Config) setName(Name string) {
	c.Name = Name
}

func (c *Config) setID(ID int) {
	c.ID = ID
}
