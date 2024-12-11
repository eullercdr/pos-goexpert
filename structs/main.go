package main

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

type Customer struct {
	Name     string
	Age      int
	Email    string
	IsActive bool
	Address
}

func (c *Customer) Deactivate() {
	c.IsActive = false
}

func main() {
	customer := Customer{
		Name:     "John Doe",
		Age:      30,
		Email:    "jupom@hukusti.cv",
		IsActive: true,
	}
	customer.Deactivate()
	fmt.Printf("%+v\n", customer)
}
