package main

// Concrete factory
type Admin struct{}

func (a *Admin) makePermissions() IPermissions {
	return &AdminPermissions{
		Permissions: Permissions{
			Scope: "admin",
		},
	}
}

func (a *Admin) makeConfig() IConfig {
	return &AdminConfig{
		Config: Config{
			Name: "John Doe",
			ID:   0,
		},
	}
}
