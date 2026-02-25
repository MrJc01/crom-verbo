package texto

import "strings"

// Tamanho retorna o tamanho de uma string
func Tamanho(s string) int { return len(s) }

// Maiusculas retorna a string em letras maiúsculas
func Maiusculas(s string) string { return strings.ToUpper(s) }

// Minusculas retorna a string em letras minúsculas
func Minusculas(s string) string { return strings.ToLower(s) }

// Contem verifica se a string contém a substring fornecida
func Contem(s, substr string) bool { return strings.Contains(s, substr) }

// Substituir substitui todas as ocorrências de 'velho' por 'novo'
func Substituir(s, velho, novo string) string { return strings.ReplaceAll(s, velho, novo) }

// Dividir particiona a string usando um separador
func Dividir(s, separador string) []interface{} {
	parts := strings.Split(s, separador)
	res := make([]interface{}, len(parts))
	for i, p := range parts {
		res[i] = p
	}
	return res
}
