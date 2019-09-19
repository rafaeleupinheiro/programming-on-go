package main

import (
	"fmt"
	"time"
)

func main() {
	Cronometrar(GerarFibonacci(8))
	//Cronometrar(GerarFibonacci(48))
	//Cronometrar(GerarFibonacci(88))
}

func GerarFibonacci(n int) func() {
	return func() {
		a, b := 0, 1

		fib := func() int {
			a, b = b, a+b
			return a
		}

		for i := 0; i < n; i++ {
			fmt.Printf("%d ", fib())
		}
		fmt.Println()
	}
}

func Cronometrar(funcao func()) {
	t := time.Now()
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	fmt.Println(t)
	//funcao()
	//fmt.Printf("\nTempo de execução: %s\n\n", time.Since(inicio))
}
