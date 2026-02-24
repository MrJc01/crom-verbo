// Package transpiler converte a AST da linguagem Verbo em código Go compilável.
// Utiliza o padrão Visitor para percorrer cada nó da árvore e gerar
// o código Go equivalente.
package transpiler

import (
	"fmt"
	"strings"

	"github.com/juanxto/crom-verbo/pkg/ast"
)

// Transpiler converte uma AST Verbo em código-fonte Go.
type Transpiler struct {
	saida       strings.Builder
	indentacao  int
	funcoes     map[string]bool // rastreia funções declaradas
}

// Novo cria um novo Transpiler.
func Novo() *Transpiler {
	return &Transpiler{
		funcoes: make(map[string]bool),
	}
}

// Transpilar converte um programa Verbo completo em código Go.
func (t *Transpiler) Transpilar(programa *ast.Programa) (string, error) {
	t.saida.Reset()

	// Cabeçalho Go
	t.escreverLinha("package main")
	t.escreverLinha("")
	t.escreverLinha("import \"fmt\"")
	t.escreverLinha("")

	// Separar declarações de funções das declarações de nível superior
	var funcs []ast.Declaracao
	var principal []ast.Declaracao

	for _, decl := range programa.Declaracoes {
		if _, ok := decl.(*ast.DeclaracaoFuncao); ok {
			funcs = append(funcs, decl)
		} else {
			principal = append(principal, decl)
		}
	}

	// Gerar funções antes do main
	for _, decl := range funcs {
		t.transpilarDeclaracao(decl)
		t.escreverLinha("")
	}

	// Gerar função main
	t.escreverLinha("func main() {")
	t.indentacao++
	for _, decl := range principal {
		t.transpilarDeclaracao(decl)
	}
	t.indentacao--
	t.escreverLinha("}")

	return t.saida.String(), nil
}

// -----------------------------------------------
// Geração de Declarações
// -----------------------------------------------

func (t *Transpiler) transpilarDeclaracao(decl ast.Declaracao) {
	switch d := decl.(type) {
	case *ast.DeclaracaoVariavel:
		t.transpilarDeclaracaoVariavel(d)
	case *ast.DeclaracaoFuncao:
		t.transpilarDeclaracaoFuncao(d)
	case *ast.DeclaracaoExibir:
		t.transpilarDeclaracaoExibir(d)
	case *ast.DeclaracaoSe:
		t.transpilarDeclaracaoSe(d)
	case *ast.DeclaracaoRepita:
		t.transpilarDeclaracaoRepita(d)
	case *ast.DeclaracaoEnquanto:
		t.transpilarDeclaracaoEnquanto(d)
	case *ast.DeclaracaoRetorne:
		t.transpilarDeclaracaoRetorne(d)
	case *ast.DeclaracaoAtribuicao:
		t.transpilarDeclaracaoAtribuicao(d)
	case *ast.DeclaracaoExpressao:
		t.escreverIndentado(t.transpilarExpressao(d.Expressao))
		t.saida.WriteString("\n")
	}
}

func (t *Transpiler) transpilarDeclaracaoVariavel(d *ast.DeclaracaoVariavel) {
	valor := t.transpilarExpressao(d.Valor)

	if d.Mutavel {
		// Um/Uma → var (mutável)
		t.escreverIndentado(fmt.Sprintf("%s := %s", d.Nome, valor))
	} else {
		// O/A → const (imutável) — usamos var para simplificar
		// pois const em Go só aceita valores constantes compiláveis
		t.escreverIndentado(fmt.Sprintf("%s := %s", d.Nome, valor))
	}
	t.saida.WriteString("\n")
}

func (t *Transpiler) transpilarDeclaracaoFuncao(d *ast.DeclaracaoFuncao) {
	t.funcoes[d.Nome] = true

	// Parâmetros
	var params []string
	for _, p := range d.Parametros {
		tipoGo := t.converterTipo(p.Tipo)
		params = append(params, fmt.Sprintf("%s %s", p.Nome, tipoGo))
	}

	t.escreverIndentado(fmt.Sprintf("func %s(%s) interface{} {", d.Nome, strings.Join(params, ", ")))
	t.saida.WriteString("\n")

	t.indentacao++
	if d.Corpo != nil {
		for _, decl := range d.Corpo.Declaracoes {
			t.transpilarDeclaracao(decl)
		}
	}
	// Adicionar return nil se o corpo não termina com Retorne
	if d.Corpo == nil || len(d.Corpo.Declaracoes) == 0 || !t.ultimoEhRetorne(d.Corpo) {
		t.escreverIndentado("return nil")
		t.saida.WriteString("\n")
	}
	t.indentacao--
	t.escreverIndentado("}")
	t.saida.WriteString("\n")
}

func (t *Transpiler) transpilarDeclaracaoExibir(d *ast.DeclaracaoExibir) {
	valor := t.transpilarExpressao(d.Valor)
	t.escreverIndentado(fmt.Sprintf("fmt.Println(%s)", valor))
	t.saida.WriteString("\n")
}

func (t *Transpiler) transpilarDeclaracaoSe(d *ast.DeclaracaoSe) {
	condicao := t.transpilarExpressao(d.Condicao)
	t.escreverIndentado(fmt.Sprintf("if %s {", condicao))
	t.saida.WriteString("\n")

	t.indentacao++
	if d.Consequencia != nil {
		for _, decl := range d.Consequencia.Declaracoes {
			t.transpilarDeclaracao(decl)
		}
	}
	t.indentacao--

	if d.Alternativa != nil && len(d.Alternativa.Declaracoes) > 0 {
		t.escreverIndentado("} else {")
		t.saida.WriteString("\n")
		t.indentacao++
		for _, decl := range d.Alternativa.Declaracoes {
			t.transpilarDeclaracao(decl)
		}
		t.indentacao--
	}

	t.escreverIndentado("}")
	t.saida.WriteString("\n")
}

