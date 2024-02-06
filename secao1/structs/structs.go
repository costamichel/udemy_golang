package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {

	var u1 usuario
	u1.nome = "Michel"
	u1.idade = 31
	u1.endereco = endereco{"Nonoai", 10}
	fmt.Println(u1)

	u2 := usuario{"Neymar", 32, endereco{"Barcelona", 11}}
	fmt.Println(u2)

	u3 := usuario{idade: 17}
	u3.endereco.logradouro = "Madri"
	fmt.Println(u3)

}
