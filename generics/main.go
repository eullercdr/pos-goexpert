package main

type MyNumber int

type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compare[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	sal := map[string]int{
		"jose":  100,
		"maria": 200,
		"pedro": 300,
	}
	println(Soma(sal))
	sal2 := map[string]float64{
		"jose":  100.20,
		"maria": 200.30,
		"pedro": 300.30,
	}
	println(Soma(sal2))
	sal3 := map[string]MyNumber{
		"jose":  100,
		"maria": 200,
		"pedro": 300,
	}
	println(Soma(sal3))
	println(Compare(10, 11.0))
}
