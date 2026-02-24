// Package lexer implementa o scanner léxico da linguagem Verbo.
// O Lexer lê o código-fonte caractere a caractere (runa Go para suporte UTF-8)
// e produz uma sequência de tokens tipados.
package lexer

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Lexer é o analisador léxico da linguagem Verbo.
type Lexer struct {
	entrada     string  // código-fonte completo
	posicao     int     // posição atual no byte
	posLeitura  int     // posição de leitura (próximo byte)
	caractere   rune    // caractere (runa) atual sendo examinado
	linha       int     // linha atual (1-indexed)
	coluna      int     // coluna atual (1-indexed)
	tokens      []Token // tokens já produzidos
}

// Novo cria um novo Lexer para o código-fonte fornecido.
func Novo(entrada string) *Lexer {
	l := &Lexer{
		entrada: entrada,
		linha:   1,
		coluna:  0,
	}
	l.lerCaractere()
	return l
}

// lerCaractere avança para o próximo caractere Unicode no código-fonte.
func (l *Lexer) lerCaractere() {
	if l.posLeitura >= len(l.entrada) {
		l.caractere = 0 // EOF
		l.posicao = l.posLeitura
		return
	}

	r, tamanho := utf8.DecodeRuneInString(l.entrada[l.posLeitura:])
	l.caractere = r
	l.posicao = l.posLeitura
	l.posLeitura += tamanho
	l.coluna++
}

// espiarProximo retorna o próximo caractere sem avançar a posição.
func (l *Lexer) espiarProximo() rune {
	if l.posLeitura >= len(l.entrada) {
		return 0
	}
	r, _ := utf8.DecodeRuneInString(l.entrada[l.posLeitura:])
	return r
}

// pularEspacos ignora espaços em branco e controla contagem de linhas.
func (l *Lexer) pularEspacos() {
	for l.caractere == ' ' || l.caractere == '\t' || l.caractere == '\r' || l.caractere == '\n' {
		if l.caractere == '\n' {
			l.linha++
			l.coluna = 0
		}
		l.lerCaractere()
	}
}

// pularComentario ignora comentários de linha (começam com //).
func (l *Lexer) pularComentario() {
	if l.caractere == '/' && l.espiarProximo() == '/' {
		for l.caractere != '\n' && l.caractere != 0 {
			l.lerCaractere()
		}
	}
}

// Tokenizar processa todo o código-fonte e retorna a lista de tokens.
func (l *Lexer) Tokenizar() ([]Token, error) {
	var tokens []Token
	var erros []string

	for {
		l.pularEspacos()
		l.pularComentario()
		l.pularEspacos()

		if l.caractere == 0 {
			tokens = append(tokens, Token{
				Tipo:   TOKEN_FIM,
				Valor:  "",
				Linha:  l.linha,
				Coluna: l.coluna,
			})
			break
		}

		tok, err := l.proximoToken()
		if err != nil {
			erros = append(erros, err.Error())
			l.lerCaractere() // avança para tentar recuperar
			continue
		}
		// Pular pseudo-tokens de comentário
		if tok.Tipo == TOKEN_ILEGAL && tok.Valor == "//comentario" {
			continue
		}
		tokens = append(tokens, tok)
	}

	if len(erros) > 0 {
		return tokens, fmt.Errorf("erros léxicos encontrados:\n%s", strings.Join(erros, "\n"))
	}

	return tokens, nil
}

