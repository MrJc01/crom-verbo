// Package lexer implementa a análise léxica (tokenização) da linguagem Verbo.
// Ele converte código-fonte .vrb em uma sequência de tokens tipados.
package lexer

import "fmt"

// TokenType representa o tipo de um token na linguagem Verbo.
type TokenType int

const (
	// Tokens especiais
	TOKEN_ILEGAL TokenType = iota
	TOKEN_FIM                // Fim do arquivo

	// Literais
	TOKEN_NUMERO             // 42, 3.14
	TOKEN_TEXTO              // "Olá, Mundo!"
	TOKEN_IDENTIFICADOR      // nomes de variáveis, funções

	// Artigos (determinam mutabilidade)
	TOKEN_ARTIGO_DEFINIDO    // O, A, Os, As
	TOKEN_ARTIGO_INDEFINIDO  // Um, Uma, Uns, Umas

	// Pronomes demonstrativos (referências)
	TOKEN_DEMONSTRATIVO      // Este, Esta, Aquele, Aquela

	// Verbos / Palavras-chave
	TOKEN_E_ACENTO           // É (atribuição estática)
	TOKEN_ESTA               // Está (atribuição de estado)
	TOKEN_PARA               // Para (declaração de função)
	TOKEN_USANDO             // usando (parâmetros de função)
	TOKEN_SE                 // Se
	TOKEN_FOR                // for (subjuntivo)
	TOKEN_ENTAO              // então
	TOKEN_SENAO              // Senão
	TOKEN_REPITA             // Repita
	TOKEN_VEZES              // vezes
	TOKEN_PARA_CADA          // cada
	TOKEN_EM                 // em
	TOKEN_RETORNE            // Retorne
	TOKEN_EXIBIR             // Exibir
	TOKEN_COM                // com
	TOKEN_ENQUANTO           // Enquanto
	TOKEN_VERDADEIRO         // Verdadeiro
	TOKEN_FALSO              // Falso
	TOKEN_NULO               // Nulo
	TOKEN_NAO                // não
	TOKEN_MENOR              // menor
	TOKEN_MAIOR              // maior
	TOKEN_IGUAL              // igual
	TOKEN_QUE                // que
	TOKEN_DE                 // de / do / da
	TOKEN_DADO               // Dado (premissa / guard clause)

	// Operadores
	TOKEN_MAIS               // +
	TOKEN_MENOS              // -
	TOKEN_MULTIPLICAR        // * ou "vezes" contextual
	TOKEN_DIVIDIR            // /
	TOKEN_ATRIBUIR           // = (usado internamente)

	// Delimitadores
	TOKEN_PONTO              // . (fim de instrução)
	TOKEN_DOIS_PONTOS        // : (início de bloco)
	TOKEN_VIRGULA            // ,
	TOKEN_PARENTESE_ABRE     // (
	TOKEN_PARENTESE_FECHA    // )
	TOKEN_COLCHETE_ABRE      // [
	TOKEN_COLCHETE_FECHA     // ]

	// Tipos
	TOKEN_TIPO               // Texto, Inteiro, Decimal, Logico, Lista
)

// Token representa um token individual produzido pelo Lexer.
type Token struct {
	Tipo    TokenType
	Valor   string
	Linha   int
	Coluna  int
}

// String retorna uma representação legível do token para debug.
func (t Token) String() string {
	return fmt.Sprintf("Token{%s, %q, L%d:C%d}", t.Tipo.NomeLegivel(), t.Valor, t.Linha, t.Coluna)
}

// NomeLegivel retorna o nome legível em português do tipo de token.
func (tt TokenType) NomeLegivel() string {
	nomes := map[TokenType]string{
		TOKEN_ILEGAL:           "ILEGAL",
		TOKEN_FIM:              "FIM",
		TOKEN_NUMERO:           "NÚMERO",
		TOKEN_TEXTO:            "TEXTO",
		TOKEN_IDENTIFICADOR:    "IDENTIFICADOR",
		TOKEN_ARTIGO_DEFINIDO:  "ARTIGO_DEFINIDO",
		TOKEN_ARTIGO_INDEFINIDO: "ARTIGO_INDEFINIDO",
		TOKEN_DEMONSTRATIVO:    "DEMONSTRATIVO",
		TOKEN_E_ACENTO:         "É",
		TOKEN_ESTA:             "ESTÁ",
		TOKEN_PARA:             "PARA",
		TOKEN_USANDO:           "USANDO",
		TOKEN_SE:               "SE",
		TOKEN_FOR:              "FOR",
		TOKEN_ENTAO:            "ENTÃO",
		TOKEN_SENAO:            "SENÃO",
		TOKEN_REPITA:           "REPITA",
		TOKEN_VEZES:            "VEZES",
		TOKEN_PARA_CADA:        "CADA",
		TOKEN_EM:               "EM",
		TOKEN_RETORNE:          "RETORNE",
		TOKEN_EXIBIR:           "EXIBIR",
		TOKEN_COM:              "COM",
		TOKEN_ENQUANTO:         "ENQUANTO",
		TOKEN_VERDADEIRO:       "VERDADEIRO",
		TOKEN_FALSO:            "FALSO",
		TOKEN_NULO:             "NULO",
		TOKEN_NAO:              "NÃO",
		TOKEN_MENOR:            "MENOR",
		TOKEN_MAIOR:            "MAIOR",
		TOKEN_IGUAL:            "IGUAL",
		TOKEN_QUE:              "QUE",
		TOKEN_DE:               "DE",
		TOKEN_DADO:             "DADO",
		TOKEN_MAIS:             "MAIS",
		TOKEN_MENOS:            "MENOS",
		TOKEN_MULTIPLICAR:      "MULTIPLICAR",
		TOKEN_DIVIDIR:          "DIVIDIR",
		TOKEN_ATRIBUIR:         "ATRIBUIR",
		TOKEN_PONTO:            "PONTO",
		TOKEN_DOIS_PONTOS:      "DOIS_PONTOS",
		TOKEN_VIRGULA:          "VÍRGULA",
		TOKEN_PARENTESE_ABRE:   "PARÊNTESE_ABRE",
		TOKEN_PARENTESE_FECHA:  "PARÊNTESE_FECHA",
		TOKEN_COLCHETE_ABRE:    "COLCHETE_ABRE",
		TOKEN_COLCHETE_FECHA:   "COLCHETE_FECHA",
		TOKEN_TIPO:             "TIPO",
	}
	if nome, ok := nomes[tt]; ok {
		return nome
	}
	return fmt.Sprintf("DESCONHECIDO(%d)", int(tt))
}

