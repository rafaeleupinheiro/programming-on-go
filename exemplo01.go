package main

import "fmt"

type Pessoa struct {
	nome string
}

func main() {
	var p Pessoa
	 p.nome = "Rafael"
	fmt.Println(p.nome)

	p.mudaNome()
	fmt.Println(p.nome)

	mudaNome(&p)
	fmt.Println(p.nome)
	c := fmt.Sprintf(p.nome)
	fmt.Println(c)
}

func (p Pessoa) mudaNome() {
	p.nome = "Armando"
}

func mudaNome(p *Pessoa) {
	p.nome = "Armando"
}
