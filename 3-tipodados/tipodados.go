package main

import (
	"errors"
	"fmt"
)

func main() {
	//int8, 16, 32, 64 bits
	var numero int16 = 100
	fmt.Println(numero)

	//int pega automatico da arq dá maquina
	var numero2 int = 99999
	fmt.Println(numero2)

	// alias rune = int32
	var numero3 rune = 12345
	fmt.Println(numero3)

	// alias byte = int8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1  float32 = 121212.34
	fmt.Println(numeroReal1)

	//flot64
	var numeroReal2 float64 = 1212121.12
	fmt.Println(numeroReal2)

	//float implicito pega bit da arq
	numeroReal3 := 12121.21
	fmt.Println(numeroReal3)

	//strings
	var str string = "string"
	fmt.Println(str)

	str2 := "string3"
	fmt.Println(str2)

	var stringVazia  string
	fmt.Println(stringVazia)

	var intVazio int
	fmt.Println(intVazio)

	var booleano bool = true
	fmt.Println(booleano)

	//erro é um tipo
	var erro error = errors.New("erro interno")
	fmt.Println(erro)
}