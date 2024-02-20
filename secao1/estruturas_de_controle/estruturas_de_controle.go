package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	cor := "VERDE"

	if cor == "VERDE" {
		fmt.Println("É verde")
	} else if cor == "AMARELA" {
		fmt.Println("É Amarela")
	} else {
		fmt.Println("Não é nem verde e nem amarelo")
	}

	if outraCor := cor; outraCor[0] == 'A' { // If init -> variável outraCor não pode ser usada fora deste escopo
		fmt.Println("O nome da cor começa com A")
	} else {
		fmt.Println("O nome da cor não começa com A")
	}
}
