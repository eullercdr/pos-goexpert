package main

import "fmt"

func main() {
	salarios := map[string]int{"jose": 1200, "maria": 1500, "pedro": 800}
	fmt.Println(salarios["jose"])
	delete(salarios, "jose")
	salarios["wesley"] = 2000
	fmt.Println(salarios)

	sal := make(map[string]int)
	sal["jose"] = 1200
	sal["maria"] = 1500
	sal["pedro"] = 800
	fmt.Println(sal)

	for nome, salario := range salarios {
		fmt.Printf("Nome: %s, Salario: %d\n", nome, salario)
	}
}
