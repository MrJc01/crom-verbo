// Package parser implementa o analisador sintático da linguagem Verbo.
// Utiliza Recursive Descent Parsing para construir a AST a partir dos tokens.
package parser

import (
	"fmt"
	"strings"

	"github.com/juanxto/crom-verbo/pkg/ast"
	"github.com/juanxto/crom-verbo/pkg/lexer"
)

// Parser é o analisador sintático da linguagem Verbo.
type Parser struct {
	tokens            []lexer.Token
	posicao           int
	erros             []string
	nivelProfundidade int // rastreia nível de aninhamento de blocos
}

// Novo cria um novo Parser a partir de uma lista de tokens.
func Novo(tokens []lexer.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		posicao: 0,
	}
}

// Analisar processa todos os tokens e retorna o programa (AST raiz).
func (p *Parser) Analisar() (*ast.Programa, error) {
	programa := &ast.Programa{}

	for !p.fimDoArquivo() {
		decl := p.analisarDeclaracao()
		if decl != nil {
			programa.Declaracoes = append(programa.Declaracoes, decl)
		}
	}

	if len(p.erros) > 0 {
		return programa, fmt.Errorf("erros sintáticos:\n%s", strings.Join(p.erros, "\n"))
	}

	return programa, nil
}

// Erros retorna a lista de erros encontrados durante a análise.
func (p *Parser) Erros() []string {
	return p.erros
}

// -----------------------------------------------
// Helpers de navegação
// -----------------------------------------------

func (p *Parser) tokenAtual() lexer.Token {
	if p.posicao >= len(p.tokens) {
		return lexer.Token{Tipo: lexer.TOKEN_FIM}
	}
	return p.tokens[p.posicao]
}

func (p *Parser) espiar() lexer.Token {
	pos := p.posicao + 1
	if pos >= len(p.tokens) {
		return lexer.Token{Tipo: lexer.TOKEN_FIM}
	}
	return p.tokens[pos]
}

func (p *Parser) avancar() lexer.Token {
	tok := p.tokenAtual()
	p.posicao++
	return tok
}

func (p *Parser) fimDoArquivo() bool {
	return p.tokenAtual().Tipo == lexer.TOKEN_FIM
}

func (p *Parser) esperarTipo(tipo lexer.TokenType) (lexer.Token, bool) {
	if p.tokenAtual().Tipo == tipo {
		return p.avancar(), true
	}
	p.erroEsperado(tipo)
	return p.tokenAtual(), false
}

func (p *Parser) erroEsperado(tipo lexer.TokenType) {
	tok := p.tokenAtual()
	p.erros = append(p.erros, fmt.Sprintf(
		"linha %d, coluna %d: esperava %s, encontrou %s (%q)",
		tok.Linha, tok.Coluna, tipo.NomeLegivel(), tok.Tipo.NomeLegivel(), tok.Valor,
	))
}

func (p *Parser) erro(msg string) {
	tok := p.tokenAtual()
	p.erros = append(p.erros, fmt.Sprintf(
		"linha %d, coluna %d: %s",
		tok.Linha, tok.Coluna, msg,
	))
}

// consumirPonto consome um ponto final opcional (fim de instrução).
func (p *Parser) consumirPonto() {
	if p.tokenAtual().Tipo == lexer.TOKEN_PONTO {
		p.avancar()
	}
}

// -----------------------------------------------
// Análise de Declarações
// -----------------------------------------------

func (p *Parser) analisarDeclaracao() ast.Declaracao {
	switch p.tokenAtual().Tipo {
	case lexer.TOKEN_ARTIGO_DEFINIDO, lexer.TOKEN_ARTIGO_INDEFINIDO:
		// Verificar se é "A Entidade ..." (declaração de struct)
		if p.espiar().Tipo == lexer.TOKEN_ENTIDADE {
			return p.analisarDeclaracaoEntidade()
		}
		return p.analisarDeclaracaoVariavel()

	case lexer.TOKEN_PARA:
		return p.analisarDeclaracaoFuncao()

	case lexer.TOKEN_SE:
		return p.analisarDeclaracaoSe()

	case lexer.TOKEN_REPITA:
		return p.analisarDeclaracaoRepita()

	case lexer.TOKEN_ENQUANTO:
		return p.analisarDeclaracaoEnquanto()

	case lexer.TOKEN_EXIBIR:
		return p.analisarDeclaracaoExibir()

	case lexer.TOKEN_RETORNE:
		return p.analisarDeclaracaoRetorne()

	case lexer.TOKEN_SIMULTANEAMENTE:
		return p.analisarDeclaracaoSimultaneamente()

	case lexer.TOKEN_TENTE:
		return p.analisarDeclaracaoTente()

	case lexer.TOKEN_SINALIZE:
		return p.analisarDeclaracaoSinalize()

	case lexer.TOKEN_ENVIAR:
		return p.analisarDeclaracaoEnviar()

	case lexer.TOKEN_INCLUIR:
		return p.analisarDeclaracaoIncluir()

	case lexer.TOKEN_SENAO:
		return nil

	case lexer.TOKEN_CAPTURE:
		// Capture é tratado dentro do Tente
		return nil

	case lexer.TOKEN_IDENTIFICADOR:
		return p.analisarDeclaracaoIdentificador()

	default:
		p.erro(fmt.Sprintf("declaração inesperada: %s (%q)", p.tokenAtual().Tipo.NomeLegivel(), p.tokenAtual().Valor))
		p.avancar()
		return nil
	}
}

