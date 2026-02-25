package main

import "fmt"

func main() {
	precos := []interface{}{10, 20, 30, 40, 50}
	primeiro_preco := precos[0]
	fmt.Println("Primeiro preço:")
	fmt.Println(primeiro_preco)
	fmt.Println("Todos os preços:")
	for _, preco := range precos {
		fmt.Println(preco)
	}
}