// proximoToken identifica e retorna o próximo token.
func (l *Lexer) proximoToken() (Token, error) {
	linha := l.linha
	coluna := l.coluna

	switch {
	// String literal
	case l.caractere == '"':
		valor, err := l.lerTexto()
		if err != nil {
			return Token{}, err
		}
		return Token{Tipo: TOKEN_TEXTO, Valor: valor, Linha: linha, Coluna: coluna}, nil

	// Número
	case unicode.IsDigit(l.caractere):
		valor := l.lerNumero()
		return Token{Tipo: TOKEN_NUMERO, Valor: valor, Linha: linha, Coluna: coluna}, nil

	// Palavra (identificador ou palavra-chave)
	case ehLetra(l.caractere):
		valor := l.lerPalavra()
		tipo := BuscarPalavraChave(valor)
		return Token{Tipo: tipo, Valor: valor, Linha: linha, Coluna: coluna}, nil

	// Operadores e delimitadores
	case l.caractere == '+':
		l.lerCaractere()
		return Token{Tipo: TOKEN_MAIS, Valor: "+", Linha: linha, Coluna: coluna}, nil

	case l.caractere == '-':
		l.lerCaractere()
		return Token{Tipo: TOKEN_MENOS, Valor: "-", Linha: linha, Coluna: coluna}, nil

	case l.caractere == '*':
		l.lerCaractere()
		return Token{Tipo: TOKEN_MULTIPLICAR, Valor: "*", Linha: linha, Coluna: coluna}, nil

	case l.caractere == '/':
		if l.espiarProximo() == '/' {
			// Comentário de linha — pular até o fim da linha
			for l.caractere != '\n' && l.caractere != 0 {
				l.lerCaractere()
			}
			// Retornar o próximo token real (recursivamente via Tokenizar loop)
			return Token{Tipo: TOKEN_ILEGAL, Valor: "//comentario"}, nil
		}
		l.lerCaractere()
		return Token{Tipo: TOKEN_DIVIDIR, Valor: "/", Linha: linha, Coluna: coluna}, nil

	case l.caractere == '=':
		l.lerCaractere()
		return Token{Tipo: TOKEN_ATRIBUIR, Valor: "=", Linha: linha, Coluna: coluna}, nil

	case l.caractere == '.':
		l.lerCaractere()
		return Token{Tipo: TOKEN_PONTO, Valor: ".", Linha: linha, Coluna: coluna}, nil

	case l.caractere == ':':
		l.lerCaractere()
		return Token{Tipo: TOKEN_DOIS_PONTOS, Valor: ":", Linha: linha, Coluna: coluna}, nil

	case l.caractere == ',':
		l.lerCaractere()
		return Token{Tipo: TOKEN_VIRGULA, Valor: ",", Linha: linha, Coluna: coluna}, nil

	case l.caractere == '(':
		l.lerCaractere()
		return Token{Tipo: TOKEN_PARENTESE_ABRE, Valor: "(", Linha: linha, Coluna: coluna}, nil

	case l.caractere == ')':
		l.lerCaractere()
		return Token{Tipo: TOKEN_PARENTESE_FECHA, Valor: ")", Linha: linha, Coluna: coluna}, nil

	case l.caractere == '[':
		l.lerCaractere()
		return Token{Tipo: TOKEN_COLCHETE_ABRE, Valor: "[", Linha: linha, Coluna: coluna}, nil

	case l.caractere == ']':
		l.lerCaractere()
		return Token{Tipo: TOKEN_COLCHETE_FECHA, Valor: "]", Linha: linha, Coluna: coluna}, nil

	default:
		ch := l.caractere
		l.lerCaractere()
		return Token{}, fmt.Errorf("linha %d, coluna %d: caractere inesperado '%c' (U+%04X)", linha, coluna, ch, ch)
	}
}

// lerTexto lê uma string literal entre aspas duplas.
func (l *Lexer) lerTexto() (string, error) {
	l.lerCaractere() // pular a aspas de abertura
	inicio := l.posicao

	for l.caractere != '"' {
		if l.caractere == 0 {
			return "", fmt.Errorf("linha %d: texto não fechado (faltou aspas de fechamento)", l.linha)
		}
		if l.caractere == '\\' {
			l.lerCaractere() // pular caractere de escape
		}
		l.lerCaractere()
	}

	valor := l.entrada[inicio:l.posicao]
	l.lerCaractere() // pular a aspas de fechamento
	return valor, nil
}

// lerNumero lê um literal numérico (inteiro ou decimal).
func (l *Lexer) lerNumero() string {
	inicio := l.posicao

	for unicode.IsDigit(l.caractere) {
		l.lerCaractere()
	}

	// Suporte a decimais (ex: 3.14)
	if l.caractere == '.' && unicode.IsDigit(l.espiarProximo()) {
		l.lerCaractere() // pular o ponto
		for unicode.IsDigit(l.caractere) {
			l.lerCaractere()
		}
	}

	return l.entrada[inicio:l.posicao]
}

// lerPalavra lê uma palavra (pode ser identificador, palavra-chave, ou tipo).
func (l *Lexer) lerPalavra() string {
	inicio := l.posicao

	for ehLetra(l.caractere) || unicode.IsDigit(l.caractere) || l.caractere == '_' {
		l.lerCaractere()
	}

	return l.entrada[inicio:l.posicao]
}

// ehLetra verifica se uma runa é uma letra válida para identificadores na linguagem Verbo.
// Inclui todo o range Unicode de letras (acentuação, cedilha, etc.)
func ehLetra(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}
