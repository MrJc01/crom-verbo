package html

import (
	"strings"
	"testing"
)

func TestCriarElemento(t *testing.T) {
	testes := []struct {
		tag, conteudo, esperado string
	}{
		{"h1", "Título", "<h1>Título</h1>"},
		{"p", "Texto aqui", "<p>Texto aqui</p>"},
		{"div", "", "<div></div>"},
		{"span", "Olá Mundo", "<span>Olá Mundo</span>"},
	}

	for _, tt := range testes {
		resultado := CriarElemento(tt.tag, tt.conteudo)
		if resultado != tt.esperado {
			t.Errorf("CriarElemento(%q, %q): esperava %q, obteve %q", tt.tag, tt.conteudo, tt.esperado, resultado)
		}
	}
}

func TestCriarElementoComAtributos(t *testing.T) {
	resultado := CriarElementoComAtributos("a", `href="/"`, "Link")
	esperado := `<a href="/">Link</a>`
	if resultado != esperado {
		t.Errorf("esperava %q, obteve %q", esperado, resultado)
	}
}

func TestCriarPagina(t *testing.T) {
	resultado := CriarPagina("Meu Site", "<h1>Olá</h1>")

	verificacoes := []string{
		"<!DOCTYPE html>",
		"<title>Meu Site</title>",
		"<h1>Olá</h1>",
		`lang="pt-BR"`,
		`charset="UTF-8"`,
	}

	for _, v := range verificacoes {
		if !strings.Contains(resultado, v) {
			t.Errorf("CriarPagina: resultado não contém %q.\nSaída:\n%s", v, resultado)
		}
	}
}

func TestCriarPaginaComEstilo(t *testing.T) {
	resultado := CriarPaginaComEstilo("Site", "body{color:red}", "<p>OK</p>")

	if !strings.Contains(resultado, "<style>body{color:red}</style>") {
		t.Errorf("esperava bloco de estilo.\nSaída:\n%s", resultado)
	}
}

func TestAtributo(t *testing.T) {
	resultado := Atributo("class", "btn verde")
	esperado := `class="btn verde"`
	if resultado != esperado {
		t.Errorf("esperava %q, obteve %q", esperado, resultado)
	}
}

func TestListaElementos(t *testing.T) {
	resultado := ListaElementos("<h1>A</h1>", "<p>B</p>")
	if !strings.Contains(resultado, "<h1>A</h1>\n<p>B</p>") {
		t.Errorf("esperava elementos concatenados:\n%s", resultado)
	}
}

func TestCriarLista(t *testing.T) {
	resultado := CriarLista("Item 1", "Item 2", "Item 3")

	verificacoes := []string{"<ul>", "</ul>", "<li>Item 1</li>", "<li>Item 2</li>", "<li>Item 3</li>"}
	for _, v := range verificacoes {
		if !strings.Contains(resultado, v) {
			t.Errorf("CriarLista: resultado não contém %q.\nSaída:\n%s", v, resultado)
		}
	}
}

func TestCriarLink(t *testing.T) {
	resultado := CriarLink("https://verbo.dev", "Verbo")
	esperado := `<a href="https://verbo.dev">Verbo</a>`
	if resultado != esperado {
		t.Errorf("esperava %q, obteve %q", esperado, resultado)
	}
}

func TestCriarImagem(t *testing.T) {
	resultado := CriarImagem("/img/logo.png", "Logo")
	esperado := `<img src="/img/logo.png" alt="Logo">`
	if resultado != esperado {
		t.Errorf("esperava %q, obteve %q", esperado, resultado)
	}
}

func TestCriarTabela(t *testing.T) {
	headers := []string{"Nome", "Idade"}
	linhas := [][]string{
		{"Ada", "36"},
		{"Juan", "25"},
	}
	resultado := CriarTabela(headers, linhas)

	verificacoes := []string{"<table>", "</table>", "<th>Nome</th>", "<td>Ada</td>", "<td>25</td>"}
	for _, v := range verificacoes {
		if !strings.Contains(resultado, v) {
			t.Errorf("CriarTabela: resultado não contém %q.\nSaída:\n%s", v, resultado)
		}
	}
}
