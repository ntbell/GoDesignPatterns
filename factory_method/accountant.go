package main

// Concrete product
type accountant struct {
	Person
}

func newAccountant() IEmployee {
	return &accountant{
		Person: Person{
			ID:   1,
			Name: "John Doe",
		},
	}
}
