package main

import (
	"fmt"
)

func getEmployeeData(e employee) {
	fmt.Printf("The employee %s earns %v€ per year.\n", e.getName(), e.getSalary())
}

type employee interface {
	getName() string
	getSalary() int
}
type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

func test02(e employee) {
	getEmployeeData(e)
	fmt.Println("========================================")
}

func main() {
	test02(contractor{
		name:         "Klaus",
		hourlyPay:    30,
		hoursPerYear: 1500,
	})
}
