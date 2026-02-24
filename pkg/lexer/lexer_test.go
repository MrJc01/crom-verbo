package lexer

import (
	"testing"
)

func TestTokenizarArtigosEAtribuicao(t *testing.T) {
	testes := []struct {
		nome     string
		entrada  string
		esperado []Token
	}{
		{
			nome:    "constante com artigo definido",
			entrada: `A mensagem é "Olá".`,
			esperado: []Token{
				{Tipo: TOKEN_ARTIGO_DEFINIDO, Valor: "A"},
				{Tipo: TOKEN_IDENTIFICADOR, Valor: "mensagem"},
				{Tipo: TOKEN_E_ACENTO, Valor: "é"},
				{Tipo: TOKEN_TEXTO, Valor: "Olá"},
				{Tipo: TOKEN_PONTO, Valor: "."},
				{Tipo: TOKEN_FIM, Valor: ""},
			},
		},
		{
			nome:    "variável com artigo indefinido",
			entrada: `Um contador está 0.`,
			esperado: []Token{
				{Tipo: TOKEN_ARTIGO_INDEFINIDO, Valor: "Um"},
				{Tipo: TOKEN_IDENTIFICADOR, Valor: "contador"},
				{Tipo: TOKEN_ESTA, Valor: "está"},
				{Tipo: TOKEN_NUMERO, Valor: "0"},
				{Tipo: TOKEN_PONTO, Valor: "."},
				{Tipo: TOKEN_FIM, Valor: ""},
			},
		},
		{
			nome:    "numero decimal",
			entrada: `Um pi é 3.14.`,
			esperado: []Token{
				{Tipo: TOKEN_ARTIGO_INDEFINIDO, Valor: "Um"},
				{Tipo: TOKEN_IDENTIFICADOR, Valor: "pi"},
				{Tipo: TOKEN_E_ACENTO, Valor: "é"},
				{Tipo: TOKEN_NUMERO, Valor: "3.14"},
				{Tipo: TOKEN_PONTO, Valor: "."},
				{Tipo: TOKEN_FIM, Valor: ""},
			},
		},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			lex := Novo(tt.entrada)
			tokens, err := lex.Tokenizar()
			if err != nil {
				t.Fatalf("erro inesperado: %v", err)
			}

			if len(tokens) != len(tt.esperado) {
				t.Fatalf("esperava %d tokens, obteve %d:\n%v", len(tt.esperado), len(tokens), tokens)
			}

			for i, esperado := range tt.esperado {
				if tokens[i].Tipo != esperado.Tipo {
					t.Errorf("token[%d]: esperava tipo %s, obteve %s (valor: %q)",
						i, esperado.Tipo.NomeLegivel(), tokens[i].Tipo.NomeLegivel(), tokens[i].Valor)
				}
				if tokens[i].Valor != esperado.Valor {
					t.Errorf("token[%d]: esperava valor %q, obteve %q",
						i, esperado.Valor, tokens[i].Valor)
				}
			}
		})
	}
}

func TestTokenizarPalavrasChave(t *testing.T) {
	testes := []struct {
		nome     string
		entrada  string
		esperado []Token
	}{
		{
			nome:    "condicional se/então",
			entrada: `Se a idade for menor que 18, então:`,
			esperado: []Token{
				{Tipo: TOKEN_SE, Valor: "Se"},
				{Tipo: TOKEN_ARTIGO_DEFINIDO, Valor: "a"},
				{Tipo: TOKEN_IDENTIFICADOR, Valor: "idade"},
				{Tipo: TOKEN_FOR, Valor: "for"},
				{Tipo: TOKEN_MENOR, Valor: "menor"},
				{Tipo: TOKEN_QUE, Valor: "que"},
				{Tipo: TOKEN_NUMERO, Valor: "18"},
				{Tipo: TOKEN_VIRGULA, Valor: ","},
				{Tipo: TOKEN_ENTAO, Valor: "então"},
				{Tipo: TOKEN_DOIS_PONTOS, Valor: ":"},
				{Tipo: TOKEN_FIM, Valor: ""},
			},
		},
		{
			nome:    "loop repita",
			entrada: `Repita 5 vezes:`,
			esperado: []Token{
				{Tipo: TOKEN_REPITA, Valor: "Repita"},
				{Tipo: TOKEN_NUMERO, Valor: "5"},
				{Tipo: TOKEN_VEZES, Valor: "vezes"},
				{Tipo: TOKEN_DOIS_PONTOS, Valor: ":"},
				{Tipo: TOKEN_FIM, Valor: ""},
			},
		},
		{
			nome:    "declaração de função",
			entrada: `Para Calcular usando (valor: Inteiro):`,
			esperado: []Token{
				{Tipo: TOKEN_PARA, Valor: "Para"},
				{Tipo: TOKEN_IDENTIFICADOR, Valor: "Calcular"},
				{Tipo: TOKEN_USANDO, Valor: "usando"},
				{Tipo: TOKEN_PARENTESE_ABRE, Valor: "("},
				{Tipo: TOKEN_IDENTIFICADOR, Valor: "valor"},
				{Tipo: TOKEN_DOIS_PONTOS, Valor: ":"},
				{Tipo: TOKEN_TIPO, Valor: "Inteiro"},
				{Tipo: TOKEN_PARENTESE_FECHA, Valor: ")"},
				{Tipo: TOKEN_DOIS_PONTOS, Valor: ":"},
				{Tipo: TOKEN_FIM, Valor: ""},
			},
		},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			lex := Novo(tt.entrada)
			tokens, err := lex.Tokenizar()
			if err != nil {
				t.Fatalf("erro inesperado: %v", err)
			}

			if len(tokens) != len(tt.esperado) {
				t.Fatalf("esperava %d tokens, obteve %d:\n%v", len(tt.esperado), len(tokens), tokens)
			}

			for i, esperado := range tt.esperado {
				if tokens[i].Tipo != esperado.Tipo {
					t.Errorf("token[%d]: esperava tipo %s, obteve %s (valor: %q)",
						i, esperado.Tipo.NomeLegivel(), tokens[i].Tipo.NomeLegivel(), tokens[i].Valor)
				}
				if tokens[i].Valor != esperado.Valor {
					t.Errorf("token[%d]: esperava valor %q, obteve %q",
						i, esperado.Valor, tokens[i].Valor)
				}
			}
		})
	}
}

