package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	//last value
	fmt.Println(meuArray[len(meuArray)-1])

	//count of elements
	fmt.Println(len(meuArray))

	//first element
	fmt.Println(meuArray[0])

	//for
	for i, v := range meuArray {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}
