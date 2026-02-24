// Package ast define os nós da Árvore de Sintaxe Abstrata (AST) da linguagem Verbo.
// Cada construto da linguagem (declaração, expressão, comando) é representado
// por um tipo que implementa a interface No.
package ast

import "github.com/juanxto/crom-verbo/pkg/lexer"

// No é a interface base para todos os nós da AST.
type No interface {
	noNo() // método sentinela para tipagem
	TokenLiteral() string
}

// Declaracao representa qualquer instrução que não produz valor direto.
type Declaracao interface {
	No
	noDeclaracao()
}

// Expressao representa qualquer construto que produz um valor.
type Expressao interface {
	No
	noExpressao()
}

// -----------------------------------------------
// Nó raiz
// -----------------------------------------------

// Programa é o nó raiz da AST — contém todas as declarações do arquivo .vrb.
type Programa struct {
	Declaracoes []Declaracao
}

func (p *Programa) noNo()            {}
func (p *Programa) TokenLiteral() string {
	if len(p.Declaracoes) > 0 {
		return p.Declaracoes[0].TokenLiteral()
	}
	return ""
}

// -----------------------------------------------
// Declarações
// -----------------------------------------------

// DeclaracaoVariavel representa: "A mensagem é ..." ou "Um contador está ..."
type DeclaracaoVariavel struct {
	Token      lexer.Token   // O artigo (O/A/Um/Uma)
	Nome       string        // Nome da variável
	Mutavel    bool          // true se artigo indefinido (Um/Uma)
	Verbo      string        // "é" ou "está"
	Valor      Expressao     // Expressão do lado direito
}

func (d *DeclaracaoVariavel) noNo()            {}
func (d *DeclaracaoVariavel) noDeclaracao()     {}
func (d *DeclaracaoVariavel) TokenLiteral() string { return d.Token.Valor }

// DeclaracaoFuncao representa: "Para Calcular usando (param: Tipo):"
type DeclaracaoFuncao struct {
	Token      lexer.Token     // Token PARA
	Nome       string          // Nome da função (verbo no infinitivo)
	Parametros []Parametro     // Lista de parâmetros
	Corpo      *Bloco          // Corpo da função
}

func (d *DeclaracaoFuncao) noNo()            {}
func (d *DeclaracaoFuncao) noDeclaracao()     {}
func (d *DeclaracaoFuncao) TokenLiteral() string { return d.Token.Valor }

// Parametro representa um parâmetro de função com nome e tipo opcional.
type Parametro struct {
	Nome string
	Tipo string // Tipo opcional (Texto, Inteiro, etc.)
}

// DeclaracaoRetorne representa: "Retorne valor."
type DeclaracaoRetorne struct {
	Token lexer.Token // Token RETORNE
	Valor Expressao   // Valor a retornar (pode ser nil para Retorne Nulo)
}

func (d *DeclaracaoRetorne) noNo()            {}
func (d *DeclaracaoRetorne) noDeclaracao()     {}
func (d *DeclaracaoRetorne) TokenLiteral() string { return d.Token.Valor }

// DeclaracaoExibir representa: "Exibir com (expressão)."
type DeclaracaoExibir struct {
	Token lexer.Token // Token EXIBIR
	Valor Expressao   // Expressão a exibir
}

func (d *DeclaracaoExibir) noNo()            {}
func (d *DeclaracaoExibir) noDeclaracao()     {}
func (d *DeclaracaoExibir) TokenLiteral() string { return d.Token.Valor }

// DeclaracaoSe representa: "Se condição for ..., então:"
type DeclaracaoSe struct {
	Token      lexer.Token   // Token SE
	Condicao   Expressao     // Expressão condicional
	Consequencia *Bloco      // Bloco "então"
	Alternativa  *Bloco      // Bloco "Senão" (pode ser nil)
}

func (d *DeclaracaoSe) noNo()            {}
func (d *DeclaracaoSe) noDeclaracao()     {}
func (d *DeclaracaoSe) TokenLiteral() string { return d.Token.Valor }

// DeclaracaoRepita representa: "Repita N vezes:" ou "Repita para cada item em lista:"
type DeclaracaoRepita struct {
	Token      lexer.Token // Token REPITA
	Contagem   Expressao   // Número de repetições (para "Repita N vezes")
	Variavel   string      // Variável de iteração (para "Repita para cada")
	Iteravel   Expressao   // Expressão iterável (para "Repita para cada")
	ForEach    bool        // true se for "para cada", false se for "N vezes"
	Corpo      *Bloco      // Corpo do loop
}

func (d *DeclaracaoRepita) noNo()            {}
func (d *DeclaracaoRepita) noDeclaracao()     {}
func (d *DeclaracaoRepita) TokenLiteral() string { return d.Token.Valor }

// DeclaracaoEnquanto representa: "Enquanto condição:"
type DeclaracaoEnquanto struct {
	Token    lexer.Token // Token ENQUANTO
	Condicao Expressao   // Expressão condicional
	Corpo    *Bloco      // Corpo do loop
}

func (d *DeclaracaoEnquanto) noNo()            {}
func (d *DeclaracaoEnquanto) noDeclaracao()     {}
func (d *DeclaracaoEnquanto) TokenLiteral() string { return d.Token.Valor }

