package main

import "fmt"

type Pessoa struct {
	Nome string
	Idade int
	Saldo float64
}

func main() {
	cliente := Pessoa{Nome: "Ada Lovelace", Idade: 36, Saldo: 1500.50}
	fmt.Println("Cliente cadastrado:")
	fmt.Println(cliente)
}
