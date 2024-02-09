package main

import "fmt"

func main() {
	var var1 int64
	var pointer1 *int64

	var1 = 150
	pointer1 = &var1
	fmt.Println(var1, pointer1, *pointer1)

	var1 = 200
	fmt.Println(var1, pointer1, *pointer1)

	var var2 int64 = 300
	pointer1 = &var2
	fmt.Println(var1, var2, pointer1, *pointer1)
}
