package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays")

	var array1 [3]int
	array1[0] = 99
	fmt.Println(array1)

	array2 := [4]string{"um", "dois", "tres", "quatro"}
	fmt.Println(array2)

	array3 := [...]int{1,2,3,4}
	fmt.Println(array3)

	slice := []int{10, 20, 30, 40} 
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array3))
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 50)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)
} 