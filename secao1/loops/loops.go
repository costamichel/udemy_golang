package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0
	for i < 10 {
		fmt.Println("i = ", i)
		i++
		//time.Sleep(time.Second) // Faz "dormir" por um segundo
	}

	fmt.Println("------------------")
	for j := 0; j < 5; j++ {
		fmt.Println("j = ", j)
	}

	fmt.Println("-------------------")
	nomes := []string{"Davi", "Isabelle", "Matteus"}
	for idx, nome := range nomes {
		fmt.Println(idx+1, " - ", nome)
	}

	for idx, letra := range "BIG BROTHER" {
		fmt.Println(idx, letra, string(letra))
	}

	brothers := map[string]int{
		"Davi":     23,
		"Cunhã":    25,
		"Alegrete": 27,
	}
	for chave, valor := range brothers {
		fmt.Println(chave, "tem", valor, "anos!")
	}

	i = 0
	for {
		fmt.Println("Loop infinito. Dormindo\tzzz")
		time.Sleep((time.Second))
		i++
		if i > 3 {
			break
		}
	}
}