// palavrasChave mapeia palavras reservadas da linguagem Verbo para seus TokenTypes.
var palavrasChave = map[string]TokenType{
	// Artigos definidos
	"O":          TOKEN_ARTIGO_DEFINIDO,
	"A":          TOKEN_ARTIGO_DEFINIDO,
	"Os":         TOKEN_ARTIGO_DEFINIDO,
	"As":         TOKEN_ARTIGO_DEFINIDO,
	"o":          TOKEN_ARTIGO_DEFINIDO,
	"a":          TOKEN_ARTIGO_DEFINIDO,
	"os":         TOKEN_ARTIGO_DEFINIDO,
	"as":         TOKEN_ARTIGO_DEFINIDO,

	// Artigos indefinidos
	"Um":         TOKEN_ARTIGO_INDEFINIDO,
	"Uma":        TOKEN_ARTIGO_INDEFINIDO,
	"um":         TOKEN_ARTIGO_INDEFINIDO,
	"uma":        TOKEN_ARTIGO_INDEFINIDO,

	// Demonstrativos
	"Este":       TOKEN_DEMONSTRATIVO,
	"Esta":       TOKEN_DEMONSTRATIVO,
	"Aquele":     TOKEN_DEMONSTRATIVO,
	"Aquela":     TOKEN_DEMONSTRATIVO,
	"este":       TOKEN_DEMONSTRATIVO,
	"esta":       TOKEN_DEMONSTRATIVO,
	"aquele":     TOKEN_DEMONSTRATIVO,
	"aquela":     TOKEN_DEMONSTRATIVO,

	// Palavras-chave semânticas
	"É":          TOKEN_E_ACENTO,
	"é":          TOKEN_E_ACENTO,
	"Está":       TOKEN_ESTA,
	"está":       TOKEN_ESTA,
	"Para":       TOKEN_PARA,
	"para":       TOKEN_PARA,
	"usando":     TOKEN_USANDO,
	"Se":         TOKEN_SE,
	"se":         TOKEN_SE,
	"for":        TOKEN_FOR,
	"então":      TOKEN_ENTAO,
	"Então":      TOKEN_ENTAO,
	"Senão":      TOKEN_SENAO,
	"senão":      TOKEN_SENAO,
	"Repita":     TOKEN_REPITA,
	"repita":     TOKEN_REPITA,
	"vezes":      TOKEN_VEZES,
	"cada":       TOKEN_PARA_CADA,
	"em":         TOKEN_EM,
	"Retorne":    TOKEN_RETORNE,
	"retorne":    TOKEN_RETORNE,
	"Exibir":     TOKEN_EXIBIR,
	"exibir":     TOKEN_EXIBIR,
	"com":        TOKEN_COM,
	"Enquanto":   TOKEN_ENQUANTO,
	"enquanto":   TOKEN_ENQUANTO,
	"Verdadeiro": TOKEN_VERDADEIRO,
	"verdadeiro": TOKEN_VERDADEIRO,
	"Falso":      TOKEN_FALSO,
	"falso":      TOKEN_FALSO,
	"Nulo":       TOKEN_NULO,
	"nulo":       TOKEN_NULO,
	"não":        TOKEN_NAO,
	"Não":        TOKEN_NAO,
	"menor":      TOKEN_MENOR,
	"maior":      TOKEN_MAIOR,
	"igual":      TOKEN_IGUAL,
	"que":        TOKEN_QUE,
	"de":         TOKEN_DE,
	"do":         TOKEN_DE,
	"da":         TOKEN_DE,
	"dos":        TOKEN_DE,
	"das":        TOKEN_DE,
	"Dado":       TOKEN_DADO,
	"dado":       TOKEN_DADO,

	// Operadores textuais
	"e":          TOKEN_MAIS,  // concatenação / adição contextual
	"menos":      TOKEN_MENOS,
	"mais":       TOKEN_MAIS,

	// Tipos
	"Texto":      TOKEN_TIPO,
	"Inteiro":    TOKEN_TIPO,
	"Decimal":    TOKEN_TIPO,
	"Logico":     TOKEN_TIPO,
	"Lógico":     TOKEN_TIPO,
	"Lista":      TOKEN_TIPO,
}

// BuscarPalavraChave verifica se uma palavra é reservada e retorna o TokenType correspondente.
// Se não for palavra reservada, retorna TOKEN_IDENTIFICADOR.
func BuscarPalavraChave(palavra string) TokenType {
	if tipo, ok := palavrasChave[palavra]; ok {
		return tipo
	}
	return TOKEN_IDENTIFICADOR
}
