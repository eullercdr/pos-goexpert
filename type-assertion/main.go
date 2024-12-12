package main

func main() {
	var minhaVar interface{} = "Teste"

	println(minhaVar.(string))
	resultado, ok := minhaVar.(int)
	if !ok {
		println("Erro ao converter")
	} else {
		println(resultado)
	}
}
