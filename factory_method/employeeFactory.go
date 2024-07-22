package main

import "fmt"

func getEmployee(employeeType string) (IEmployee, error) {
	switch employeeType {
	case "programmer":
		return newProgrammer(), nil
	case "accountant":
		return newAccountant(), nil
	default:
		return nil, fmt.Errorf("wrong employee type passed")
	}
}
