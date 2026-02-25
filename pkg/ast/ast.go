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
	Objeto     Expressao   // Opcional, para chamadas de método: "obj de metodo com (args)"
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

// -----------------------------------------------
// V2: Declarações Avançadas
// -----------------------------------------------

// CampoEntidade representa um campo de uma Entidade (struct).
type CampoEntidade struct {
	Nome string
	Tipo string
}

// DeclaracaoEntidade representa: "A Entidade Usuario contendo (nome: Texto, idade: Inteiro)."
type DeclaracaoEntidade struct {
	Token  lexer.Token     // Token ARTIGO
	Nome   string          // Nome da entidade
	Campos []CampoEntidade // Campos da entidade
}

func (d *DeclaracaoEntidade) noNo()            {}
func (d *DeclaracaoEntidade) noDeclaracao()     {}
func (d *DeclaracaoEntidade) TokenLiteral() string { return d.Token.Valor }

// DeclaracaoSimultaneamente representa: "Simultaneamente:"
type DeclaracaoSimultaneamente struct {
	Token lexer.Token // Token SIMULTANEAMENTE
	Corpo *Bloco      // Bloco de declarações a executar em paralelo
}

func (d *DeclaracaoSimultaneamente) noNo()            {}
func (d *DeclaracaoSimultaneamente) noDeclaracao()     {}
func (d *DeclaracaoSimultaneamente) TokenLiteral() string { return d.Token.Valor }

// DeclaracaoTente representa: "Tente: ... Capture erro: ..."
type DeclaracaoTente struct {
	Token        lexer.Token // Token TENTE
	Tentativa    *Bloco      // Bloco "Tente"
	VariavelErro string      // Nome da variável de erro em "Capture"
	Captura      *Bloco      // Bloco "Capture" (pode ser nil)
}

func (d *DeclaracaoTente) noNo()            {}
func (d *DeclaracaoTente) noDeclaracao()     {}
func (d *DeclaracaoTente) TokenLiteral() string { return d.Token.Valor }

// DeclaracaoSinalize representa: "Sinalize com (mensagem)."
type DeclaracaoSinalize struct {
	Token lexer.Token
	Valor Expressao
}

func (d *DeclaracaoSinalize) noNo()            {}
func (d *DeclaracaoSinalize) noDeclaracao()     {}
func (d *DeclaracaoSinalize) TokenLiteral() string { return d.Token.Valor }

// -----------------------------------------------
// V2: Expressões Avançadas
// -----------------------------------------------

// ExpressaoLista representa: ["elem1", "elem2", "elem3"]
type ExpressaoLista struct {
	Token    lexer.Token
	Elementos []Expressao
}

func (e *ExpressaoLista) noNo()            {}
func (e *ExpressaoLista) noExpressao()      {}
func (e *ExpressaoLista) TokenLiteral() string { return e.Token.Valor }

// ExpressaoAcessoIndice representa: lista[0]
type ExpressaoAcessoIndice struct {
	Token  lexer.Token
	Objeto Expressao // A lista sendo acessada
	Indice Expressao // O índice
}

func (e *ExpressaoAcessoIndice) noNo()            {}
func (e *ExpressaoAcessoIndice) noExpressao()      {}
func (e *ExpressaoAcessoIndice) TokenLiteral() string { return e.Token.Valor }

// ExpressaoAcessoCampo representa: "nome de usuario" → usuario.Nome
type ExpressaoAcessoCampo struct {
	Token  lexer.Token
	Campo  string    // O campo sendo acessado
	Objeto Expressao // O objeto que contém o campo
}

func (e *ExpressaoAcessoCampo) noNo()            {}
func (e *ExpressaoAcessoCampo) noExpressao()      {}
func (e *ExpressaoAcessoCampo) TokenLiteral() string { return e.Token.Valor }

// ExpressaoInstanciacao representa: "Usuario com (nome: "Juan", idade: 25)"
type ExpressaoInstanciacao struct {
	Token      lexer.Token
	Tipo       string                   // Nome da entidade
	Argumentos []Expressao              // Valores dos campos
}

func (e *ExpressaoInstanciacao) noNo()            {}
func (e *ExpressaoInstanciacao) noExpressao()      {}
func (e *ExpressaoInstanciacao) TokenLiteral() string { return e.Token.Valor }

// -----------------------------------------------
// V2: Canais e Concorrência Avançada
// -----------------------------------------------

// DeclaracaoCanal representa a criação de um canal: "Um via é um Canal de Inteiros."
// Ou melhor, o tipo Canal é usado na declaração. "Canal" age como um tipo primitivo "Canal de Tipo"
// Mas a especificação inicial para parser sugeriu:
// "Uma via é um Canal de Inteiros."
// Sendo que "Canal de Inteiros" será tratado no parser como ExpressaoCanal.
// Vamos criar as expressões necessárias.

// ExpressaoCriarCanal representa: "Canal de Tipo"
type ExpressaoCriarCanal struct {
	Token    lexer.Token // O token "Canal"
	TipoItem string      // O tipo dos itens, ex: "Inteiros" -> "int"
}

func (ec *ExpressaoCriarCanal) noExpressao()          {}
func (ec *ExpressaoCriarCanal) noNo()                 {}
func (ec *ExpressaoCriarCanal) TokenLiteral() string  { return ec.Token.Valor }

// DeclaracaoEnviar representa enviar valor para canal: "Enviar 10 para via."
type DeclaracaoEnviar struct {
	Token lexer.Token // O token "Enviar"
	Valor Expressao   // O valor enviado
	Canal string      // Nome do canal
}

func (de *DeclaracaoEnviar) noDeclaracao()         {}
func (de *DeclaracaoEnviar) noNo()                 {}
func (de *DeclaracaoEnviar) TokenLiteral() string  { return de.Token.Valor }

// ExpressaoReceber representa ler valor do canal: "Receber de via"
type ExpressaoReceber struct {
	Token lexer.Token // O token "Receber"
	Canal string      // Nome do canal
}

func (er *ExpressaoReceber) noExpressao()          {}
func (er *ExpressaoReceber) noNo()                 {}
func (er *ExpressaoReceber) TokenLiteral() string  { return er.Token.Valor }

// -----------------------------------------------
// V2: Biblioteca Padrão
// -----------------------------------------------

// DeclaracaoIncluir representa a importação de um pacote padrão: "Incluir Matematica."
type DeclaracaoIncluir struct {
	Token  lexer.Token // O token "Incluir"
	Pacote string      // Nome do pacote a incluir
}

func (di *DeclaracaoIncluir) noDeclaracao()         {}
func (di *DeclaracaoIncluir) noNo()                 {}
func (di *DeclaracaoIncluir) TokenLiteral() string  { return di.Token.Valor }
