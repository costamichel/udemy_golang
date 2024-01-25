package main

import "fmt"

func main() {
	var var1 string = "Texto 1"
	fmt.Println(var1)

	var2 := "Texto 2"
	fmt.Println(var2)

	var (
		var3 string = "Texto 3"
		var4 int    = 4
	)
	fmt.Println(var3)
	fmt.Println(var4)

	var5, var6 := "Texto 5", "Texto 6"
	fmt.Println(var5)
	fmt.Println(var6)

	const c1 string = "c1 é constante"
	fmt.Println(c1)

	var5, var6 = var6, var5
	fmt.Println(var5)
	fmt.Println(var6)

}
