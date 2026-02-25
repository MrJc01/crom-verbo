package main

import "fmt"

func main() {
	func() {
		defer func() {
			if erro := recover(); erro != nil {
				fmt.Println("Operação falhou. Motivo:")
			}
		}()
		fmt.Println("Tentando operação arriscada...")
		panic("Saldo insuficiente!")
	}()
	fmt.Println(erro)
	fmt.Println("Programa continua normalmente.")
}
