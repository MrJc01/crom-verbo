package transpiler

import (
	"strings"
	"testing"

	"github.com/juanxto/crom-verbo/pkg/ast"
	"github.com/juanxto/crom-verbo/pkg/lexer"
	"github.com/juanxto/crom-verbo/pkg/parser"
)

func transpilarCodigo(t *testing.T, entrada string) string {
	t.Helper()

	lex := lexer.Novo(entrada)
	tokens, err := lex.Tokenizar()
	if err != nil {
		t.Fatalf("erro léxico: %v", err)
	}

	p := parser.Novo(tokens)
	programa, err := p.Analisar()
	if err != nil {
		t.Fatalf("erro sintático: %v", err)
	}

	trans := Novo()
	saida, err := trans.Transpilar(programa)
	if err != nil {
		t.Fatalf("erro de transpilação: %v", err)
	}

	return saida
}

func TestTranspilarOlaMundo(t *testing.T) {
	entrada := `Exibir com ("Olá, Mundo!").`
	saida := transpilarCodigo(t, entrada)

	if !strings.Contains(saida, `fmt.Println("Olá, Mundo!")`) {
		t.Errorf("código Go não contém fmt.Println esperado.\nSaída:\n%s", saida)
	}

	if !strings.Contains(saida, "package main") {
		t.Error("código Go não contém 'package main'")
	}

	if !strings.Contains(saida, `import "fmt"`) {
		t.Error("código Go não contém import de fmt")
	}

	if !strings.Contains(saida, "func main()") {
		t.Error("código Go não contém func main()")
	}
}

func TestTranspilarVariavelConstante(t *testing.T) {
	testes := []struct {
		nome         string
		entrada      string
		esperaContem string
	}{
		{
			nome:         "constante texto",
			entrada:      `A mensagem é "teste".`,
			esperaContem: `mensagem := "teste"`,
		},
		{
			nome:         "variável número",
			entrada:      `Um contador está 0.`,
			esperaContem: `contador := 0`,
		},
		{
			nome:         "número decimal",
			entrada:      `Um pi é 3.14.`,
			esperaContem: `pi := 3.14`,
		},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			saida := transpilarCodigo(t, tt.entrada)
			if !strings.Contains(saida, tt.esperaContem) {
				t.Errorf("esperava conter %q.\nSaída:\n%s", tt.esperaContem, saida)
			}
		})
	}
}

func TestTranspilarFuncao(t *testing.T) {
	entrada := `Para Saudar usando (nome: Texto):
    Exibir com (nome).`

	saida := transpilarCodigo(t, entrada)

	if !strings.Contains(saida, "func Saudar(nome string)") {
		t.Errorf("esperava declaração de função 'Saudar'.\nSaída:\n%s", saida)
	}

	if !strings.Contains(saida, "fmt.Println(nome)") {
		t.Errorf("esperava fmt.Println(nome) no corpo.\nSaída:\n%s", saida)
	}
}

func TestTranspilarRepita(t *testing.T) {
	entrada := `Repita 3 vezes:
    Exibir com ("x").`

	saida := transpilarCodigo(t, entrada)

	if !strings.Contains(saida, "for i := 0; i < 3; i++") {
		t.Errorf("esperava loop for com 3 iterações.\nSaída:\n%s", saida)
	}
}

func TestTranspilarExpressaoBinaria(t *testing.T) {
	entrada := `Um resultado é 10 + 5.`
	saida := transpilarCodigo(t, entrada)

	if !strings.Contains(saida, "resultado := (10 + 5)") {
		t.Errorf("esperava expressão binária.\nSaída:\n%s", saida)
	}
}

func TestTranspilarRetorne(t *testing.T) {
	entrada := `Para Dobrar usando (x: Inteiro):
    Retorne x.`

	saida := transpilarCodigo(t, entrada)

	if !strings.Contains(saida, "return x") {
		t.Errorf("esperava 'return x'.\nSaída:\n%s", saida)
	}
}

func TestTranspilarProgramaCompleto(t *testing.T) {
	entrada := `A mensagem é "Bem-vindo ao Verbo!".
Um contador está 0.

Exibir com (mensagem).

Repita 5 vezes:
    contador está contador + 1.
    Exibir com (contador).`

	saida := transpilarCodigo(t, entrada)

	verificacoes := []string{
		"package main",
		`import "fmt"`,
		"func main()",
		`mensagem := "Bem-vindo ao Verbo!"`,
		"contador := 0",
		"fmt.Println(mensagem)",
		"for i := 0; i < 5; i++",
		"contador = (contador + 1)",
		"fmt.Println(contador)",
	}

	for _, v := range verificacoes {
		if !strings.Contains(saida, v) {
			t.Errorf("esperava conter %q.\nSaída:\n%s", v, saida)
		}
	}
}

// TestConverterTipo testa a conversão de tipos Verbo → Go.
func TestConverterTipo(t *testing.T) {
	trans := Novo()

	testes := []struct {
		verbo string
		goTip string
	}{
		{"Texto", "string"},
		{"Inteiro", "int"},
		{"Decimal", "float64"},
		{"Logico", "bool"},
		{"Lógico", "bool"},
		{"", "interface{}"},
		{"Desconhecido", "interface{}"},
	}

	for _, tt := range testes {
		resultado := trans.converterTipo(tt.verbo)
		if resultado != tt.goTip {
			t.Errorf("converterTipo(%q): esperava %q, obteve %q", tt.verbo, tt.goTip, resultado)
		}
	}
}

