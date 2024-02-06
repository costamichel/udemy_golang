package main

import "fmt"

/* Não há herança "nativa" no Go, mas pode-se usar structs para imitar o conceito */

type pessoa struct {
	nome  string
	idade uint8
}

type estudante struct {
	pessoa
	curso string
}

func main() {

	var e1 estudante
	e1.nome = "Greg Shaw"
	e1.idade = 55
	e1.curso = "Ciência da Computação"
	fmt.Println(e1)
	fmt.Println(e1.nome)

	e2 := estudante{pessoa{"Hipócrates", 99}, "Medicina"}
	fmt.Println(e2)
	e2.idade = 77
	fmt.Println(e2)

	p1 := pessoa{"Marcos Uchoa", 51}
	e3 := estudante{p1, "Jornalismo"}
	fmt.Println(e3)
	fmt.Println(e3.curso)

}
