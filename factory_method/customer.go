package main

// Concrete product
type customer struct {
	User
}

func newCustomer() IUser {
	return &customer{
		User: User{
			ID:   1,
			Name: "John Doe",
		},
	}
}
