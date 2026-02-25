package main

import (
	"fmt"
	"github.com/juanxto/crom-verbo/pkg/stdlib/matematica"
	"github.com/juanxto/crom-verbo/pkg/stdlib/texto"
)

type Pessoa struct {
	Nome string
	Idade int
	Saldo float64
}

func Saudacao(pessoa Pessoa) interface{} {
	nome_maiusculo := texto.Maiusculas(pessoa.Nome)
	fmt.Println((("Olá " + nome_maiusculo) + "!"))
	return nil
}

func principal() interface{} {
	cliente := Pessoa{Nome: "Ada Lovelace", Idade: 36, Saldo: 1500.50}
	fmt.Println("Cliente cadastrado:")
	fmt.Println(cliente)
	Saudacao(cliente)
	fmt.Println("--- Testando Tratamento de Erros ---")
	func() {
		defer func() {
			if erro := recover(); erro != nil {
				fmt.Println("Operação falhou. Motivo:")
				fmt.Println(erro)
			}
		}()
		fmt.Println("Tentando executar uma operação arriscada...")
		panic("Saldo insuficiente para a operação.")
	}()
	fmt.Println("--- Testando Biblioteca Padrão ---")
	valor_absoluto := matematica.Absoluto(-42.5)
	fmt.Println("O valor absoluto de -42.5 é:")
	fmt.Println(valor_absoluto)
	fmt.Println("--- Testando Listas ---")
	precos := []interface{}{10, 20, 30, 40, 50}
	primeiro_preco := precos[0]
	fmt.Println("O primeiro preço da lista é:")
	fmt.Println(primeiro_preco)
	fmt.Println("--- Iterando sobre a lista ---")
	for _, preco := range precos {
		fmt.Println(preco)
	}
	principal()
	return nil
}

func main() {
}