// analisarDeclaracaoVariavel analisa: "A/O/Um/Uma nome é/está valor."
func (p *Parser) analisarDeclaracaoVariavel() ast.Declaracao {
	artigo := p.avancar()
	mutavel := artigo.Tipo == lexer.TOKEN_ARTIGO_INDEFINIDO

	if p.tokenAtual().Tipo != lexer.TOKEN_IDENTIFICADOR {
		p.erroEsperado(lexer.TOKEN_IDENTIFICADOR)
		p.avancar()
		return nil
	}
	nome := p.avancar().Valor

	// Verificar se é "é" ou "está"
	var verbo string
	switch p.tokenAtual().Tipo {
	case lexer.TOKEN_E_ACENTO:
		verbo = "é"
		p.avancar()
	case lexer.TOKEN_ESTA:
		verbo = "está"
		p.avancar()
	default:
		p.erro(fmt.Sprintf("esperava 'é' ou 'está' após '%s', encontrou '%s'", nome, p.tokenAtual().Valor))
		p.avancar()
		return nil
	}

	valor := p.analisarExpressao()
	p.consumirPonto()

	return &ast.DeclaracaoVariavel{
		Token:   artigo,
		Nome:    nome,
		Mutavel: mutavel,
		Verbo:   verbo,
		Valor:   valor,
	}
}

// analisarDeclaracaoFuncao analisa: "Para NomeFuncao usando (params):"
func (p *Parser) analisarDeclaracaoFuncao() ast.Declaracao {
	tokenPara := p.avancar() // consome "Para"

	// Nome da função
	if p.tokenAtual().Tipo != lexer.TOKEN_IDENTIFICADOR {
		p.erroEsperado(lexer.TOKEN_IDENTIFICADOR)
		return nil
	}
	nome := p.avancar().Valor

	// Parâmetros opcionais
	var parametros []ast.Parametro
	if p.tokenAtual().Tipo == lexer.TOKEN_USANDO {
		p.avancar() // consumir "usando"
		parametros = p.analisarParametros()
	}

	// Dois pontos (início do bloco)
	p.esperarTipo(lexer.TOKEN_DOIS_PONTOS)

	// Corpo da função
	corpo := p.analisarBloco()

	return &ast.DeclaracaoFuncao{
		Token:      tokenPara,
		Nome:       nome,
		Parametros: parametros,
		Corpo:      corpo,
	}
}

// analisarParametros analisa: (nome: Tipo, nome2: Tipo2)
func (p *Parser) analisarParametros() []ast.Parametro {
	var params []ast.Parametro

	if _, ok := p.esperarTipo(lexer.TOKEN_PARENTESE_ABRE); !ok {
		return params
	}

	for p.tokenAtual().Tipo != lexer.TOKEN_PARENTESE_FECHA && !p.fimDoArquivo() {
		if p.tokenAtual().Tipo == lexer.TOKEN_VIRGULA {
			p.avancar()
			continue
		}

		param := ast.Parametro{}

		if p.tokenAtual().Tipo != lexer.TOKEN_IDENTIFICADOR {
			p.erroEsperado(lexer.TOKEN_IDENTIFICADOR)
			p.avancar()
			continue
		}
		param.Nome = p.avancar().Valor

		// Tipo opcional: "nome: Tipo"
		if p.tokenAtual().Tipo == lexer.TOKEN_DOIS_PONTOS {
			p.avancar()
			if p.tokenAtual().Tipo == lexer.TOKEN_TIPO || p.tokenAtual().Tipo == lexer.TOKEN_IDENTIFICADOR || p.tokenAtual().Tipo == lexer.TOKEN_CANAL {
				param.Tipo = p.avancar().Valor
			}
		}

		params = append(params, param)
	}

	p.esperarTipo(lexer.TOKEN_PARENTESE_FECHA)
	return params
}

// analisarDeclaracaoSe analisa: "Se condição, então:" com "Senão:" opcional.
func (p *Parser) analisarDeclaracaoSe() ast.Declaracao {
	tokenSe := p.avancar() // consome "Se"

	condicao := p.analisarExpressaoCondicional()

	// Consumir "então" e ":"
	// Skip comma before "então" if present
	if p.tokenAtual().Tipo == lexer.TOKEN_VIRGULA {
		p.avancar()
	}
	if p.tokenAtual().Tipo == lexer.TOKEN_ENTAO {
		p.avancar()
	}
	p.esperarTipo(lexer.TOKEN_DOIS_PONTOS)

	consequencia := p.analisarBloco()

	var alternativa *ast.Bloco
	if p.tokenAtual().Tipo == lexer.TOKEN_SENAO {
		p.avancar() // consome "Senão"
		if p.tokenAtual().Tipo == lexer.TOKEN_DOIS_PONTOS {
			p.avancar()
		}
		alternativa = p.analisarBloco()
	}

	return &ast.DeclaracaoSe{
		Token:        tokenSe,
		Condicao:     condicao,
		Consequencia: consequencia,
		Alternativa:  alternativa,
	}
}

