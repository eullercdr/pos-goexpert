package main

func soma(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

func main() {
	a := 10
	b := 20
	c := soma(&a, &b)
	println(c)
	println(a)
	println(b)
}
