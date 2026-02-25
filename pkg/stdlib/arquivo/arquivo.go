package arquivo

import "os"

// LerTexto obtém o conteúdo de um arquivo de texto como string
func LerTexto(caminho string) string {
	b, err := os.ReadFile(caminho)
	if err != nil {
		panic("Erro ao ler arquivo " + caminho + ": " + err.Error())
	}
	return string(b)
}

// EscreverTexto grava o conteúdo de texto no arquivo
func EscreverTexto(caminho, conteudo string) {
	err := os.WriteFile(caminho, []byte(conteudo), 0644)
	if err != nil {
		panic("Erro ao escrever arquivo " + caminho + ": " + err.Error())
	}
}
