package main

import "fmt"

func somar2(n1, n2 int8) int8 {
	return somar(n1, n2)
}

func calculosMatematicos(n1, n2 int8) (int8, int8) { // Retorno múltiplo
	return n1 + n2, n1 - n2
}

func main() {

	resultadoSoma := somar(3, 2)
	fmt.Println(resultadoSoma)

	resultadoSoma = somar2(5, 7)
	fmt.Println(resultadoSoma)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(5, 4)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	resultadoSoma, _ = calculosMatematicos(4, 8)
	fmt.Println(resultadoSoma)

	_, resultadoSubtracao = calculosMatematicos(4, 6)
	fmt.Println(resultadoSubtracao)
}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}
