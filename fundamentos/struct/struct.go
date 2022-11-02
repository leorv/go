package main

import "fmt"

func main() {
	pessoa1 := Pessoa{"Tiago", "Temporin", 31}
	pessoa2 := Pessoa{Nome: "João", Sobrenome: "Silva", Idade: 35}
	pessoa3 := Pessoa{Nome: "Aline", Idade: 35}

	fmt.Println(pessoa1, pessoa2, pessoa3)
}

// Estrutura guarda apenas estado, e não comportamento.

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int8
}
