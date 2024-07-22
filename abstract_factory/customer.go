package main

// Concrete factory
type Customer struct{}

func (c Customer) makePermissions() IPermissions {
	return &CustomerPermissions{
		Permissions: Permissions{
			Scope: "customer",
		},
	}
}

func (c Customer) makeConfig() IConfig {
	return &CustomerConfig{
		Config: Config{
			Name: "Jane Doe",
			ID:   1,
		},
	}
}
