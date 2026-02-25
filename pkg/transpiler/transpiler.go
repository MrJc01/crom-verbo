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
	funcoes     map[string]bool           // rastreia funções declaradas
	imutaveis   map[string]bool           // V2: rastreia variáveis imutáveis (decl. com 'é')
	entidades   map[string][]ast.CampoEntidade // V2: rastreia entidades declaradas
	usaSync     bool                      // V2: precisa importar "sync"
	imports     map[string]bool           // V2: pacotes a importar (ex: Matematica)
	erros       []string                  // V2: erros de compilação
}

// Novo cria um novo Transpiler.
func Novo() *Transpiler {
	return &Transpiler{
		funcoes:   make(map[string]bool),
		imutaveis: make(map[string]bool),
		entidades: make(map[string][]ast.CampoEntidade),
		imports:   make(map[string]bool),
	}
}

// Transpilar converte um programa Verbo completo em código Go.
func (t *Transpiler) Transpilar(programa *ast.Programa) (string, error) {
	t.saida.Reset()

	// Função auxiliar para procurar "Simultaneamente" recursivamente
	var checarSync func(decl ast.Declaracao)
	checarSync = func(decl ast.Declaracao) {
		if _, ok := decl.(*ast.DeclaracaoSimultaneamente); ok {
			t.usaSync = true
		}
		if f, ok := decl.(*ast.DeclaracaoFuncao); ok && f.Corpo != nil {
			for _, d := range f.Corpo.Declaracoes {
				checarSync(d)
			}
		}
		if tentativa, ok := decl.(*ast.DeclaracaoTente); ok {
			if tentativa.Tentativa != nil {
				for _, d := range tentativa.Tentativa.Declaracoes { checarSync(d) }
			}
			if tentativa.Captura != nil {
				for _, d := range tentativa.Captura.Declaracoes { checarSync(d) }
			}
		}
	}

	// Primeira passada: detectar se precisa de "sync", pacotes importados e registrar entidades
	for _, decl := range programa.Declaracoes {
		checarSync(decl)
		if ent, ok := decl.(*ast.DeclaracaoEntidade); ok {
			t.entidades[ent.Nome] = ent.Campos
		}
		if inc, ok := decl.(*ast.DeclaracaoIncluir); ok {
			t.imports[inc.Pacote] = true
		}
	}

	// Cabeçalho Go
	t.escreverLinha("package main")
	t.escreverLinha("")
	
	if len(t.imports) > 0 || t.usaSync {
		t.escreverLinha("import (")
		t.escreverLinha("\t\"fmt\"")
		if t.usaSync {
			t.escreverLinha("\t\"sync\"")
		}
		for imp := range t.imports {
			switch imp {
			case "Matematica":
				t.escreverLinha("\t\"github.com/juanxto/crom-verbo/pkg/stdlib/matematica\"")
			case "Texto":
				t.escreverLinha("\t\"github.com/juanxto/crom-verbo/pkg/stdlib/texto\"")
			case "Arquivo":
				t.escreverLinha("\t\"github.com/juanxto/crom-verbo/pkg/stdlib/arquivo\"")
			case "Html":
				t.escreverLinha("\t\"github.com/juanxto/crom-verbo/pkg/stdlib/html\"")
			default:
				// Fallback simplificado se não achar na BibVerbo: import direto
				t.escreverLinha(fmt.Sprintf("\t%q", strings.ToLower(imp)))
			}
		}
		t.escreverLinha(")")
	} else {
		t.escreverLinha("import \"fmt\"")
	}
	t.escreverLinha("")

	// Separar declarações
	var funcs []ast.Declaracao
	var entidades []ast.Declaracao
	var principal []ast.Declaracao

	for _, decl := range programa.Declaracoes {
		switch decl.(type) {
		case *ast.DeclaracaoFuncao:
			funcs = append(funcs, decl)
		case *ast.DeclaracaoEntidade:
			entidades = append(entidades, decl)
		default:
			principal = append(principal, decl)
		}
	}

	// V2: Gerar structs (entidades) no topo
	for _, decl := range entidades {
		t.transpilarDeclaracao(decl)
		t.escreverLinha("")
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

	// V2: erros de compilação (imutabilidade)
	if len(t.erros) > 0 {
		return "", fmt.Errorf("erros de compilação Verbo:\n%s", strings.Join(t.erros, "\n"))
	}

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
	// V2
	case *ast.DeclaracaoEntidade:
		t.transpilarDeclaracaoEntidade(d)
	case *ast.DeclaracaoSimultaneamente:
		t.transpilarDeclaracaoSimultaneamente(d)
	case *ast.DeclaracaoTente:
		t.transpilarDeclaracaoTente(d)
	case *ast.DeclaracaoSinalize:
		t.transpilarDeclaracaoSinalize(d)
	case *ast.DeclaracaoEnviar:
		t.transpilarDeclaracaoEnviar(d)
	case *ast.DeclaracaoIncluir:
		// Não gera código aqui, pois os imports foram feitos no topo
	}
}

