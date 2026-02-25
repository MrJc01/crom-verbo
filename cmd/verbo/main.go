// Verbo CLI — Ferramenta de linha de comando para a linguagem Verbo.
// Subcomandos:
//   verbo compilar <arquivo.vrb>   — Transpila para Go e compila
//   verbo executar <arquivo.vrb>   — Transpila, compila e executa
//   verbo verificar <arquivo.vrb>  — Apenas verifica a sintaxe
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/juanxto/crom-verbo/pkg/lexer"
	"github.com/juanxto/crom-verbo/pkg/parser"
	"github.com/juanxto/crom-verbo/pkg/transpiler"
)

const versao = "0.1.0"

func main() {
	if len(os.Args) < 2 {
		exibirAjuda()
		os.Exit(1)
	}

	comando := os.Args[1]

	switch comando {
	case "compilar":
		if len(os.Args) < 3 {
			erroFatal("uso: verbo compilar <arquivo.vrb>")
		}
		executarCompilar(os.Args[2])

	case "executar":
		if len(os.Args) < 3 {
			erroFatal("uso: verbo executar <arquivo.vrb>")
		}
		executarExecutar(os.Args[2])

	case "verificar":
		if len(os.Args) < 3 {
			erroFatal("uso: verbo verificar <arquivo.vrb>")
		}
		executarVerificar(os.Args[2])

	case "versao", "versão", "--version", "-v":
		fmt.Printf("Verbo v%s\n", versao)

	case "ajuda", "help", "--help", "-h":
		exibirAjuda()

	default:
		fmt.Fprintf(os.Stderr, "❌ Comando desconhecido: '%s'\n\n", comando)
		exibirAjuda()
		os.Exit(1)
	}
}

// executarCompilar transpila o arquivo .vrb para Go e compila em binário.
func executarCompilar(caminhoArquivo string) {
	codigoGo := transpilar(caminhoArquivo)

	// Escrever o código Go
	nomeBase := strings.TrimSuffix(filepath.Base(caminhoArquivo), ".vrb")
	arquivoGo := nomeBase + "_verbo.go"

	if err := os.WriteFile(arquivoGo, []byte(codigoGo), 0644); err != nil {
		erroFatal(fmt.Sprintf("erro ao escrever arquivo Go: %v", err))
	}

	fmt.Printf("📝 Código Go gerado: %s\n", arquivoGo)

	// Compilar com go build
	binario := nomeBase
	cmd := exec.Command("go", "build", "-o", binario, arquivoGo)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		erroFatal(fmt.Sprintf("erro ao compilar código Go: %v", err))
	}

	fmt.Printf("✅ Binário compilado: ./%s\n", binario)
}

// executarExecutar transpila, compila e executa o programa.
func executarExecutar(caminhoArquivo string) {
	codigoGo := transpilar(caminhoArquivo)

	// Escrever arquivo temporário
	arquivoTemp := filepath.Join(os.TempDir(), "verbo_exec.go")
	if err := os.WriteFile(arquivoTemp, []byte(codigoGo), 0644); err != nil {
		erroFatal(fmt.Sprintf("erro ao escrever arquivo temporário: %v", err))
	}
	defer os.Remove(arquivoTemp)

	fmt.Println("🚀 Executando programa Verbo...")
	fmt.Println(strings.Repeat("─", 40))

	// Executar com go run
	cmd := exec.Command("go", "run", arquivoTemp)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		erroFatal(fmt.Sprintf("erro ao executar programa: %v", err))
	}

	fmt.Println(strings.Repeat("─", 40))
	fmt.Println("✅ Programa finalizado com sucesso.")
}

// executarVerificar apenas verifica a sintaxe sem compilar.
func executarVerificar(caminhoArquivo string) {
	codigo := lerArquivo(caminhoArquivo)

	// Análise Léxica
	lex := lexer.Novo(codigo)
	tokens, err := lex.Tokenizar()
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Erros léxicos em '%s':\n%v\n", caminhoArquivo, err)
		os.Exit(1)
	}

	fmt.Printf("📊 Tokens encontrados: %d\n", len(tokens))

	// Análise Sintática
	p := parser.Novo(tokens)
	programa, err := p.Analisar()
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Erros sintáticos em '%s':\n%v\n", caminhoArquivo, err)
		os.Exit(1)
	}

	fmt.Printf("🌳 Declarações na AST: %d\n", len(programa.Declaracoes))
	fmt.Printf("✅ Arquivo '%s' está sintaticamente correto!\n", caminhoArquivo)
}

// transpilar lê o arquivo .vrb e retorna o código Go equivalente.
func transpilar(caminhoArquivo string) string {
	codigo := lerArquivo(caminhoArquivo)

	// Análise Léxica
	lex := lexer.Novo(codigo)
	tokens, err := lex.Tokenizar()
	if err != nil {
		erroFatal(fmt.Sprintf("erros léxicos:\n%v", err))
	}

	// Análise Sintática
	p := parser.Novo(tokens)
	programa, err := p.Analisar()
	if err != nil {
		erroFatal(fmt.Sprintf("erros sintáticos:\n%v", err))
	}

	// Transpilação
	trans := transpiler.Novo()
	codigoGo, err := trans.Transpilar(programa)
	if err != nil {
		erroFatal(fmt.Sprintf("erro de transpilação:\n%v", err))
	}

	return codigoGo
}

// lerArquivo lê o conteúdo de um arquivo .vrb.
func lerArquivo(caminho string) string {
	if !strings.HasSuffix(caminho, ".vrb") {
		erroFatal(fmt.Sprintf("arquivo deve ter extensão .vrb (recebido: '%s')", caminho))
	}

	conteudo, err := os.ReadFile(caminho)
	if err != nil {
		erroFatal(fmt.Sprintf("erro ao ler arquivo '%s': %v", caminho, err))
	}

	return string(conteudo)
}

// erroFatal exibe uma mensagem de erro e encerra.
func erroFatal(msg string) {
	fmt.Fprintf(os.Stderr, "❌ %s\n", msg)
	os.Exit(1)
}

// exibirAjuda mostra a ajuda do CLI.
func exibirAjuda() {
	fmt.Printf(`🇧🇷 Verbo v%s — Linguagem de Programação em Português

Uso:
  verbo <comando> [argumentos]

Comandos:
  compilar <arquivo.vrb>   Transpila para Go e compila em binário
  executar <arquivo.vrb>   Transpila, compila e executa o programa
  verificar <arquivo.vrb>  Apenas verifica a sintaxe

  versão                   Exibe a versão do Verbo
  ajuda                    Exibe esta ajuda

Exemplos:
  verbo executar ola_mundo.vrb
  verbo compilar calculadora.vrb
  verbo verificar meu_programa.vrb

Para mais informações: https://github.com/MrJc01/crom-verbo
`, versao)
}