// analisarExpressaoCondicional analisa condições como: "a idade for menor que 18"
func (p *Parser) analisarExpressaoCondicional() ast.Expressao {
	// Pular artigo antes do sujeito se presente
	if p.tokenAtual().Tipo == lexer.TOKEN_ARTIGO_DEFINIDO || p.tokenAtual().Tipo == lexer.TOKEN_ARTIGO_INDEFINIDO {
		p.avancar()
	}

	esquerda := p.analisarExpressaoPrimaria()

	// Verificar "for" (subjuntivo)
	if p.tokenAtual().Tipo == lexer.TOKEN_FOR {
		p.avancar()
	}

	// Operador de comparação: "menor que", "maior que", "igual"
	var operador string
	switch p.tokenAtual().Tipo {
	case lexer.TOKEN_MENOR:
		p.avancar()
		if p.tokenAtual().Tipo == lexer.TOKEN_QUE {
			p.avancar()
		}
		operador = "menor que"
	case lexer.TOKEN_MAIOR:
		p.avancar()
		if p.tokenAtual().Tipo == lexer.TOKEN_QUE {
			p.avancar()
		}
		operador = "maior que"
	case lexer.TOKEN_IGUAL:
		p.avancar()
		operador = "igual"
	default:
		return esquerda
	}

	direita := p.analisarExpressaoPrimaria()

	return &ast.ExpressaoBinaria{
		Token:    p.tokenAtual(),
		Esquerda: esquerda,
		Operador: operador,
		Direita:  direita,
	}
}

// analisarDeclaracaoRepita analisa: "Repita N vezes:" ou "Repita para cada X em lista:"
func (p *Parser) analisarDeclaracaoRepita() ast.Declaracao {
	tokenRepita := p.avancar() // consome "Repita"

	// Verificar se é "para cada" ou "N vezes"
	if p.tokenAtual().Tipo == lexer.TOKEN_PARA {
		// "Repita para cada item em lista:"
		p.avancar() // consome "para"
		if p.tokenAtual().Tipo == lexer.TOKEN_PARA_CADA {
			p.avancar() // consome "cada"
		}

		variavel := ""
		if p.tokenAtual().Tipo == lexer.TOKEN_IDENTIFICADOR {
			variavel = p.avancar().Valor
		}

		if p.tokenAtual().Tipo == lexer.TOKEN_EM {
			p.avancar() // consome "em"
		}

		iteravel := p.analisarExpressaoPrimaria()
		p.esperarTipo(lexer.TOKEN_DOIS_PONTOS)

		corpo := p.analisarBloco()

		return &ast.DeclaracaoRepita{
			Token:    tokenRepita,
			Variavel: variavel,
			Iteravel: iteravel,
			ForEach:  true,
			Corpo:    corpo,
		}
	}

	// "Repita N vezes:"
	contagem := p.analisarExpressaoPrimaria()

	if p.tokenAtual().Tipo == lexer.TOKEN_VEZES {
		p.avancar()
	}

	p.esperarTipo(lexer.TOKEN_DOIS_PONTOS)

	corpo := p.analisarBloco()

	return &ast.DeclaracaoRepita{
		Token:    tokenRepita,
		Contagem: contagem,
		ForEach:  false,
		Corpo:    corpo,
	}
}

// analisarDeclaracaoEnquanto analisa: "Enquanto condição:"
func (p *Parser) analisarDeclaracaoEnquanto() ast.Declaracao {
	tokenEnquanto := p.avancar() // consome "Enquanto"

	condicao := p.analisarExpressaoCondicional()

	p.esperarTipo(lexer.TOKEN_DOIS_PONTOS)

	corpo := p.analisarBloco()

	return &ast.DeclaracaoEnquanto{
		Token:    tokenEnquanto,
		Condicao: condicao,
		Corpo:    corpo,
	}
}

// analisarDeclaracaoExibir analisa: "Exibir com (expressão)."
func (p *Parser) analisarDeclaracaoExibir() ast.Declaracao {
	tokenExibir := p.avancar() // consome "Exibir"

	// "com" é opcional
	if p.tokenAtual().Tipo == lexer.TOKEN_COM {
		p.avancar()
	}

	var valor ast.Expressao
	if p.tokenAtual().Tipo == lexer.TOKEN_PARENTESE_ABRE {
		p.avancar()
		valor = p.analisarExpressao()
		p.esperarTipo(lexer.TOKEN_PARENTESE_FECHA)
	} else {
		valor = p.analisarExpressao()
	}

	p.consumirPonto()

	return &ast.DeclaracaoExibir{
		Token: tokenExibir,
		Valor: valor,
	}
}

// analisarDeclaracaoRetorne analisa: "Retorne valor."
func (p *Parser) analisarDeclaracaoRetorne() ast.Declaracao {
	tokenRetorne := p.avancar() // consome "Retorne"

	var valor ast.Expressao
	if p.tokenAtual().Tipo != lexer.TOKEN_PONTO && p.tokenAtual().Tipo != lexer.TOKEN_FIM {
		if p.tokenAtual().Tipo == lexer.TOKEN_NULO {
			valor = &ast.ExpressaoNulo{Token: p.avancar()}
		} else {
			valor = p.analisarExpressao()
		}
	}

	p.consumirPonto()

	return &ast.DeclaracaoRetorne{
		Token: tokenRetorne,
		Valor: valor,
	}
}

