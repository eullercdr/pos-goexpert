package main

import "errors"

func main() {
	valor, err := Sum(100, 2)
	if err != nil {
		panic(err)
	}
	println(valor)
}

// func Sum(a, b int) (int, bool) {
// 	if a+b >= 50 {
// 		return a + b, true
// 	}
// 	return a + b, false
// }

func Sum(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("The sum is greater than 50")
	}
	return a + b, nil
}