func (t *Transpiler) transpilarDeclaracaoVariavel(d *ast.DeclaracaoVariavel) {
	valor := t.transpilarExpressao(d.Valor)

	if d.Verbo == "é" {
		// V2: Imutabilidade — rastrear e gerar como constante quando possível
		t.imutaveis[d.Nome] = true
		t.escreverIndentado(fmt.Sprintf("%s := %s", d.Nome, valor))
	} else {
		// "está" = mutável
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
		// Remover conversão explícita pois slices declarados já tem o tipo correto
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
	// V2: Verificação de imutabilidade
	if t.imutaveis[d.Nome] {
		t.erros = append(t.erros, fmt.Sprintf(
			"erro semântico: '%s' foi declarado com 'é' (imutável) e não pode ser reatribuído. Use 'está' para variáveis mutáveis.", d.Nome))
		return
	}
	valor := t.transpilarExpressao(d.Valor)
	t.escreverIndentado(fmt.Sprintf("%s = %s", d.Nome, valor))
	t.saida.WriteString("\n")
}

// -----------------------------------------------
// V2: Geração de Declarações Avançadas
// -----------------------------------------------

func (t *Transpiler) transpilarDeclaracaoEntidade(d *ast.DeclaracaoEntidade) {
	t.escreverIndentado(fmt.Sprintf("type %s struct {", d.Nome))
	t.saida.WriteString("\n")
	t.indentacao++
	for _, campo := range d.Campos {
		// Capitalizar nome do campo para export em Go
		nomeCap := strings.ToUpper(campo.Nome[:1]) + campo.Nome[1:]
		tipoGo := t.converterTipo(campo.Tipo)
		t.escreverIndentado(fmt.Sprintf("%s %s", nomeCap, tipoGo))
		t.saida.WriteString("\n")
	}
	t.indentacao--
	t.escreverIndentado("}")
	t.saida.WriteString("\n")
}

func (t *Transpiler) transpilarDeclaracaoSimultaneamente(d *ast.DeclaracaoSimultaneamente) {
	if d.Corpo == nil || len(d.Corpo.Declaracoes) == 0 {
		return
	}

	n := len(d.Corpo.Declaracoes)
	
	// Usar um bloco anônimo para escopar a variável wg
	t.escreverIndentado("{")
	t.saida.WriteString("\n")
	t.indentacao++
	
	t.escreverIndentado("var wg sync.WaitGroup")
	t.saida.WriteString("\n")
	t.escreverIndentado(fmt.Sprintf("wg.Add(%d)", n))
	t.saida.WriteString("\n")

	for _, decl := range d.Corpo.Declaracoes {
		t.escreverIndentado("go func() {")
		t.saida.WriteString("\n")
		t.indentacao++
		t.escreverIndentado("defer wg.Done()")
		t.saida.WriteString("\n")
		t.transpilarDeclaracao(decl)
		t.indentacao--
		t.escreverIndentado("}()")
		t.saida.WriteString("\n")
	}

	t.escreverIndentado("wg.Wait()")
	t.saida.WriteString("\n")
	
	t.indentacao--
	t.escreverIndentado("}")
	t.saida.WriteString("\n")
}

func (t *Transpiler) transpilarDeclaracaoTente(d *ast.DeclaracaoTente) {
	// defer/recover pattern
	t.escreverIndentado("func() {")
	t.saida.WriteString("\n")
	t.indentacao++

	// defer com recover
	varErro := d.VariavelErro
	if varErro == "" {
		varErro = "r"
	}
	t.escreverIndentado(fmt.Sprintf("defer func() {"))
	t.saida.WriteString("\n")
	t.indentacao++
	t.escreverIndentado(fmt.Sprintf("if %s := recover(); %s != nil {", varErro, varErro))
	t.saida.WriteString("\n")

	if d.Captura != nil && len(d.Captura.Declaracoes) > 0 {
		t.indentacao++
		for _, decl := range d.Captura.Declaracoes {
			t.transpilarDeclaracao(decl)
		}
		t.indentacao--
	}

	t.escreverIndentado("}")
	t.saida.WriteString("\n")
	t.indentacao--
	t.escreverIndentado("}()")
	t.saida.WriteString("\n")

	// corpo da tentativa
	if d.Tentativa != nil {
		for _, decl := range d.Tentativa.Declaracoes {
			t.transpilarDeclaracao(decl)
		}
	}

	t.indentacao--
	t.escreverIndentado("}()")
	t.saida.WriteString("\n")
}

func (t *Transpiler) transpilarDeclaracaoSinalize(d *ast.DeclaracaoSinalize) {
	valor := t.transpilarExpressao(d.Valor)
	t.escreverIndentado(fmt.Sprintf("panic(%s)", valor))
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

		// V2: Se tiver Objeto, gerar obj.Metodo(args)
		if e.Objeto != nil {
			objStr := t.transpilarExpressao(e.Objeto)
			
			// Se o objeto for um tipo primitivo (package importado), garantir que o package fique em minusculas,
			// mas a função seja em maiúsculas ("Matematica de Absoluto" -> "matematica.Absoluto")
			if _, isImported := t.imports[objStr]; isImported {
				objStr = strings.ToLower(objStr)
			}
			
			metodoCap := strings.ToUpper(e.Nome[:1]) + e.Nome[1:]
			return fmt.Sprintf("%s.%s(%s)", objStr, metodoCap, strings.Join(args, ", "))
		}
		// V2: Se o nome for uma entidade conhecida, gerar struct literal
		if campos, ok := t.entidades[e.Nome]; ok {
			var pares []string
			for i, campo := range campos {
				if i < len(args) {
					nomeCap := strings.ToUpper(campo.Nome[:1]) + campo.Nome[1:]
					pares = append(pares, fmt.Sprintf("%s: %s", nomeCap, args[i]))
				}
			}
			return fmt.Sprintf("%s{%s}", e.Nome, strings.Join(pares, ", "))
		}
		return fmt.Sprintf("%s(%s)", e.Nome, strings.Join(args, ", "))

	case *ast.ExpressaoAgrupada:
		return fmt.Sprintf("(%s)", t.transpilarExpressao(e.Expressao))

	// V2: Novas expressões
	case *ast.ExpressaoLista:
		var elems []string
		for _, elem := range e.Elementos {
			elems = append(elems, t.transpilarExpressao(elem))
		}
		return fmt.Sprintf("[]interface{}{%s}", strings.Join(elems, ", "))

	case *ast.ExpressaoAcessoIndice:
		obj := t.transpilarExpressao(e.Objeto)
		idx := t.transpilarExpressao(e.Indice)
		return fmt.Sprintf("%s[%s]", obj, idx)

	case *ast.ExpressaoAcessoCampo:
		obj := t.transpilarExpressao(e.Objeto)
		// Capitalizar campo para acesso Go (campos exported)
		nomeCap := strings.ToUpper(e.Campo[:1]) + e.Campo[1:]
		return fmt.Sprintf("%s.%s", obj, nomeCap)

	case *ast.ExpressaoInstanciacao:
		var args []string
		for _, arg := range e.Argumentos {
			args = append(args, t.transpilarExpressao(arg))
		}
		// Mapear argumentos para campos da entidade
		if campos, ok := t.entidades[e.Tipo]; ok {
			var pares []string
			for i, campo := range campos {
				if i < len(args) {
					nomeCap := strings.ToUpper(campo.Nome[:1]) + campo.Nome[1:]
					pares = append(pares, fmt.Sprintf("%s: %s", nomeCap, args[i]))
				}
			}
			return fmt.Sprintf("%s{%s}", e.Tipo, strings.Join(pares, ", "))
		}
		return fmt.Sprintf("%s{%s}", e.Tipo, strings.Join(args, ", "))

	case *ast.ExpressaoCriarCanal:
		return t.transpilarExpressaoCriarCanal(e)

	case *ast.ExpressaoReceber:
		return t.transpilarExpressaoReceber(e)

	default:
		return "/* expressão não suportada */"
	}
}

