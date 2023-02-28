package main

import "fmt"

func main() {
	fmt.Println("map")

	usuario := map[string] string {
		"nome": "carlos",
		"sobrenome" : "almeida",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro": "joao",
			"segundo": "pedro",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)
}