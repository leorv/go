package main

import "fmt"

// não precisa declarar o tipo, pode ser convenção não declarar para ter um código mais limpo.
const ebulicaoF float64 = 212.0

func main() {
	var tempF float64 = ebulicaoF
	var tempC float64 = (tempF - 32) * 5 / 9

	fmt.Println("A temperatura de ebulição da água em F:", tempF)
	fmt.Println("A temperatura de ebulição da água em C:", tempC)

	// Usando Printf
	fmt.Println("===== Agora usando Printf =====")
	fmt.Printf("A temperatura de ebulição da água em F é %g e em C é %g.\n", tempF, tempC)

	// Descobrindo o tipo da variável
	fmt.Printf("%T\n", tempC)
}
