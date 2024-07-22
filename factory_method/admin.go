package main

// Concrete product
type admin struct {
	User
}

func newAdmin() IUser {
	return &admin{
		User: User{
			ID:   2,
			Name: "Jane Doe",
		},
	}
}
