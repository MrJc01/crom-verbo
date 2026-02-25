package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Tarefa 1: Processando dados...")
	}()
	wg.Wait()
	fmt.Println("Tarefa 2: Enviando relatório...")
	fmt.Println("Tarefa 3: Atualizando cache...")
	fmt.Println("Todas as tarefas completadas!")
}
