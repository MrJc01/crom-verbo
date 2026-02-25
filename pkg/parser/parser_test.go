package parser

import (
	"testing"

	"github.com/juanxto/crom-verbo/pkg/ast"
	"github.com/juanxto/crom-verbo/pkg/lexer"
)

func analisarPrograma(t *testing.T, entrada string) *ast.Programa {
	t.Helper()
	lex := lexer.Novo(entrada)
	tokens, err := lex.Tokenizar()
	if err != nil {
		t.Fatalf("erro léxico: %v", err)
	}

	parser := Novo(tokens)
	programa, err := parser.Analisar()
	if err != nil {
		t.Fatalf("erro sintático: %v", err)
	}

	return programa
}

func TestAnalisarDeclaracaoConstante(t *testing.T) {
	programa := analisarPrograma(t, `A mensagem é "Olá, Mundo!".`)

	if len(programa.Declaracoes) != 1 {
		t.Fatalf("esperava 1 declaração, obteve %d", len(programa.Declaracoes))
	}

	decl, ok := programa.Declaracoes[0].(*ast.DeclaracaoVariavel)
	if !ok {
		t.Fatalf("esperava DeclaracaoVariavel, obteve %T", programa.Declaracoes[0])
	}

	if decl.Nome != "mensagem" {
		t.Errorf("esperava nome 'mensagem', obteve '%s'", decl.Nome)
	}

	if decl.Mutavel {
		t.Error("esperava constante (imutável), mas está marcada como mutável")
	}

	if decl.Verbo != "é" {
		t.Errorf("esperava verbo 'é', obteve '%s'", decl.Verbo)
	}

	lit, ok := decl.Valor.(*ast.ExpressaoLiteralTexto)
	if !ok {
		t.Fatalf("esperava ExpressaoLiteralTexto, obteve %T", decl.Valor)
	}

	if lit.Valor != "Olá, Mundo!" {
		t.Errorf("esperava 'Olá, Mundo!', obteve '%s'", lit.Valor)
	}
}

func TestAnalisarDeclaracaoVariavel(t *testing.T) {
	programa := analisarPrograma(t, `Um contador está 0.`)

	if len(programa.Declaracoes) != 1 {
		t.Fatalf("esperava 1 declaração, obteve %d", len(programa.Declaracoes))
	}

	decl, ok := programa.Declaracoes[0].(*ast.DeclaracaoVariavel)
	if !ok {
		t.Fatalf("esperava DeclaracaoVariavel, obteve %T", programa.Declaracoes[0])
	}

	if decl.Nome != "contador" {
		t.Errorf("esperava nome 'contador', obteve '%s'", decl.Nome)
	}

	if !decl.Mutavel {
		t.Error("esperava variável mutável, mas está marcada como constante")
	}

	if decl.Verbo != "está" {
		t.Errorf("esperava verbo 'está', obteve '%s'", decl.Verbo)
	}

	lit, ok := decl.Valor.(*ast.ExpressaoLiteralNumero)
	if !ok {
		t.Fatalf("esperava ExpressaoLiteralNumero, obteve %T", decl.Valor)
	}

	if lit.Valor != "0" {
		t.Errorf("esperava '0', obteve '%s'", lit.Valor)
	}
}

func TestAnalisarDeclaracaoFuncao(t *testing.T) {
	entrada := `Para Calcular usando (valor: Inteiro):
    Retorne valor.`

	programa := analisarPrograma(t, entrada)

	if len(programa.Declaracoes) != 1 {
		t.Fatalf("esperava 1 declaração, obteve %d", len(programa.Declaracoes))
	}

	decl, ok := programa.Declaracoes[0].(*ast.DeclaracaoFuncao)
	if !ok {
		t.Fatalf("esperava DeclaracaoFuncao, obteve %T", programa.Declaracoes[0])
	}

	if decl.Nome != "Calcular" {
		t.Errorf("esperava nome 'Calcular', obteve '%s'", decl.Nome)
	}

	if len(decl.Parametros) != 1 {
		t.Fatalf("esperava 1 parâmetro, obteve %d", len(decl.Parametros))
	}

	if decl.Parametros[0].Nome != "valor" {
		t.Errorf("esperava parâmetro 'valor', obteve '%s'", decl.Parametros[0].Nome)
	}

	if decl.Parametros[0].Tipo != "Inteiro" {
		t.Errorf("esperava tipo 'Inteiro', obteve '%s'", decl.Parametros[0].Tipo)
	}

	if decl.Corpo == nil || len(decl.Corpo.Declaracoes) != 1 {
		t.Fatalf("esperava corpo com 1 declaração")
	}
}

