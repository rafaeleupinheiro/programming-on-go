package main

import (
	"fmt"
	"time"
)

type Data1 struct {
	d1 string
}

func main() {
	var t time.Time

	t = time.Date(2019, time.February, 06, 14, 43, 0, 0, time.UTC)
	fmt.Println(t)

	tm, _ := time.Parse(time.RFC3339, "2019-02-06T16:00:00Z")
	fmt.Println(tm)

	fmt.Println(tm.Format(time.RFC3339))

	if Tempo().IsZero() {
		fmt.Println("É zero")
	} else {
		fmt.Println("Negado")
	}

	if (Data() == Data1{}) {
		fmt.Println("É zero")
	} else {
		fmt.Println("Negado")
	}
}

func Tempo() (time.Time) {
	return time.Time{}
}

func Data() (Data1) {
	return Data1{"aaaaaaa"}
}
