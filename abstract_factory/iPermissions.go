package main

// Abstract product
type IPermissions interface {
	getScope() string
	setScope(Scope string)
}

type Permissions struct {
	Scope string
}

func (p *Permissions) getScope() string {
	return p.Scope
}

func (p *Permissions) setScope(Scope string) {
	p.Scope = Scope
}
