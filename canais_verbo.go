package main

import (
	"fmt"
	"sync"
)

func ProcessarPacote(pacote int) interface{} {
	fmt.Println("Processando pacote...")
	resultado := (pacote * 10)
	resultados <- resultado
	return nil
}

func principal() interface{} {
	fmt.Println("Iniciando workers simultâneos (Goroutines)...")
	{
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			ProcessarPacote(1)
		}()
		wg.Wait()
	}
	{
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			ProcessarPacote(2)
		}()
		wg.Wait()
	}
	{
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			ProcessarPacote(3)
		}()
		wg.Wait()
	}
	res1 := <-resultados
	res2 := <-resultados
	res3 := <-resultados
	fmt.Println("Todos os pacotes processados!")
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	return nil
}

func main() {
	resultados := make(chan int)
}