// DeclaracaoAtribuicao representa reatribuição: "variável está novo_valor."
type DeclaracaoAtribuicao struct {
	Token  lexer.Token // Token do identificador
	Nome   string      // Nome da variável
	Valor  Expressao   // Novo valor
}

func (d *DeclaracaoAtribuicao) noNo()            {}
func (d *DeclaracaoAtribuicao) noDeclaracao()     {}
func (d *DeclaracaoAtribuicao) TokenLiteral() string { return d.Token.Valor }

// DeclaracaoExpressao encapsula uma expressão usada como declaração
// (ex: chamada de função como instrução).
type DeclaracaoExpressao struct {
	Token     lexer.Token
	Expressao Expressao
}

func (d *DeclaracaoExpressao) noNo()            {}
func (d *DeclaracaoExpressao) noDeclaracao()     {}
func (d *DeclaracaoExpressao) TokenLiteral() string { return d.Token.Valor }

// Bloco representa um bloco de declarações (corpo de função, if, loop).
type Bloco struct {
	Declaracoes []Declaracao
}

func (b *Bloco) noNo()            {}
func (b *Bloco) TokenLiteral() string { return "{bloco}" }

// -----------------------------------------------
// Expressões
// -----------------------------------------------

// ExpressaoLiteralNumero representa um literal numérico: 42, 3.14
type ExpressaoLiteralNumero struct {
	Token lexer.Token
	Valor string
}

func (e *ExpressaoLiteralNumero) noNo()            {}
func (e *ExpressaoLiteralNumero) noExpressao()      {}
func (e *ExpressaoLiteralNumero) TokenLiteral() string { return e.Token.Valor }

// ExpressaoLiteralTexto representa um literal de texto: "Olá"
type ExpressaoLiteralTexto struct {
	Token lexer.Token
	Valor string
}

func (e *ExpressaoLiteralTexto) noNo()            {}
func (e *ExpressaoLiteralTexto) noExpressao()      {}
func (e *ExpressaoLiteralTexto) TokenLiteral() string { return e.Token.Valor }

// ExpressaoLiteralLogico representa Verdadeiro ou Falso.
type ExpressaoLiteralLogico struct {
	Token lexer.Token
	Valor bool
}

func (e *ExpressaoLiteralLogico) noNo()            {}
func (e *ExpressaoLiteralLogico) noExpressao()      {}
func (e *ExpressaoLiteralLogico) TokenLiteral() string { return e.Token.Valor }

// ExpressaoNulo representa o valor Nulo.
type ExpressaoNulo struct {
	Token lexer.Token
}

func (e *ExpressaoNulo) noNo()            {}
func (e *ExpressaoNulo) noExpressao()      {}
func (e *ExpressaoNulo) TokenLiteral() string { return e.Token.Valor }

// ExpressaoIdentificador representa uma referência a uma variável pelo nome.
type ExpressaoIdentificador struct {
	Token lexer.Token
	Nome  string
}

func (e *ExpressaoIdentificador) noNo()            {}
func (e *ExpressaoIdentificador) noExpressao()      {}
func (e *ExpressaoIdentificador) TokenLiteral() string { return e.Token.Valor }

// ExpressaoBinaria representa uma operação entre duas expressões.
// Ex: "contador + 1", "idade menor que 18"
type ExpressaoBinaria struct {
	Token    lexer.Token // Token do operador
	Esquerda Expressao
	Operador string     // "+", "-", "*", "/", "menor que", "maior que", "igual", "e"
	Direita  Expressao
}

func (e *ExpressaoBinaria) noNo()            {}
func (e *ExpressaoBinaria) noExpressao()      {}
func (e *ExpressaoBinaria) TokenLiteral() string { return e.Token.Valor }

// ExpressaoUnaria representa uma operação com um operando.
// Ex: "não ativo"
type ExpressaoUnaria struct {
	Token    lexer.Token
	Operador string    // "não", "-"
	Operando Expressao
}

func (e *ExpressaoUnaria) noNo()            {}
func (e *ExpressaoUnaria) noExpressao()      {}
func (e *ExpressaoUnaria) TokenLiteral() string { return e.Token.Valor }

// ExpressaoChamadaFuncao representa uma chamada de função.
// Ex: "Calcular com (10, 20)" ou "Saudar com (nome)"
type ExpressaoChamadaFuncao struct {
	Token      lexer.Token
	Nome       string
	Argumentos []Expressao
}

func (e *ExpressaoChamadaFuncao) noNo()            {}
func (e *ExpressaoChamadaFuncao) noExpressao()      {}
func (e *ExpressaoChamadaFuncao) TokenLiteral() string { return e.Token.Valor }

// ExpressaoAgrupada representa uma expressão entre parênteses: (expr)
type ExpressaoAgrupada struct {
	Token     lexer.Token
	Expressao Expressao
}

func (e *ExpressaoAgrupada) noNo()            {}
func (e *ExpressaoAgrupada) noExpressao()      {}
func (e *ExpressaoAgrupada) TokenLiteral() string { return e.Token.Valor }