func TestAnalisarExibir(t *testing.T) {
	programa := analisarPrograma(t, `Exibir com ("Olá, Mundo!").`)

	if len(programa.Declaracoes) != 1 {
		t.Fatalf("esperava 1 declaração, obteve %d", len(programa.Declaracoes))
	}

	decl, ok := programa.Declaracoes[0].(*ast.DeclaracaoExibir)
	if !ok {
		t.Fatalf("esperava DeclaracaoExibir, obteve %T", programa.Declaracoes[0])
	}

	lit, ok := decl.Valor.(*ast.ExpressaoLiteralTexto)
	if !ok {
		t.Fatalf("esperava ExpressaoLiteralTexto, obteve %T", decl.Valor)
	}

	if lit.Valor != "Olá, Mundo!" {
		t.Errorf("esperava 'Olá, Mundo!', obteve '%s'", lit.Valor)
	}
}

func TestAnalisarRepitaVezes(t *testing.T) {
	entrada := `Repita 5 vezes:
    Exibir com ("x").`

	programa := analisarPrograma(t, entrada)

	if len(programa.Declaracoes) != 1 {
		t.Fatalf("esperava 1 declaração, obteve %d", len(programa.Declaracoes))
	}

	decl, ok := programa.Declaracoes[0].(*ast.DeclaracaoRepita)
	if !ok {
		t.Fatalf("esperava DeclaracaoRepita, obteve %T", programa.Declaracoes[0])
	}

	if decl.ForEach {
		t.Error("esperava 'N vezes', mas está marcado como 'para cada'")
	}

	cont, ok := decl.Contagem.(*ast.ExpressaoLiteralNumero)
	if !ok {
		t.Fatalf("esperava ExpressaoLiteralNumero para contagem, obteve %T", decl.Contagem)
	}

	if cont.Valor != "5" {
		t.Errorf("esperava contagem '5', obteve '%s'", cont.Valor)
	}
}

func TestAnalisarExpressaoBinaria(t *testing.T) {
	programa := analisarPrograma(t, `Um resultado é 10 + 5.`)

	if len(programa.Declaracoes) != 1 {
		t.Fatalf("esperava 1 declaração, obteve %d", len(programa.Declaracoes))
	}

	decl, ok := programa.Declaracoes[0].(*ast.DeclaracaoVariavel)
	if !ok {
		t.Fatalf("esperava DeclaracaoVariavel, obteve %T", programa.Declaracoes[0])
	}

	bin, ok := decl.Valor.(*ast.ExpressaoBinaria)
	if !ok {
		t.Fatalf("esperava ExpressaoBinaria, obteve %T", decl.Valor)
	}

	if bin.Operador != "+" {
		t.Errorf("esperava operador '+', obteve '%s'", bin.Operador)
	}
}

func TestAnalisarMultiplasDeclaracoes(t *testing.T) {
	entrada := `A nome é "Verbo".
Um versao está 1.
Exibir com (nome).`

	programa := analisarPrograma(t, entrada)

	if len(programa.Declaracoes) != 3 {
		t.Fatalf("esperava 3 declarações, obteve %d", len(programa.Declaracoes))
	}

	// Primeira: constante
	if _, ok := programa.Declaracoes[0].(*ast.DeclaracaoVariavel); !ok {
		t.Errorf("declaração 0: esperava DeclaracaoVariavel, obteve %T", programa.Declaracoes[0])
	}

	// Segunda: variável
	if _, ok := programa.Declaracoes[1].(*ast.DeclaracaoVariavel); !ok {
		t.Errorf("declaração 1: esperava DeclaracaoVariavel, obteve %T", programa.Declaracoes[1])
	}

	// Terceira: exibir
	if _, ok := programa.Declaracoes[2].(*ast.DeclaracaoExibir); !ok {
		t.Errorf("declaração 2: esperava DeclaracaoExibir, obteve %T", programa.Declaracoes[2])
	}
}

