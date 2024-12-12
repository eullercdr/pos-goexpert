package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "hello word"
	showTypes(x)
	showTypes(y)
}

func showTypes(t interface{}) {
	fmt.Printf("Type: %T e value %v\n", t, t)
}
