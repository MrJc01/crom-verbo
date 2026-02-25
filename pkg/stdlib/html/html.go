// Package html implementa funções para geração de HTML tipada na linguagem Verbo.
// Parte da BibVerbo (Biblioteca Padrão).
//
// Uso em Verbo:
//
//	Incluir Html.
//	O pagina é CriarPagina de Html com ("Meu Site", "<h1>Olá!</h1>").
//	Exibir com (pagina).
package html

import (
	"fmt"
	"strings"
)

// CriarElemento cria uma tag HTML com conteúdo.
//
// Exemplo: CriarElemento("h1", "Título") → <h1>Título</h1>
func CriarElemento(tag, conteudo string) string {
	return fmt.Sprintf("<%s>%s</%s>", tag, conteudo, tag)
}

// CriarElementoComAtributos cria uma tag HTML com atributos e conteúdo.
//
// Exemplo: CriarElementoComAtributos("a", "href=\"/\"", "Link") → <a href="/">Link</a>
func CriarElementoComAtributos(tag, atributos, conteudo string) string {
	return fmt.Sprintf("<%s %s>%s</%s>", tag, atributos, conteudo, tag)
}

// CriarPagina gera uma página HTML completa.
//
// Exemplo: CriarPagina("Meu Site", "<h1>Olá</h1>") → <!DOCTYPE html>...
func CriarPagina(titulo, corpo string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s</title>
</head>
<body>
%s
</body>
</html>`, titulo, corpo)
}

// CriarPaginaComEstilo gera uma página HTML com CSS inline.
func CriarPaginaComEstilo(titulo, estilo, corpo string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s</title>
    <style>%s</style>
</head>
<body>
%s
</body>
</html>`, titulo, estilo, corpo)
}

// Atributo cria um par chave="valor" para uso em tags HTML.
//
// Exemplo: Atributo("class", "btn") → class="btn"
func Atributo(chave, valor string) string {
	return fmt.Sprintf(`%s="%s"`, chave, valor)
}

// ListaElementos concatena múltiplos elementos HTML com quebra de linha.
func ListaElementos(elementos ...string) string {
	return strings.Join(elementos, "\n")
}

// CriarLista gera uma lista HTML (ul) a partir de itens.
//
// Exemplo: CriarLista("Item 1", "Item 2") → <ul><li>Item 1</li><li>Item 2</li></ul>
func CriarLista(itens ...string) string {
	var lis []string
	for _, item := range itens {
		lis = append(lis, fmt.Sprintf("  <li>%s</li>", item))
	}
	return fmt.Sprintf("<ul>\n%s\n</ul>", strings.Join(lis, "\n"))
}

// toStringSlice converte []interface{} ou []string para []string.
func toStringSlice(v interface{}) []string {
	switch val := v.(type) {
	case []string:
		return val
	case []interface{}:
		result := make([]string, len(val))
		for i, item := range val {
			result[i] = fmt.Sprintf("%v", item)
		}
		return result
	default:
		return nil
	}
}

// CriarTabela gera uma tabela HTML a partir de headers e linhas.
// Aceita []string, []interface{}, [][]string ou [][]interface{} graças
// à conversão automática de tipos do runtime Verbo.
func CriarTabela(headersRaw interface{}, linhasRaw interface{}) string {
	headers := toStringSlice(headersRaw)

	// Converter linhas: pode ser [][]string, [][]interface{} ou []interface{} contendo sub-slices
	var linhas [][]string
	switch lv := linhasRaw.(type) {
	case [][]string:
		linhas = lv
	case []interface{}:
		for _, row := range lv {
			linhas = append(linhas, toStringSlice(row))
		}
	}

	var sb strings.Builder
	sb.WriteString("<table>\n  <thead>\n    <tr>\n")
	for _, h := range headers {
		sb.WriteString(fmt.Sprintf("      <th>%s</th>\n", h))
	}
	sb.WriteString("    </tr>\n  </thead>\n  <tbody>\n")
	for _, linha := range linhas {
		sb.WriteString("    <tr>\n")
		for _, cel := range linha {
			sb.WriteString(fmt.Sprintf("      <td>%s</td>\n", cel))
		}
		sb.WriteString("    </tr>\n")
	}
	sb.WriteString("  </tbody>\n</table>")
	return sb.String()
}

// CriarLink cria um link HTML.
//
// Exemplo: CriarLink("https://verbo.dev", "Verbo") → <a href="https://verbo.dev">Verbo</a>
func CriarLink(url, texto string) string {
	return fmt.Sprintf(`<a href="%s">%s</a>`, url, texto)
}

// CriarImagem cria uma tag de imagem HTML.
//
// Exemplo: CriarImagem("/img/logo.png", "Logo") → <img src="/img/logo.png" alt="Logo">
func CriarImagem(src, alt string) string {
	return fmt.Sprintf(`<img src="%s" alt="%s">`, src, alt)
}
