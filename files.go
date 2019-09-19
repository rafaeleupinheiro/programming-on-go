package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const teste = "./teste/teste1/21141904.PED"

func main() {
	arquivo := filepath.Base(teste)

	parent := filepath.Dir(filepath.Dir(teste))

	backup := fmt.Sprint(parent, "/BkpRetornos/", time.Now().Format("20060102"))

	if !ExisteArquivo(backup) {
		fmt.Println("Criou")
		CriaDiretorio(backup)
	}

	backup = fmt.Sprint(backup, "/", arquivo)

	input, err := ioutil.ReadFile(teste)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(backup, input, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CriaDiretorio(directory string) error {
	err := os.MkdirAll(directory, os.ModePerm)
	return err
}

func ExisteArquivo(diretorio string) bool {
	if _, err := os.Stat(diretorio); err == nil {
		if err != nil {
			return false
		}
		return true
	}
	return false
}