// -----------------------------------------------
// Testes V2
// -----------------------------------------------

func TestAnalisarEntidade(t *testing.T) {
	entrada := `A Entidade Produto contendo (nome: Texto, preco: Inteiro).`
	programa := analisarPrograma(t, entrada)

	if len(programa.Declaracoes) != 1 {
		t.Fatalf("esperava 1 declaração, obteve %d", len(programa.Declaracoes))
	}

	decl, ok := programa.Declaracoes[0].(*ast.DeclaracaoEntidade)
	if !ok {
		t.Fatalf("esperava DeclaracaoEntidade, obteve %T", programa.Declaracoes[0])
	}

	if decl.Nome != "Produto" {
		t.Errorf("esperava nome 'Produto', obteve '%s'", decl.Nome)
	}

	if len(decl.Campos) != 2 {
		t.Fatalf("esperava 2 campos, obteve %d", len(decl.Campos))
	}

	if decl.Campos[0].Nome != "nome" || decl.Campos[0].Tipo != "Texto" {
		t.Errorf("campo 0: esperava 'nome: Texto', obteve '%s: %s'", decl.Campos[0].Nome, decl.Campos[0].Tipo)
	}

	if decl.Campos[1].Nome != "preco" || decl.Campos[1].Tipo != "Inteiro" {
		t.Errorf("campo 1: esperava 'preco: Inteiro', obteve '%s: %s'", decl.Campos[1].Nome, decl.Campos[1].Tipo)
	}
}

func TestAnalisarLista(t *testing.T) {
	entrada := `Uma frutas é ["maçã", "banana", "uva"].`
	programa := analisarPrograma(t, entrada)

	if len(programa.Declaracoes) != 1 {
		t.Fatalf("esperava 1 declaração, obteve %d", len(programa.Declaracoes))
	}

	decl, ok := programa.Declaracoes[0].(*ast.DeclaracaoVariavel)
	if !ok {
		t.Fatalf("esperava DeclaracaoVariavel, obteve %T", programa.Declaracoes[0])
	}

	lista, ok := decl.Valor.(*ast.ExpressaoLista)
	if !ok {
		t.Fatalf("esperava ExpressaoLista, obteve %T", decl.Valor)
	}

	if len(lista.Elementos) != 3 {
		t.Errorf("esperava 3 elementos, obteve %d", len(lista.Elementos))
	}
}

func TestAnalisarSimultaneamente(t *testing.T) {
	entrada := `Simultaneamente:
    Exibir com ("Tarefa 1").`

	programa := analisarPrograma(t, entrada)

	if len(programa.Declaracoes) != 1 {
		t.Fatalf("esperava 1 declaração, obteve %d", len(programa.Declaracoes))
	}

	decl, ok := programa.Declaracoes[0].(*ast.DeclaracaoSimultaneamente)
	if !ok {
		t.Fatalf("esperava DeclaracaoSimultaneamente, obteve %T", programa.Declaracoes[0])
	}

	if decl.Corpo == nil || len(decl.Corpo.Declaracoes) == 0 {
		t.Fatal("esperava corpo com pelo menos 1 declaração")
	}
}

