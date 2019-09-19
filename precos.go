package main

import "fmt"

func main() {
	precoDolar, precoReal := PrecoFinal(34.99)

	fmt.Printf("Preco final em dólar: %.2f\n" + "Preço final em reais: %.2f\n", precoDolar, precoReal)
}

func PrecoFinal(precoCusto float64) (precoDolar float64, precoReal float64) {
	fatorLucro := 1.33
	taxaConversao := 2.34

	precoDolar = precoCusto * fatorLucro
	precoReal = precoDolar * taxaConversao

	return precoDolar, precoReal
}
