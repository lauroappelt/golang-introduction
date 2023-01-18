package main

import "fmt"

//func nome (param) retorno
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//multiplos resultados
func calculosMatematicos(n1 int8, n2 int8) (int8, int8){
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(1, 2)
	fmt.Println(soma)

	var f = func (txt string) {
		fmt.Println(txt)
	}

	f("texto")

	soma, sub := calculosMatematicos(10, 5)
	fmt.Println(soma, sub)
}