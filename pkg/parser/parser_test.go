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