func TestTokenizarOperadores(t *testing.T) {
	entrada := `contador + 1 - 2 * 3 / 4`
	lex := Novo(entrada)
	tokens, err := lex.Tokenizar()
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	tiposEsperados := []TokenType{
		TOKEN_IDENTIFICADOR, TOKEN_MAIS, TOKEN_NUMERO,
		TOKEN_MENOS, TOKEN_NUMERO, TOKEN_MULTIPLICAR,
		TOKEN_NUMERO, TOKEN_DIVIDIR, TOKEN_NUMERO, TOKEN_FIM,
	}

	if len(tokens) != len(tiposEsperados) {
		t.Fatalf("esperava %d tokens, obteve %d", len(tiposEsperados), len(tokens))
	}

	for i, tipo := range tiposEsperados {
		if tokens[i].Tipo != tipo {
			t.Errorf("token[%d]: esperava %s, obteve %s",
				i, tipo.NomeLegivel(), tokens[i].Tipo.NomeLegivel())
		}
	}
}

func TestTokenizarUTF8Acentuacao(t *testing.T) {
	entrada := `A ação é "válida". Um índice está 0.`
	lex := Novo(entrada)
	tokens, err := lex.Tokenizar()
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	// Verifica que "ação", "é", "válida", "índice", "está" são tokenizados corretamente
	palavrasEsperadas := map[string]bool{
		"ação":    true,
		"é":       true,
		"válida":  true,
		"índice":  true,
		"está":    true,
	}

	for _, tok := range tokens {
		if tok.Valor != "" {
			delete(palavrasEsperadas, tok.Valor)
		}
	}

	if len(palavrasEsperadas) > 0 {
		t.Errorf("palavras acentuadas não encontradas nos tokens: %v", palavrasEsperadas)
	}
}

func TestTokenizarComentario(t *testing.T) {
	entrada := `// Este é um comentário
A mensagem é "teste".`
	lex := Novo(entrada)
	tokens, err := lex.Tokenizar()
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	// O comentário deve ser ignorado
	if tokens[0].Tipo != TOKEN_ARTIGO_DEFINIDO {
		t.Errorf("esperava primeiro token ser ARTIGO_DEFINIDO, obteve %s", tokens[0].Tipo.NomeLegivel())
	}
}

func TestTokenizarTextoNaoFechado(t *testing.T) {
	entrada := `A mensagem é "incompleto`
	lex := Novo(entrada)
	_, err := lex.Tokenizar()
	if err == nil {
		t.Fatal("esperava erro para texto não fechado, mas não obteve nenhum")
	}
}

func TestTokenizarExibirComFuncao(t *testing.T) {
	entrada := `Exibir com ("Olá, Mundo!").`
	lex := Novo(entrada)
	tokens, err := lex.Tokenizar()
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	esperado := []TokenType{
		TOKEN_EXIBIR, TOKEN_COM, TOKEN_PARENTESE_ABRE,
		TOKEN_TEXTO, TOKEN_PARENTESE_FECHA, TOKEN_PONTO, TOKEN_FIM,
	}

	if len(tokens) != len(esperado) {
		t.Fatalf("esperava %d tokens, obteve %d:\n%v", len(esperado), len(tokens), tokens)
	}

	for i, tipo := range esperado {
		if tokens[i].Tipo != tipo {
			t.Errorf("token[%d]: esperava %s, obteve %s (valor: %q)",
				i, tipo.NomeLegivel(), tokens[i].Tipo.NomeLegivel(), tokens[i].Valor)
		}
	}
}

func TestTokenizarLinhaColuna(t *testing.T) {
	entrada := "A x é 1.\nUm y está 2."
	lex := Novo(entrada)
	tokens, err := lex.Tokenizar()
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	// "Um" deve estar na linha 2
	for _, tok := range tokens {
		if tok.Valor == "Um" {
			if tok.Linha != 2 {
				t.Errorf("esperava 'Um' na linha 2, obteve linha %d", tok.Linha)
			}
			break
		}
	}
}
