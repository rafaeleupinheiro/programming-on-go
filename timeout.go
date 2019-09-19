package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool, 1)
	go executar(c)

	fmt.Println("Esperando...")

	fim := false
	for !fim {
		select {
		case fim = <-c:
			fmt.Println("Fim!")
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout!")
			fim = true
		}
	}
}

func executar(c chan<- bool) {
	time.Sleep(1 * time.Second)
	c <- true
}