// -----------------------------------------------
// Helpers
// -----------------------------------------------

func (t *Transpiler) converterOperador(op string) string {
	switch op {
	case "+", "e", "soma", "mais":
		return "+"
	case "-", "menos", "subtrai":
		return "-"
	case "*", "multiplica":
		return "*"
	case "/", "divide":
		return "/"
	case "%", "módulo", "modulo", "porcentagem", "resto":
		return "%"
	case "menor que":
		return "<"
	case "maior que":
		return ">"
	case "igual", "idêntico", "identico":
		return "=="
	case "diferente":
		return "!="
	default:
		return op
	}
}

func (t *Transpiler) converterTipo(tipo string) string {
	switch tipo {
	case "Texto", "Textos":
		return "string"
	case "Inteiro", "Inteiros":
		return "int"
	case "Decimal", "Decimais":
		return "float64"
	case "Logico", "Lógico", "Logicos", "Lógicos":
		return "bool"
	case "Lista", "Listas":
		return "[]interface{}"
	case "Canal_Inteiros":
		return "chan int"
	case "":
		return "interface{}"
	default:
		// V2: Pode ser um nome de Entidade
		if _, ok := t.entidades[tipo]; ok {
			return tipo
		}
		// Vamos assumir interface{} de forma segura
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

// -----------------------------------------------
// V2: Canais e Concorrência Avançada
// -----------------------------------------------

func (t *Transpiler) transpilarExpressaoCriarCanal(e *ast.ExpressaoCriarCanal) string {
	tipoGo := t.converterTipo(e.TipoItem)
	if tipoGo == "interface{}" {
		tipoGo = "interface{}"
	}
	return fmt.Sprintf("make(chan %s, 100)", tipoGo)
}

func (t *Transpiler) transpilarExpressaoReceber(e *ast.ExpressaoReceber) string {
	return fmt.Sprintf("<-%s", e.Canal)
}

func (t *Transpiler) transpilarDeclaracaoEnviar(d *ast.DeclaracaoEnviar) {
	valor := t.transpilarExpressao(d.Valor)
	t.escreverIndentado(fmt.Sprintf("%s <- %s", d.Canal, valor))
	t.saida.WriteString("\n")
}
