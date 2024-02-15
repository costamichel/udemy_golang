package main

import "fmt"

func main() {
	fmt.Println("Maps!!!")

	usuario := map[string]string{
		"nome":      "Michel",
		"sobrenome": "Costa",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"]) // Não é possível acessar com `usuario.nome`

	alunosNotas := map[string]map[string]float32{
		"Soteldo": {
			"História":  8.5,
			"Português": 9.2,
		},
		"Galvao": {
			"História":  4.5,
			"Português": 6.3,
		},
	}
	fmt.Println(alunosNotas)
	fmt.Println(alunosNotas["Soteldo"]["Português"])

	delete(alunosNotas["Galvao"], "Português")
	fmt.Println(alunosNotas)

	alunosNotas["Galvao"]["Matemática"] = 2.1
	fmt.Println(alunosNotas)
}