func TestAnalisarTenteCapture(t *testing.T) {
	entrada := `Tente:
    Sinalize com ("erro!").
Capture erro:
    Exibir com ("capturado").`

	programa := analisarPrograma(t, entrada)

	if len(programa.Declaracoes) != 1 {
		t.Fatalf("esperava 1 declaração, obteve %d", len(programa.Declaracoes))
	}

	decl, ok := programa.Declaracoes[0].(*ast.DeclaracaoTente)
	if !ok {
		t.Fatalf("esperava DeclaracaoTente, obteve %T", programa.Declaracoes[0])
	}

	if decl.Tentativa == nil || len(decl.Tentativa.Declaracoes) == 0 {
		t.Fatal("esperava bloco tentativa com declarações")
	}

	if decl.VariavelErro != "erro" {
		t.Errorf("esperava variável de erro 'erro', obteve '%s'", decl.VariavelErro)
	}

	if decl.Captura == nil || len(decl.Captura.Declaracoes) == 0 {
		t.Fatal("esperava bloco capture com declarações")
	}
}

func TestAnalisarAcessoCampo(t *testing.T) {
	entrada := `Exibir com (nome de usuario).`
	programa := analisarPrograma(t, entrada)

	if len(programa.Declaracoes) != 1 {
		t.Fatalf("esperava 1 declaração, obteve %d", len(programa.Declaracoes))
	}

	exibir, ok := programa.Declaracoes[0].(*ast.DeclaracaoExibir)
	if !ok {
		t.Fatalf("esperava DeclaracaoExibir, obteve %T", programa.Declaracoes[0])
	}

	acesso, ok := exibir.Valor.(*ast.ExpressaoAcessoCampo)
	if !ok {
		t.Fatalf("esperava ExpressaoAcessoCampo, obteve %T", exibir.Valor)
	}

	if acesso.Campo != "nome" {
		t.Errorf("esperava campo 'nome', obteve '%s'", acesso.Campo)
	}

	ident, ok := acesso.Objeto.(*ast.ExpressaoIdentificador)
	if !ok {
		t.Fatalf("esperava ExpressaoIdentificador como objeto, obteve %T", acesso.Objeto)
	}

	if ident.Nome != "usuario" {
		t.Errorf("esperava objeto 'usuario', obteve '%s'", ident.Nome)
	}
}

func TestAnalisarAcessoIndice(t *testing.T) {
	entrada := `Exibir com (lista[0]).`
	programa := analisarPrograma(t, entrada)

	if len(programa.Declaracoes) != 1 {
		t.Fatalf("esperava 1 declaração, obteve %d", len(programa.Declaracoes))
	}

	exibir, ok := programa.Declaracoes[0].(*ast.DeclaracaoExibir)
	if !ok {
		t.Fatalf("esperava DeclaracaoExibir, obteve %T", programa.Declaracoes[0])
	}

	acesso, ok := exibir.Valor.(*ast.ExpressaoAcessoIndice)
	if !ok {
		t.Fatalf("esperava ExpressaoAcessoIndice, obteve %T", exibir.Valor)
	}

	ident, ok := acesso.Objeto.(*ast.ExpressaoIdentificador)
	if !ok {
		t.Fatalf("esperava ExpressaoIdentificador como objeto, obteve %T", acesso.Objeto)
	}

	if ident.Nome != "lista" {
		t.Errorf("esperava objeto 'lista', obteve '%s'", ident.Nome)
	}

	idx, ok := acesso.Indice.(*ast.ExpressaoLiteralNumero)
	if !ok {
		t.Fatalf("esperava ExpressaoLiteralNumero como índice, obteve %T", acesso.Indice)
	}

	if idx.Valor != "0" {
		t.Errorf("esperava índice '0', obteve '%s'", idx.Valor)
	}
}

func TestAnalisarChamadaNatural(t *testing.T) {
	entrada := `Exibir o valor.`
	programa := analisarPrograma(t, entrada)

	if len(programa.Declaracoes) != 1 {
		t.Fatalf("esperava 1 declaração, obteve %d", len(programa.Declaracoes))
	}

	exibir, ok := programa.Declaracoes[0].(*ast.DeclaracaoExibir)
	if !ok {
		t.Fatalf("esperava DeclaracaoExibir, obteve %T", programa.Declaracoes[0])
	}

	ident, ok := exibir.Valor.(*ast.ExpressaoIdentificador)
	if !ok {
		t.Fatalf("esperava ExpressaoIdentificador, obteve %T", exibir.Valor)
	}

	if ident.Nome != "valor" {
		t.Errorf("esperava 'valor', obteve '%s'", ident.Nome)
	}
}

