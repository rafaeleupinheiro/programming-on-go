package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Olá Go")

	x := "Rafael Pinheiro"

	fmt.Println(x)
	fmt.Println(x[4:8])

	y := "PEDIDO_XXXXXXXXXX_YYYYYYYYYYYYYY_ZZZ.PED"
	y = y[0 : len(y)-4]
	fmt.Println(y[7:])

	var numerof float32
	numerof = 2154.45965

	valor := fmt.Sprintf("%012d", int32(numerof*100))

	fmt.Println(valor)
	fmt.Println(len(valor))

	data := time.Now()

	fmt.Println(data)
	fmt.Println(data.Format("304"))
	fmt.Println(data.Format("02012006"))

	fmt.Println(time.Now().Format("02012006"))

	fmt.Println(time.Now().Format("150405"))

	x = "OLPEDxxxxxDDMMYYYYhhnnssNNNNNNNNNNNN.txt"

	fmt.Println("|" + x[:len(x)-3] + "|")
	j := fmt.Sprint(x[:len(x)-3], "tmp")
	fmt.Println(j)

	fmt.Println(fmt.Sprintf("|%-13d|", 789600471258))

	fmt.Println(fmt.Sprint("OLPEDxxxxxDDMMYYYYhhnnssNNNNNNNNNNNN.", "sdasda"))

	var qtd int

	fmt.Println(qtd+1)

	var v int32
	v = 99999999
	fmt.Println(v)

}
