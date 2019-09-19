package main

import "fmt"

func main() {
	estados := make(map[string]Estado, 6)
	estados["GO"] = Estado{"Goiás", 6434052, "Goiânia"}
	estados["PB"] = Estado{"Paraíba", 3914418, "João Pessoa"}
	estados["PR"] = Estado{"Paraná", 10997462, "Curitiba"}
	estados["RN"] = Estado{"Rio Grande do Norte", 3373960, "Natal"}
	estados["AM"] = Estado{"Amazonas", 3807923, "Manaus"}
	estados["SE"] = Estado{"Sergipe", 2228489, "Aracaju"}

	fmt.Println(estados)

	sergipe := estados["SE"]
	fmt.Println(sergipe)

	saoPaulo, encontrado := estados["SP"]
	if encontrado {
		fmt.Println(saoPaulo)
	} else {
		fmt.Println("Não encontrado!")
	}

	delete(estados, "AM")
	fmt.Println(estados)

	for sigla, estado := range estados {
		fmt.Printf("%s (%s) possui %d habitantes. \n", estado.nome, sigla, estado.populacao)
	}
}

type Estado struct {
	nome string
	populacao int
	capital string
}
