package main

import "fmt"

type usuario struct {
	nome string
	idade int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main() {

	fmt.Println("arquivo structs")

	var u usuario
	u.nome = "joaozinho"
	u.idade = 12
	fmt.Println(u)

	endereco := endereco{"rua xazam", 129}
	usuario2 := usuario{"davi", 22, endereco}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 29, endereco: endereco}
	fmt.Println(usuario3)
}