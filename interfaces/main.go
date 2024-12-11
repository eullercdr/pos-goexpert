package main

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

type Customer struct {
	Name     string
	Age      int
	Email    string
	IsActive bool
	Address
}

type Pessoa interface {
	Deactivate()
}

func (c *Customer) Deactivate() {
	c.IsActive = false
}

func Desativacao(pessoa Pessoa) {
	pessoa.Deactivate()
}

func main() {

}
