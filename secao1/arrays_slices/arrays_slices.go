package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e slices")

	// Arrays
	var array1 [4]int32
	fmt.Println(array1)
	fmt.Println(array1[1])
	array1[1] = 5
	fmt.Println(array1[1])
	fmt.Println(array1)
	fmt.Println("-----------------")

	array2 := [5]int64{5, 10, 15, 20, 25}
	fmt.Println(array2)
	fmt.Println("-----------------")

	array3 := [...]int32{10, 20, 30}
	fmt.Println(array3)
	fmt.Println("-----------------")

	//Slices
	slice1 := []int{3, 6, 9}
	fmt.Println(slice1)
	slice1 = append(slice1, 12)
	fmt.Println(slice1)
	fmt.Println("-----------------")

	slice2 := array2[1:3]
	fmt.Println(slice2)
	array2[1] = 11
	fmt.Println(slice2)
	fmt.Println(reflect.TypeOf(slice2))
	fmt.Println(reflect.TypeOf(array2))
	fmt.Println("-----------------")

	slice3 := make([]string, 4, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	slice3 = append(slice3, "Segunda")
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	slice3 = append(slice3, "Terça")
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	slice3 = append(slice3, "Quarta")
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	fmt.Println("-----------------")

	slice4 := []float32{2, 4, 6}
	fmt.Println(slice4)
	fmt.Println(len(slice4), cap(slice4))
	slice4 = append(slice4, 8.1)
	fmt.Println(slice4)
	fmt.Println(len(slice4), cap(slice4))
}
