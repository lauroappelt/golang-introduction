package main

import "fmt"

func main() {
	//explicito
	var variavel1 string = "teste"
	fmt.Println(variavel1)

	//implicito
	variavel2 := "teste 2"
	fmt.Println(variavel2)


	//declarar n variaveis junto
	var  (
		variavel3 string = "teste 3"
		variavel4 string = "teste 4"
	)

	variavel5, variavel6 := "teste 5", "teste 6"

	fmt.Println(variavel3)
	fmt.Println(variavel4)
	fmt.Println(variavel5)
	fmt.Println(variavel6)

	//constante
	const constant1 = "const 1"
}