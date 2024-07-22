package main

// Concrete product
type programmer struct {
	Person
}

func newProgrammer() IEmployee {
	return &programmer{
		Person: Person{
			ID:   2,
			Name: "Jane Doe",
		},
	}
}
