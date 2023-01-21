package main

import "fmt"

type pessoa struct {
	nome string
	idade int8
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Println("heranca")

	p1 := pessoa{"joaozinho", 45}
	fmt.Println(p1)

	e1 := estudante{p1, "engenharia", "usp"}
	fmt.Println(e1)

	fmt.Println(e1.nome)
}