// analisarDeclaracaoIdentificador trata identificadores como reatribuição ou chamada de função.
func (p *Parser) analisarDeclaracaoIdentificador() ast.Declaracao {
	tok := p.avancar()

	// Reatribuição: "variável está novo_valor."
	if p.tokenAtual().Tipo == lexer.TOKEN_ESTA {
		p.avancar()
		valor := p.analisarExpressao()
		p.consumirPonto()
		return &ast.DeclaracaoAtribuicao{
			Token: tok,
			Nome:  tok.Valor,
			Valor: valor,
		}
	}

	// Chamada de função: "Funcao com (args)" ou "Funcao obj para obj"
	if p.tokenAtual().Tipo == lexer.TOKEN_COM || p.ehInicioExpressao(p.espiarComConectivos()) {
		if p.tokenAtual().Tipo == lexer.TOKEN_COM {
			p.avancar()
		}
		args := p.analisarArgumentos()
		p.consumirPonto()
		return &ast.DeclaracaoExpressao{
			Token: tok,
			Expressao: &ast.ExpressaoChamadaFuncao{
				Token:      tok,
				Nome:       tok.Valor,
				Argumentos: args,
			},
		}
	}

	// Expressão standalone
	expr := &ast.ExpressaoIdentificador{Token: tok, Nome: tok.Valor}
	p.consumirPonto()
	return &ast.DeclaracaoExpressao{
		Token:     tok,
		Expressao: expr,
	}
}

// -----------------------------------------------
// V2: Parsers Avançados
// -----------------------------------------------

// analisarDeclaracaoEntidade analisa: "A Entidade Usuario contendo (nome: Texto, idade: Inteiro)."
func (p *Parser) analisarDeclaracaoEntidade() ast.Declaracao {
	artigo := p.avancar() // consome artigo (A/O)
	p.avancar()           // consome "Entidade"

	if p.tokenAtual().Tipo != lexer.TOKEN_IDENTIFICADOR {
		p.erroEsperado(lexer.TOKEN_IDENTIFICADOR)
		return nil
	}
	nome := p.avancar().Valor

	// "contendo" é opcional
	if p.tokenAtual().Tipo == lexer.TOKEN_CONTENDO {
		p.avancar()
	}

	campos := p.analisarCamposEntidade()
	p.consumirPonto()

	return &ast.DeclaracaoEntidade{
		Token:  artigo,
		Nome:   nome,
		Campos: campos,
	}
}

// analisarCamposEntidade analisa: (nome: Tipo, idade: Inteiro)
func (p *Parser) analisarCamposEntidade() []ast.CampoEntidade {
	var campos []ast.CampoEntidade

	if _, ok := p.esperarTipo(lexer.TOKEN_PARENTESE_ABRE); !ok {
		return campos
	}

	for p.tokenAtual().Tipo != lexer.TOKEN_PARENTESE_FECHA && !p.fimDoArquivo() {
		if p.tokenAtual().Tipo == lexer.TOKEN_VIRGULA {
			p.avancar()
			continue
		}

		campo := ast.CampoEntidade{}
		if p.tokenAtual().Tipo != lexer.TOKEN_IDENTIFICADOR {
			p.erroEsperado(lexer.TOKEN_IDENTIFICADOR)
			p.avancar()
			continue
		}
		campo.Nome = p.avancar().Valor

		if p.tokenAtual().Tipo == lexer.TOKEN_DOIS_PONTOS {
			p.avancar()
			if p.tokenAtual().Tipo == lexer.TOKEN_TIPO || p.tokenAtual().Tipo == lexer.TOKEN_IDENTIFICADOR || p.tokenAtual().Tipo == lexer.TOKEN_CANAL {
				campo.Tipo = p.avancar().Valor
			}
		}

		campos = append(campos, campo)
	}

	p.esperarTipo(lexer.TOKEN_PARENTESE_FECHA)
	return campos
}

// analisarDeclaracaoSimultaneamente analisa: "Simultaneamente:"
func (p *Parser) analisarDeclaracaoSimultaneamente() ast.Declaracao {
	tok := p.avancar() // consome "Simultaneamente"
	p.esperarTipo(lexer.TOKEN_DOIS_PONTOS)
	corpo := p.analisarBloco()
	return &ast.DeclaracaoSimultaneamente{Token: tok, Corpo: corpo}
}

// analisarDeclaracaoTente analisa: "Tente: ... Capture erro: ..."
func (p *Parser) analisarDeclaracaoTente() ast.Declaracao {
	tok := p.avancar() // consome "Tente"
	p.esperarTipo(lexer.TOKEN_DOIS_PONTOS)

	tentativa := p.analisarBlocoAte(lexer.TOKEN_CAPTURE)

	var varErro string
	var captura *ast.Bloco

	if p.tokenAtual().Tipo == lexer.TOKEN_CAPTURE {
		p.avancar() // consome "Capture"
		if p.tokenAtual().Tipo == lexer.TOKEN_IDENTIFICADOR {
			varErro = p.avancar().Valor
		}
		if p.tokenAtual().Tipo == lexer.TOKEN_DOIS_PONTOS {
			p.avancar()
		}
		captura = p.analisarBloco()
	}

	return &ast.DeclaracaoTente{
		Token:        tok,
		Tentativa:    tentativa,
		VariavelErro: varErro,
		Captura:      captura,
	}
}

