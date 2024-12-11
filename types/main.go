package main

import "fmt"

type ID int

var (
	b bool = true
	c ID   = 1
	e int  = 1
)

func main() {
	fmt.Printf("O tipo de C é %T", c)
}