// TestConverterOperador testa a conversão de operadores Verbo → Go.
func TestConverterOperador(t *testing.T) {
	trans := Novo()

	testes := []struct {
		verbo string
		goOp  string
	}{
		{"+", "+"},
		{"-", "-"},
		{"*", "*"},
		{"/", "/"},
		{"e", "+"},
		{"menos", "-"},
		{"menor que", "<"},
		{"maior que", ">"},
		{"igual", "=="},
	}

	for _, tt := range testes {
		resultado := trans.converterOperador(tt.verbo)
		if resultado != tt.goOp {
			t.Errorf("converterOperador(%q): esperava %q, obteve %q", tt.verbo, tt.goOp, resultado)
		}
	}
}

// -----------------------------------------------
// Testes V2
// -----------------------------------------------

func TestTranspilarEntidade(t *testing.T) {
	entrada := `A Entidade Produto contendo (nome: Texto, preco: Inteiro).`
	saida := transpilarCodigo(t, entrada)

	if !strings.Contains(saida, "type Produto struct {") {
		t.Errorf("esperava declaração de struct Produto.\nSaída:\n%s", saida)
	}

	if !strings.Contains(saida, "Nome string") {
		t.Errorf("esperava campo 'Nome string'.\nSaída:\n%s", saida)
	}

	if !strings.Contains(saida, "Preco int") {
		t.Errorf("esperava campo 'Preco int'.\nSaída:\n%s", saida)
	}
}

func TestTranspilarLista(t *testing.T) {
	entrada := `Uma frutas é ["maçã", "banana"].`
	saida := transpilarCodigo(t, entrada)

	if !strings.Contains(saida, `frutas := []interface{}{"maçã", "banana"}`) {
		t.Errorf("esperava literal de lista em Go.\nSaída:\n%s", saida)
	}
}

func TestTranspilarTenteCapture(t *testing.T) {
	entrada := `Tente:
    Sinalize com ("erro").
Capture erro:
    Exibir com (erro).`
	saida := transpilarCodigo(t, entrada)

	if !strings.Contains(saida, "defer func() {") {
		t.Errorf("esperava 'defer func() {'.\nSaída:\n%s", saida)
	}

	if !strings.Contains(saida, "if erro := recover(); erro != nil {") {
		t.Errorf("esperava 'if erro := recover(); erro != nil {'.\nSaída:\n%s", saida)
	}

	if !strings.Contains(saida, `panic("erro")`) {
		t.Errorf("esperava 'panic(\"erro\")'.\nSaída:\n%s", saida)
	}
}

func TestTranspilarSimultaneamente(t *testing.T) {
	entrada := `Simultaneamente:
    Exibir com ("1").
    Exibir com ("2").`
	saida := transpilarCodigo(t, entrada)

	if !strings.Contains(saida, "var wg sync.WaitGroup") {
		t.Errorf("esperava 'var wg sync.WaitGroup'.\nSaída:\n%s", saida)
	}

	if !strings.Contains(saida, "wg.Add(2)") {
		t.Errorf("esperava 'wg.Add(2)'.\nSaída:\n%s", saida)
	}

	if !strings.Contains(saida, "go func() {") {
		t.Errorf("esperava 'go func() {'.\nSaída:\n%s", saida)
	}

	if !strings.Contains(saida, "defer wg.Done()") {
		t.Errorf("esperava 'defer wg.Done()'.\nSaída:\n%s", saida)
	}

	if !strings.Contains(saida, "wg.Wait()") {
		t.Errorf("esperava 'wg.Wait()'.\nSaída:\n%s", saida)
	}
}

func TestTranspilarAcessoCampo(t *testing.T) {
	entrada := `A Entidade Produto contendo (nome: Texto).
Um prod é Produto com ("Mesa").
Exibir com (nome de prod).`
	saida := transpilarCodigo(t, entrada)

	if !strings.Contains(saida, "fmt.Println(prod.Nome)") {
		t.Errorf("esperava 'fmt.Println(prod.Nome)'.\nSaída:\n%s", saida)
	}
}

func TestTranspilarImutabilidade(t *testing.T) {
	entrada := `A x é 10.
x está 20.`
	
	lex := lexer.Novo(entrada)
	tokens, _ := lex.Tokenizar()
	p := parser.Novo(tokens)
	programa, _ := p.Analisar()

	trans := Novo()
	_, err := trans.Transpilar(programa)
	
	if err == nil {
		t.Fatal("esperava erro semântico de imutabilidade, mas compilou com sucesso")
	}

	if !strings.Contains(err.Error(), "imutável") {
		t.Errorf("esperava mensagem de erro sobre imutabilidade, obteve: %v", err)
	}
}

// Ensure imports are used
var _ = ast.Programa{}
func TestTranspilarCanais(t *testing.T) {
	codigo := `
	Uma via é um Canal de Inteiros.
	Enviar 10 para via.
	O valor é Receber de via.
	`

	codigoGerado := transpilarCodigo(t, codigo)

	esperados := []string{
		"via := make(chan int, 100)",
		"via <- 10",
		"valor := <-via",
	}

	for _, esp := range esperados {
		if !strings.Contains(codigoGerado, esp) {
			t.Errorf("Código gerado não contém a string esperada. Esperava:\n%s\n\nObtido:\n%s", esp, codigoGerado)
		}
	}
}

func TestTranspilarIncluir(t *testing.T) {
	codigo := `
	Incluir Matematica.
	Incluir ExemploCustom.
	O valor é 10.
	`
	codigoGerado := transpilarCodigo(t, codigo)

	esperados := []string{
		`"github.com/juanxto/crom-verbo/pkg/stdlib/matematica"`,
		`"exemplocustom"`,
		`valor := 10`,
	}

	for _, esp := range esperados {
		if !strings.Contains(codigoGerado, esp) {
			t.Errorf("Código gerado não contém a string esperada. Esperava:\n%s\n\nObtido:\n%s", esp, codigoGerado)
		}
	}
}
