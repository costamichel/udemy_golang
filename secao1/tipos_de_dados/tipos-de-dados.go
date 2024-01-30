package main

import (
	"errors"
	"fmt"
)

func main() {

	var numero_inteiro1 int8 = 123
	fmt.Println(numero_inteiro1)

	var numero_inteiro2 int16 = 12345
	fmt.Println(numero_inteiro2)

	var numero_inteiro3 int32 = 1234567890
	fmt.Println(numero_inteiro3)

	var numero_inteiro4 int64 = 1234567890123456789
	fmt.Println(numero_inteiro4)

	var numero_inteiro5 int = 1234567890123456789 // não funcionará em uma arquitetura de 32 bits
	fmt.Println(numero_inteiro5)

	var numero_unsigned1 uint8 = 123
	fmt.Println(numero_unsigned1)

	var numero_unsigned2 uint16 = 12345
	fmt.Println(numero_unsigned2)

	var numero_unsigned3 uint32 = 1234567890
	fmt.Println(numero_unsigned3)

	var numero_unsigned4 uint64 = 12345678901234567890
	fmt.Println(numero_unsigned4)

	var numero_unsigned5 uint = 12345678901234567890 // isso não funcionará numa arquitetura de 32 bits
	fmt.Println(numero_unsigned5)

	var numero_real1 float32 = 123.45454545454545454545454
	fmt.Println(numero_real1)

	var numero_real2 float64 = 12345678901234567890.333333333333333333333333333
	fmt.Println(numero_real2)

	var numero_rune rune = 1234567890
	fmt.Println(numero_rune)

	var numero_byte byte = 100
	fmt.Println(numero_byte)

	numero_infere1 := 123
	fmt.Println(numero_infere1)

	numero_infere2 := 123.45
	fmt.Println(numero_infere2)

	var v0_inteiro int
	fmt.Println(v0_inteiro)

	var v0_string string
	fmt.Println(v0_string)

	var v0_float64 float64
	fmt.Println(v0_float64)

	var texto string = "Texto"
	fmt.Println(texto)

	texto2 := 'B'
	fmt.Println(texto2)

	var booleano bool = true
	fmt.Println(booleano)
	var v0_booleano bool
	fmt.Println(v0_booleano)

	var err error
	fmt.Println(err)
	err = errors.New("Internal error!!!")
	fmt.Println(err)
	err = nil
	fmt.Println(err)
}