func (t *Transpiler) transpilarDeclaracaoRepita(d *ast.DeclaracaoRepita) {
	if d.ForEach {
		iteravel := t.transpilarExpressao(d.Iteravel)
		t.escreverIndentado(fmt.Sprintf("for _, %s := range %s {", d.Variavel, iteravel))
	} else {
		contagem := t.transpilarExpressao(d.Contagem)
		t.escreverIndentado(fmt.Sprintf("for i := 0; i < %s; i++ {", contagem))
	}
	t.saida.WriteString("\n")

	t.indentacao++
	if d.Corpo != nil {
		for _, decl := range d.Corpo.Declaracoes {
			t.transpilarDeclaracao(decl)
		}
	}
	t.indentacao--
	t.escreverIndentado("}")
	t.saida.WriteString("\n")
}

func (t *Transpiler) transpilarDeclaracaoEnquanto(d *ast.DeclaracaoEnquanto) {
	condicao := t.transpilarExpressao(d.Condicao)
	t.escreverIndentado(fmt.Sprintf("for %s {", condicao))
	t.saida.WriteString("\n")

	t.indentacao++
	if d.Corpo != nil {
		for _, decl := range d.Corpo.Declaracoes {
			t.transpilarDeclaracao(decl)
		}
	}
	t.indentacao--
	t.escreverIndentado("}")
	t.saida.WriteString("\n")
}

func (t *Transpiler) transpilarDeclaracaoRetorne(d *ast.DeclaracaoRetorne) {
	if d.Valor != nil {
		valor := t.transpilarExpressao(d.Valor)
		t.escreverIndentado(fmt.Sprintf("return %s", valor))
	} else {
		t.escreverIndentado("return nil")
	}
	t.saida.WriteString("\n")
}

func (t *Transpiler) transpilarDeclaracaoAtribuicao(d *ast.DeclaracaoAtribuicao) {
	valor := t.transpilarExpressao(d.Valor)
	t.escreverIndentado(fmt.Sprintf("%s = %s", d.Nome, valor))
	t.saida.WriteString("\n")
}

// -----------------------------------------------
// Geração de Expressões
// -----------------------------------------------

func (t *Transpiler) transpilarExpressao(expr ast.Expressao) string {
	if expr == nil {
		return "nil"
	}

	switch e := expr.(type) {
	case *ast.ExpressaoLiteralNumero:
		return e.Valor

	case *ast.ExpressaoLiteralTexto:
		return fmt.Sprintf("%q", e.Valor)

	case *ast.ExpressaoLiteralLogico:
		if e.Valor {
			return "true"
		}
		return "false"

	case *ast.ExpressaoNulo:
		return "nil"

	case *ast.ExpressaoIdentificador:
		return e.Nome

	case *ast.ExpressaoBinaria:
		esq := t.transpilarExpressao(e.Esquerda)
		dir := t.transpilarExpressao(e.Direita)
		op := t.converterOperador(e.Operador)
		return fmt.Sprintf("(%s %s %s)", esq, op, dir)

	case *ast.ExpressaoUnaria:
		operando := t.transpilarExpressao(e.Operando)
		switch e.Operador {
		case "não":
			return fmt.Sprintf("!%s", operando)
		case "-":
			return fmt.Sprintf("-%s", operando)
		default:
			return operando
		}

	case *ast.ExpressaoChamadaFuncao:
		var args []string
		for _, arg := range e.Argumentos {
			args = append(args, t.transpilarExpressao(arg))
		}
		return fmt.Sprintf("%s(%s)", e.Nome, strings.Join(args, ", "))

	case *ast.ExpressaoAgrupada:
		return fmt.Sprintf("(%s)", t.transpilarExpressao(e.Expressao))

	default:
		return "/* expressão não suportada */"
	}
}

// -----------------------------------------------
// Helpers
// -----------------------------------------------

func (t *Transpiler) converterOperador(op string) string {
	switch op {
	case "+", "e":
		return "+"
	case "-", "menos":
		return "-"
	case "*":
		return "*"
	case "/":
		return "/"
	case "menor que":
		return "<"
	case "maior que":
		return ">"
	case "igual":
		return "=="
	default:
		return op
	}
}

func (t *Transpiler) converterTipo(tipo string) string {
	switch tipo {
	case "Texto":
		return "string"
	case "Inteiro":
		return "int"
	case "Decimal":
		return "float64"
	case "Logico", "Lógico":
		return "bool"
	case "":
		return "interface{}"
	default:
		return "interface{}"
	}
}

func (t *Transpiler) escreverLinha(texto string) {
	t.saida.WriteString(texto)
	t.saida.WriteString("\n")
}

func (t *Transpiler) escreverIndentado(texto string) {
	for i := 0; i < t.indentacao; i++ {
		t.saida.WriteString("\t")
	}
	t.saida.WriteString(texto)
}

func (t *Transpiler) ultimoEhRetorne(bloco *ast.Bloco) bool {
	if len(bloco.Declaracoes) == 0 {
		return false
	}
	_, ok := bloco.Declaracoes[len(bloco.Declaracoes)-1].(*ast.DeclaracaoRetorne)
	return ok
}