// analisarBlocoAte analisa um bloco até encontrar um token específico.
func (p *Parser) analisarBlocoAte(ate lexer.TokenType) *ast.Bloco {
	bloco := &ast.Bloco{}
	for !p.fimDoArquivo() {
		if p.tokenAtual().Tipo == ate {
			break
		}
		decl := p.analisarDeclaracao()
		if decl != nil {
			bloco.Declaracoes = append(bloco.Declaracoes, decl)
		}
	}
	return bloco
}

// analisarDeclaracaoSinalize analisa: "Sinalize com (mensagem)."
func (p *Parser) analisarDeclaracaoSinalize() ast.Declaracao {
	tok := p.avancar() // consome "Sinalize"
	if p.tokenAtual().Tipo == lexer.TOKEN_COM {
		p.avancar()
	}
	var valor ast.Expressao
	if p.tokenAtual().Tipo == lexer.TOKEN_PARENTESE_ABRE {
		p.avancar()
		valor = p.analisarExpressao()
		p.esperarTipo(lexer.TOKEN_PARENTESE_FECHA)
	} else {
		valor = p.analisarExpressao()
	}
	p.consumirPonto()
	return &ast.DeclaracaoSinalize{Token: tok, Valor: valor}
}

// analisarExpressaoLista analisa: [elem1, elem2, ...]
func (p *Parser) analisarExpressaoLista() ast.Expressao {
	tok := p.avancar() // consome "["
	var elementos []ast.Expressao

	for p.tokenAtual().Tipo != lexer.TOKEN_COLCHETE_FECHA && !p.fimDoArquivo() {
		if p.tokenAtual().Tipo == lexer.TOKEN_VIRGULA {
			p.avancar()
			continue
		}
		elem := p.analisarExpressao()
		if elem != nil {
			elementos = append(elementos, elem)
		}
	}

	p.esperarTipo(lexer.TOKEN_COLCHETE_FECHA)
	return &ast.ExpressaoLista{Token: tok, Elementos: elementos}
}

// analisarAcessoIndice analisa: obj[indice]
func (p *Parser) analisarAcessoIndice(obj ast.Expressao) ast.Expressao {
	tok := p.avancar() // consome "["
	indice := p.analisarExpressao()
	p.esperarTipo(lexer.TOKEN_COLCHETE_FECHA)
	return &ast.ExpressaoAcessoIndice{Token: tok, Objeto: obj, Indice: indice}
}

// -----------------------------------------------
// Análise de Blocos
// -----------------------------------------------

// analisarBloco analisa um bloco indentado de declarações.
// Um bloco termina quando encontramos uma palavra-chave de nível superior
// ou outro construto que indica fim do bloco.
func (p *Parser) analisarBloco() *ast.Bloco {
	p.nivelProfundidade++
	defer func() { p.nivelProfundidade-- }()

	bloco := &ast.Bloco{}
	
	for !p.fimDoArquivo() {
		tok := p.tokenAtual()

		// Fim do bloco — tokens que indicam saída
		if tok.Tipo == lexer.TOKEN_SENAO {
			break
		}

		// Se o usuário colocar um ponto solto na linha de baixo para fechar o bloco explícitamente: ". "
		if tok.Tipo == lexer.TOKEN_PONTO {
			p.avancar()
			break
		}

		// Quebra automática caso encontremos uma declaração de nível superior (ex: nova função)
		if p.ehInicioDeclaracaoNivelSuperior() {
			break
		}

		decl := p.analisarDeclaracao()
		if decl != nil {
			bloco.Declaracoes = append(bloco.Declaracoes, decl)
		}
	}

	return bloco
}

// ehInicioDeclaracaoNivelSuperior verifica se o token atual inicia uma declaração de nível superior.
// Quando estamos dentro de um bloco (nivelProfundidade > 1), padrões como artigos,
// Exibir, e chamadas de função indicam que saímos do bloco.
func (p *Parser) ehInicioDeclaracaoNivelSuperior() bool {
	tok := p.tokenAtual()

	// "Para" seguido de identificador = nova função (SEMPRE é nível superior, em qualquer profundidade)
	if tok.Tipo == lexer.TOKEN_PARA && p.espiar().Tipo == lexer.TOKEN_IDENTIFICADOR {
		return true
	}

	// "A Entidade" = declaração de entidade (sempre nível superior)
	if (tok.Tipo == lexer.TOKEN_ARTIGO_DEFINIDO || tok.Tipo == lexer.TOKEN_ARTIGO_INDEFINIDO) &&
		p.espiar().Tipo == lexer.TOKEN_ENTIDADE {
		return true
	}

	// Só aplicar heurística de artigos/Exibir se estamos em bloco profundo (>= 2).
	// nivelProfundidade == 1 é o corpo de uma função/if/loop, que pode
	// conter livremente declarações como Se, Exibir, Repita sem sair.
	if p.nivelProfundidade < 2 {
		return false
	}

	// Artigo definido/indefinido = declaração de variável/constante de nível superior
	if tok.Tipo == lexer.TOKEN_ARTIGO_DEFINIDO || tok.Tipo == lexer.TOKEN_ARTIGO_INDEFINIDO {
		return true
	}

	// "Exibir" no nível superior
	if tok.Tipo == lexer.TOKEN_EXIBIR {
		return true
	}

	// Identificador seguido de "com" = chamada de função no nível superior
	if tok.Tipo == lexer.TOKEN_IDENTIFICADOR && p.espiar().Tipo == lexer.TOKEN_COM {
		return true
	}

	// V2: Novos tokens de nível superior
	if tok.Tipo == lexer.TOKEN_SIMULTANEAMENTE ||
		tok.Tipo == lexer.TOKEN_TENTE ||
		tok.Tipo == lexer.TOKEN_SINALIZE {
		return true
	}

	return false
}