func TestAnalisarMultiplosArgumentosNaturais(t *testing.T) {
	entrada := `Substituir na frase "olá" por "oi".`
	programa := analisarPrograma(t, entrada)

	if len(programa.Declaracoes) != 1 {
		t.Fatalf("esperava 1 declaração, obteve %d", len(programa.Declaracoes))
	}

	exprDecl, ok := programa.Declaracoes[0].(*ast.DeclaracaoExpressao)
	if !ok {
		t.Fatalf("esperava DeclaracaoExpressao, obteve %T", programa.Declaracoes[0])
	}

	chamada, ok := exprDecl.Expressao.(*ast.ExpressaoChamadaFuncao)
	if !ok {
		t.Fatalf("esperava ExpressaoChamadaFuncao, obteve %T", exprDecl.Expressao)
	}

	if chamada.Nome != "Substituir" {
		t.Errorf("esperava nome 'Substituir', obteve '%s'", chamada.Nome)
	}

	if len(chamada.Argumentos) != 3 {
		t.Fatalf("esperava 3 argumentos, obteve %d", len(chamada.Argumentos))
	}

	arg1, ok := chamada.Argumentos[0].(*ast.ExpressaoIdentificador)
	if !ok || arg1.Nome != "frase" {
		t.Errorf("arg[0] esperava Identificador 'frase', obteve %T", chamada.Argumentos[0])
	}

	arg2, ok := chamada.Argumentos[1].(*ast.ExpressaoLiteralTexto)
	if !ok || arg2.Valor != "olá" {
		t.Errorf("arg[1] esperava Texto 'olá', obteve %T", chamada.Argumentos[1])
	}

	arg3, ok := chamada.Argumentos[2].(*ast.ExpressaoLiteralTexto)
	if !ok || arg3.Valor != "oi" {
		t.Errorf("arg[2] esperava Texto 'oi', obteve %T", chamada.Argumentos[2])
	}
}
func TestAnalisarCanais(t *testing.T) {
	codigo := `
	Uma via é um Canal de Inteiros.
	Enviar 10 para via.
	O valor é Receber de via.
	`

	programa := analisarPrograma(t, codigo)

	if len(programa.Declaracoes) != 3 {
		t.Fatalf("O programa não possui declarações suficientes. Esperava 3. Obteve %d", len(programa.Declaracoes))
	}

	// 1. O canal
	decl1, ok := programa.Declaracoes[0].(*ast.DeclaracaoVariavel)
	if !ok {
		t.Fatalf("esperava ast.DeclaracaoVariavel, obteve %T", programa.Declaracoes[0])
	}

	exprCanal, ok := decl1.Valor.(*ast.ExpressaoCriarCanal)
	if !ok {
		t.Fatalf("esperava ast.ExpressaoCriarCanal, obteve %T", decl1.Valor)
	}

	if exprCanal.TipoItem != "Inteiros" {
		t.Errorf("esperava tipo 'Inteiros', obteve %q", exprCanal.TipoItem)
	}

	// 2. Enviar
	decl2, ok := programa.Declaracoes[1].(*ast.DeclaracaoEnviar)
	if !ok {
		t.Fatalf("esperava ast.DeclaracaoEnviar, obteve %T", programa.Declaracoes[1])
	}
	if decl2.Canal != "via" {
		t.Errorf("esperava canal 'via', obteve %q", decl2.Canal)
	}

	// 3. Receber
	decl3, ok := programa.Declaracoes[2].(*ast.DeclaracaoVariavel)
	if !ok {
		t.Fatalf("esperava ast.DeclaracaoVariavel, obteve %T", programa.Declaracoes[2])
	}

	exprRecv, ok := decl3.Valor.(*ast.ExpressaoReceber)
	if !ok {
		t.Fatalf("esperava ast.ExpressaoReceber, obteve %T", decl3.Valor)
	}
	if exprRecv.Canal != "via" {
		t.Errorf("esperava canal 'via', obteve %q", exprRecv.Canal)
	}
}
