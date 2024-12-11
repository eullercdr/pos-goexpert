package main

func main() {
	total := func() int {
		return sum(1, 2, 3, 4, 5) * 2
	}()
	println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}
