package main

type Conta struct {
	saldo int
}

func NewConta(saldo int) *Conta {
	return &Conta{saldo: 0}
}

func (c Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func (c *Conta) depositar(valor int) int {
	c.saldo += valor
	return c.saldo
}

func main() {
	conta := Conta{saldo: 100}
	conta.simular(200)
	println(conta.saldo)
	conta.depositar(200)
	println(conta.saldo)
}
