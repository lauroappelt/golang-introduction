package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int8 = 28
	var variavel2 int8 = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	//ponteiro = endereco de memoria
	var variavel3 int8 = 100
	var ponteiro *int8 = &variavel3

	fmt.Println(variavel3, ponteiro)
	fmt.Println(*ponteiro)
}	