// -----------------------------------------------
// Análise de Expressões
// -----------------------------------------------

// analisarExpressao analisa uma expressão com operadores.
func (p *Parser) analisarExpressao() ast.Expressao {
	return p.analisarExpressaoAditiva()
}

// analisarExpressaoAditiva analisa + e - (e "e" como concatenação).
func (p *Parser) analisarExpressaoAditiva() ast.Expressao {
	esquerda := p.analisarExpressaoMultiplicativa()

	for p.tokenAtual().Tipo == lexer.TOKEN_MAIS || p.tokenAtual().Tipo == lexer.TOKEN_MENOS {
		op := p.avancar()
		direita := p.analisarExpressaoMultiplicativa()
		esquerda = &ast.ExpressaoBinaria{
			Token:    op,
			Esquerda: esquerda,
			Operador: op.Valor,
			Direita:  direita,
		}
	}

	return esquerda
}

// analisarExpressaoMultiplicativa analisa * e /.
func (p *Parser) analisarExpressaoMultiplicativa() ast.Expressao {
	esquerda := p.analisarExpressaoPrimaria()

	for p.tokenAtual().Tipo == lexer.TOKEN_MULTIPLICAR || p.tokenAtual().Tipo == lexer.TOKEN_DIVIDIR {
		op := p.avancar()
		direita := p.analisarExpressaoPrimaria()
		esquerda = &ast.ExpressaoBinaria{
			Token:    op,
			Esquerda: esquerda,
			Operador: op.Valor,
			Direita:  direita,
		}
	}

	return esquerda
}

// pularConectivos consome preposições e artigos que servem apenas
// para dar fluidez gramatical à linguagem (ex: "para", "o", "no").
func (p *Parser) pularConectivos() {
	for !p.fimDoArquivo() {
		tok := p.tokenAtual().Tipo
		if tok == lexer.TOKEN_ARTIGO_DEFINIDO ||
			tok == lexer.TOKEN_ARTIGO_INDEFINIDO ||
			tok == lexer.TOKEN_DE ||
			tok == lexer.TOKEN_COM ||
			tok == lexer.TOKEN_PARA ||
			tok == lexer.TOKEN_EM ||
			tok == lexer.TOKEN_AO ||
			tok == lexer.TOKEN_NO ||
			tok == lexer.TOKEN_PELO ||
			tok == lexer.TOKEN_POR {
			p.avancar()
		} else {
			break
		}
	}
}

// analisarExpressaoPrimaria analisa literais, identificadores e parênteses.
func (p *Parser) analisarExpressaoPrimaria() ast.Expressao {
	// Pular artigos ou preposições soltas antes da expressão primária
	p.pularConectivos()
	
	tok := p.tokenAtual()

	switch tok.Tipo {
	case lexer.TOKEN_NUMERO:
		p.avancar()
		return &ast.ExpressaoLiteralNumero{Token: tok, Valor: tok.Valor}

	case lexer.TOKEN_TEXTO:
		p.avancar()
		return &ast.ExpressaoLiteralTexto{Token: tok, Valor: tok.Valor}

	case lexer.TOKEN_VERDADEIRO:
		p.avancar()
		return &ast.ExpressaoLiteralLogico{Token: tok, Valor: true}

	case lexer.TOKEN_FALSO:
		p.avancar()
		return &ast.ExpressaoLiteralLogico{Token: tok, Valor: false}

	case lexer.TOKEN_NULO:
		p.avancar()
		return &ast.ExpressaoNulo{Token: tok}

	case lexer.TOKEN_COLCHETE_ABRE:
		// V2: Lista literal: [elem1, elem2, ...]
		return p.analisarExpressaoLista()

	case lexer.TOKEN_NOVO:
		return p.analisarExpressaoInstanciacao()

	case lexer.TOKEN_IDENTIFICADOR, lexer.TOKEN_TIPO:
		p.avancar()
		// V2: Acesso por campo: "campo de objeto"
		if p.tokenAtual().Tipo == lexer.TOKEN_DE {
			p.avancar() // consome "de"/"do"/"da"
			p.pularConectivos()
			if p.tokenAtual().Tipo == lexer.TOKEN_IDENTIFICADOR || p.tokenAtual().Tipo == lexer.TOKEN_TIPO {
				objTok := p.avancar()
				expr := &ast.ExpressaoAcessoCampo{
					Token:  tok,
					Campo:  tok.Valor,
					Objeto: &ast.ExpressaoIdentificador{Token: objTok, Nome: objTok.Valor},
				}
				// Check for index access after field access
				if p.tokenAtual().Tipo == lexer.TOKEN_COLCHETE_ABRE {
					return p.analisarAcessoIndice(expr)
				}
				// Verify explicit method call: "Objeto de Metodo com (args)"
				if p.tokenAtual().Tipo == lexer.TOKEN_COM {
					p.avancar()
					args := p.analisarArgumentos()
					return &ast.ExpressaoChamadaFuncao{
						Token:      tok,
						Nome:       expr.Campo,
						Objeto:     expr.Objeto,
						Argumentos: args,
					}
				}
				return expr
			}
		}
		// Verificar chamada de função explícita: "NomeFuncao com (args)"
		// Chamadas naturais puras (sem 'com') são processadas apenas como
		// declarações de nível superior (analisarDeclaracaoIdentificador)
		// para evitar ambiguidades com identificadores que são argumentos em sequência.
		if p.tokenAtual().Tipo == lexer.TOKEN_COM {
			p.avancar()
			args := p.analisarArgumentos()
			return &ast.ExpressaoChamadaFuncao{
				Token:      tok,
				Nome:       tok.Valor,
				Argumentos: args,
			}
		}
		// V2: Acesso por índice: lista[0]
		ident := &ast.ExpressaoIdentificador{Token: tok, Nome: tok.Valor}
		if p.tokenAtual().Tipo == lexer.TOKEN_COLCHETE_ABRE {
			return p.analisarAcessoIndice(ident)
		}
		return ident

	case lexer.TOKEN_PARENTESE_ABRE:
		p.avancar()
		expr := p.analisarExpressao()
		p.esperarTipo(lexer.TOKEN_PARENTESE_FECHA)
		return &ast.ExpressaoAgrupada{Token: tok, Expressao: expr}

	case lexer.TOKEN_NAO, lexer.TOKEN_MENOS:
		operador := tok.Valor
		p.avancar()
		operando := p.analisarExpressaoPrimaria()
		return &ast.ExpressaoUnaria{Token: tok, Operador: operador, Operando: operando}

	case lexer.TOKEN_RECEBER:
		return p.analisarExpressaoReceber()

	case lexer.TOKEN_CANAL:
		return p.analisarExpressaoCriarCanal()

	default:
		p.erro(fmt.Sprintf("expressão inesperada: %s (%q)", tok.Tipo.NomeLegivel(), tok.Valor))
		p.avancar()
		return nil
	}
}

