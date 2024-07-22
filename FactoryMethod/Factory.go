package main

import "fmt"

// Employee interface
type Employee interface {
	PaySalary()
	Dismiss()
}

// Programmer struct
type Programmer struct {
	ID int
}

func (p Programmer) PaySalary() {
	fmt.Printf("Paying salary to programmer %d\n", p.ID)
}

func (p Programmer) Dismiss() {
	fmt.Printf("Dismissing programmer %d\n", p.ID)
}

// Accountant struct
type Accountant struct {
	ID int
}

func (a Accountant) PaySalary() {
	fmt.Printf("Paying salary to accountant %d\n", a.ID)
}

func (a Accountant) Dismiss() {
	fmt.Printf("Dismissing accountant %d\n", a.ID)
}

// Department interface
type Department interface {
	CreateEmployee(id int) Employee
	Fire(id int)
}

// BaseDepartment struct
type BaseDepartment struct{}

func (bd BaseDepartment) Fire(d Department, id int) {
	employee := d.CreateEmployee(id)
	employee.PaySalary()
	employee.Dismiss()
}

// ITDepartment struct
type ITDepartment struct {
	BaseDepartment
}

func (it ITDepartment) CreateEmployee(id int) Employee {
	return Programmer{ID: id}
}

// AccountingDepartment struct
type AccountingDepartment struct {
	BaseDepartment
}

func (ad AccountingDepartment) CreateEmployee(id int) Employee {
	return Accountant{ID: id}
}

func main() {
	itDept := ITDepartment{}
	accDept := AccountingDepartment{}

	itDept.Fire(itDept, 1)
	accDept.Fire(accDept, 2)
}
