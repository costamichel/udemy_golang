package main

import "fmt"

func esportes(num int) string {
	switch num {
	case 1:
		return "Futebol"
	case 2:
		return "Vôlei"
	case 3:
		return "Natação"
	default:
		return "Outro esporte"
	}
}

func esportes2(num int) string {
	var esporte_nome string

	switch {
	case num == 1:
		esporte_nome = "Futebol"
		fallthrough // faz cair direto no próximo case SEM VERIFICAR (neste caso retornará o valor "Vôlei")
	case num == 2:
		esporte_nome = "Vôlei"
	case num == 3:
		esporte_nome = "Natação"
	default:
		esporte_nome = "Outro Esporte"
	}

	return esporte_nome
}

func main() {
	fmt.Println("Switch\n")

	esp1 := esportes(1)
	fmt.Println(esp1)
	fmt.Println(esportes(10))
	fmt.Println("------------------------")

	esp2 := esportes2(3)
	fmt.Println(esp2)
	fmt.Println(esportes2(1))

}