// espiarComConectivos retorna o tipo do próximo token relevante, ignorando conectivos.
func (p *Parser) espiarComConectivos() lexer.TokenType {
	pos := p.posicao
	for pos < len(p.tokens) {
		tok := p.tokens[pos].Tipo
		if tok == lexer.TOKEN_ARTIGO_DEFINIDO ||
			tok == lexer.TOKEN_ARTIGO_INDEFINIDO ||
			tok == lexer.TOKEN_DE ||
			tok == lexer.TOKEN_COM ||
			tok == lexer.TOKEN_PARA ||
			tok == lexer.TOKEN_EM ||
			tok == lexer.TOKEN_AO ||
			tok == lexer.TOKEN_NO ||
			tok == lexer.TOKEN_PELO ||
			tok == lexer.TOKEN_POR {
			pos++
		} else {
			return tok
		}
	}
	return lexer.TOKEN_FIM
}

// ehInicioExpressao determina se um token pode iniciar uma expressão primária.
// Usado para saber se um identificador é uma chamada de função natural (Exibir o valor).
// NÃO incluímos COLCHETE_ABRE aqui para evitar que lista[0] seja interpretado
// como uma chamada de função lista([0]).
func (p *Parser) ehInicioExpressao(tipo lexer.TokenType) bool {
	return tipo == lexer.TOKEN_NUMERO ||
		tipo == lexer.TOKEN_TEXTO ||
		tipo == lexer.TOKEN_IDENTIFICADOR ||
		tipo == lexer.TOKEN_TIPO ||
		tipo == lexer.TOKEN_NOVO ||
		tipo == lexer.TOKEN_VERDADEIRO ||
		tipo == lexer.TOKEN_FALSO ||
		tipo == lexer.TOKEN_NULO ||
		tipo == lexer.TOKEN_PARENTESE_ABRE ||
		tipo == lexer.TOKEN_RECEBER ||
		tipo == lexer.TOKEN_CANAL ||
		tipo == lexer.TOKEN_NAO
}

// analisarArgumentos analisa: (expr1, expr2, ...) ou Mover o bloco para o destino
func (p *Parser) analisarArgumentos() []ast.Expressao {
	var args []ast.Expressao

	if p.tokenAtual().Tipo != lexer.TOKEN_PARENTESE_ABRE {
		// V2: Chamada natural sem parênteses, ex: Mover o bloco para o destino
		for !p.fimDoArquivo() && p.tokenAtual().Tipo != lexer.TOKEN_PONTO {
			fmt.Printf("DEBUG INICIO LOOP: token atual = %s (%q)\n", p.tokenAtual().Tipo.NomeLegivel(), p.tokenAtual().Valor)
			p.pularConectivos()
			fmt.Printf("DEBUG APOS PULAR: token atual = %s (%q)\n", p.tokenAtual().Tipo.NomeLegivel(), p.tokenAtual().Valor)
			
			// Se chegamos a um ponto ou fim, fim do comando
			if p.tokenAtual().Tipo == lexer.TOKEN_PONTO || p.fimDoArquivo() {
				break
			}

			// Lida com vírgulas explícitas
			if p.tokenAtual().Tipo == lexer.TOKEN_VIRGULA {
				p.avancar()
				continue
			}
			
			// Devemos ter o início de uma expressão aqui
			if !p.ehInicioAuxiliarExpressao(p.tokenAtual().Tipo) {
				fmt.Printf("DEBUG: analisador quebrou no token: %s (valor: %q)\n", p.tokenAtual().Tipo.NomeLegivel(), p.tokenAtual().Valor)
				break // encontrou algo inesperado, provavelmente fim da chamada
			}
			
			arg := p.analisarExpressao()
			if arg != nil {
				args = append(args, arg)
			}
		}
		return args
	}

	p.avancar() // consome "("

	for p.tokenAtual().Tipo != lexer.TOKEN_PARENTESE_FECHA && !p.fimDoArquivo() {
		if p.tokenAtual().Tipo == lexer.TOKEN_VIRGULA {
			p.avancar()
			continue
		}
		arg := p.analisarExpressao()
		if arg != nil {
			args = append(args, arg)
		}
	}

	p.esperarTipo(lexer.TOKEN_PARENTESE_FECHA)
	return args
}

