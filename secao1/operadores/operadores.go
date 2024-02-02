package main

import "fmt"

func main() {

	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 2 * 6
	divisao := 12 / 3
	resto := 5 / 2
	fmt.Println(soma, subtracao, multiplicacao, divisao, resto)
	fmt.Println("-----------------")

	bt := true
	bf := false
	fmt.Println(bt && bf)
	fmt.Println(bt || bf)
	fmt.Println(!bf)
	fmt.Println("----------------")

	fmt.Println(1 > 2)
	fmt.Println(2 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(2 <= 1)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println("----------------")

	num := 5
	fmt.Println(num)
	num++
	fmt.Println(num)
	num += 10
	fmt.Println(num)
	num -= 10
	fmt.Println(num)
	num--
	fmt.Println(num)
	num %= 3
	fmt.Println(num)

}