// ehInicioAuxiliarExpressao aceita colchetes (para listas) como inicio
func (p *Parser) ehInicioAuxiliarExpressao(tipo lexer.TokenType) bool {
	return p.ehInicioExpressao(tipo) || tipo == lexer.TOKEN_COLCHETE_ABRE
}

// -----------------------------------------------
// Canais e Concorrência (V2)
// -----------------------------------------------

// analisarDeclaracaoEnviar analisa: "Enviar 10 para via."
func (p *Parser) analisarDeclaracaoEnviar() ast.Declaracao {
	tok := p.avancar() // consome Enviar

	// Opcionalmente consumir conectivos antes do valor (ex: "Enviar o valor...")
	// Mas como 'analisarExpressao' lida com os artigos, não é estritamente necessário,
	// porém é bom pular antes de tentar processar.
	p.pularConectivos()
	
	valor := p.analisarExpressao()

	// A preposição "para" é conectivo e será pulada se vier na frente de "via" 
	// por 'pularConectivos()', mas também podemos pular agora para chegar no identificador do canal.
	p.pularConectivos()

	if p.tokenAtual().Tipo != lexer.TOKEN_IDENTIFICADOR {
		p.erroEsperado(lexer.TOKEN_IDENTIFICADOR)
		p.avancar()
		return nil
	}

	canal := p.avancar().Valor
	p.consumirPonto()

	return &ast.DeclaracaoEnviar{
		Token: tok,
		Valor: valor,
		Canal: canal,
	}
}

// analisarExpressaoReceber analisa: "Receber de via"
func (p *Parser) analisarExpressaoReceber() ast.Expressao {
	tok := p.avancar() // consome Receber

	// pular "de", "do", "da"
	p.pularConectivos()

	if p.tokenAtual().Tipo != lexer.TOKEN_IDENTIFICADOR {
		p.erro("esperava identificador de canal após Receber")
		p.avancar()
		return nil
	}

	canal := p.avancar().Valor

	return &ast.ExpressaoReceber{
		Token: tok,
		Canal: canal,
	}
}

// analisarExpressaoCriarCanal analisa: "Canal de Inteiros"
func (p *Parser) analisarExpressaoCriarCanal() ast.Expressao {
	tok := p.avancar() // consome Canal

	// pular "de"
	p.pularConectivos()

	if p.tokenAtual().Tipo != lexer.TOKEN_IDENTIFICADOR && p.tokenAtual().Tipo != lexer.TOKEN_TIPO {
		p.erro("esperava tipo do canal após Canal")
		p.avancar()
		return nil
	}

	tipoItem := p.avancar().Valor

	return &ast.ExpressaoCriarCanal{
		Token:    tok,
		TipoItem: tipoItem,
	}
}

// analisarDeclaracaoIncluir analisa: "Incluir Matematica."
func (p *Parser) analisarDeclaracaoIncluir() ast.Declaracao {
	tok := p.avancar() // consome Incluir

	if p.tokenAtual().Tipo != lexer.TOKEN_IDENTIFICADOR && p.tokenAtual().Tipo != lexer.TOKEN_TIPO {
		p.erroEsperado(lexer.TOKEN_IDENTIFICADOR)
		p.avancar()
		return nil
	}

	pacote := p.avancar().Valor
	p.consumirPonto()

	return &ast.DeclaracaoIncluir{
		Token:  tok,
		Pacote: pacote,
	}
}

// analisarExpressaoInstanciacao analisa: "novo Pessoa contendo (nome: \"Ada\")"
func (p *Parser) analisarExpressaoInstanciacao() ast.Expressao {
	tok := p.avancar() // consome "novo"

	if p.tokenAtual().Tipo != lexer.TOKEN_IDENTIFICADOR {
		p.erroEsperado(lexer.TOKEN_IDENTIFICADOR)
		p.avancar()
		return nil
	}
	
	tipo := p.avancar().Valor

	// "contendo" é opcional
	if p.tokenAtual().Tipo == lexer.TOKEN_CONTENDO {
		p.avancar()
	}

	args := p.analisarArgumentos()

	return &ast.ExpressaoInstanciacao{
		Token:      tok,
		Tipo:       tipo,
		Argumentos: args,
	}
}
