// Verbo CLI — Ferramenta de linha de comando para a linguagem Verbo.
// Subcomandos:
//   verbo compilar <arquivo.vrb>   — Transpila para Go e compila
//   verbo executar <arquivo.vrb>   — Transpila, compila e executa
//   verbo verificar <arquivo.vrb>  — Apenas verifica a sintaxe
package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/juanxto/crom-verbo/pkg/lexer"
	"github.com/juanxto/crom-verbo/pkg/parser"
	"github.com/juanxto/crom-verbo/pkg/transpiler"
)

const versao = "0.1.0"

type installerAcoes struct {
	Copiar []struct {
		De   string `json:"de"`
		Para string `json:"para"`
	} `json:"copiar"`

	Dependencias []string `json:"dependencias"`

	CriarPasta []struct {
		Caminho string `json:"caminho"`
	} `json:"criar_pasta"`

	Remover []struct {
		Caminho string `json:"caminho"`
	} `json:"remover"`

	Patch []struct {
		Arquivo  string `json:"arquivo"`
		Procurar string `json:"procurar"`
		Trocar   string `json:"trocar"`
		Limite   int    `json:"limite"`
	} `json:"patch"`

	ExecutarComando []struct {
		Comando string   `json:"comando"`
		Args    []string `json:"args"`
		Cwd     string   `json:"cwd"`
	} `json:"executar_comando"`

	Allowlist struct {
		ExecutarComando []string `json:"executar_comando"`
	} `json:"allowlist"`
}

func executarInstallerComEstado(pastaPacote string, pilha []string, instalados map[string]bool) {
	installer := filepath.Join(pastaPacote, "installer.vrb")
	if _, err := os.Stat(installer); err != nil {
		return
	}
	for _, p := range pilha {
		if filepath.Clean(p) == filepath.Clean(pastaPacote) {
			erroFatal(fmt.Sprintf("installer: dependência cíclica detectada (pasta): %s", pastaPacote))
		}
	}

	fmt.Printf("🧩 Executando installer: %s\n", installer)
	cmd := exec.Command(os.Args[0], "executar", installer)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		erroFatal(fmt.Sprintf("installer falhou: %v", err))
	}
	// Se o installer emitir um JSON válido, aplicar ações.
	var a installerAcoes
	if err := json.Unmarshal(out.Bytes(), &a); err != nil {
		// não é erro: o installer pode só executar lógica e não emitir JSON.
		return
	}

	// Dependências primeiro
	for _, dep := range a.Dependencias {
		spec := strings.TrimSpace(dep)
		if spec == "" {
			continue
		}
		nomeDep, _, _, err := parsePacoteSpec(spec)
		if err != nil {
			erroFatal(fmt.Sprintf("installer dependencias: spec inválido: %s: %v", spec, err))
		}
		if instalados[strings.ToLower(nomeDep)] {
			continue
		}
		fmt.Printf("📦 Instalando dependência: %s\n", spec)
		instalarPacoteComEstado(spec, append(pilha, pastaPacote), instalados)
	}

	allowCmd := make(map[string]bool)
	for _, c := range a.Allowlist.ExecutarComando {
		allowCmd[c] = true
	}

	for _, d := range a.CriarPasta {
		dst, err := validarCaminhoRelativoSeguro("installer criar_pasta", d.Caminho)
		if err != nil {
			erroFatal(err.Error())
		}
		fmt.Printf("📁 Criando pasta %s\n", dst)
		if err := os.MkdirAll(dst, 0755); err != nil {
			erroFatal(fmt.Sprintf("installer criar_pasta: %v", err))
		}
	}

	for _, r := range a.Remover {
		target, err := validarCaminhoRelativoSeguro("installer remover", r.Caminho)
		if err != nil {
			erroFatal(err.Error())
		}
		// Segurança extra: impedir remover pasta de pacotes inteira ou lock/manifests
		if strings.EqualFold(target, "pacotes") || strings.EqualFold(target, "verbo.mod.json") || strings.EqualFold(target, "verbo.lock.json") {
			erroFatal(fmt.Sprintf("installer remover: operação bloqueada para '%s'", target))
		}
		fmt.Printf("🗑️  Removendo %s\n", target)
		if err := os.RemoveAll(target); err != nil {
			erroFatal(fmt.Sprintf("installer remover: %v", err))
		}
	}

	for _, c := range a.Copiar {
		srcRel, err := validarCaminhoRelativoSeguro("installer copiar.de", c.De)
		if err != nil {
			erroFatal(err.Error())
		}
		dstRel, err := validarCaminhoRelativoSeguro("installer copiar.para", c.Para)
		if err != nil {
			erroFatal(err.Error())
		}
		src := filepath.Join(pastaPacote, srcRel)
		dst := dstRel
		fmt.Printf("📄 Copiando %s -> %s\n", src, dst)
		info, err := os.Stat(src)
		if err != nil {
			erroFatal(fmt.Sprintf("installer copiar: arquivo não encontrado: %s", src))
		}
		if info.IsDir() {
			if err := copiarDiretorio(src, dst); err != nil {
				erroFatal(fmt.Sprintf("installer copiar dir: %v", err))
			}
			continue
		}
		if err := copiarArquivo(src, dst); err != nil {
			erroFatal(fmt.Sprintf("installer copiar arquivo: %v", err))
		}
	}

	for _, p := range a.Patch {
		arquivoRel, err := validarCaminhoRelativoSeguro("installer patch.arquivo", p.Arquivo)
		if err != nil {
			erroFatal(err.Error())
		}
		fmt.Printf("🧷 Patch em %s\n", arquivoRel)
		if err := aplicarPatchArquivo(arquivoRel, p.Procurar, p.Trocar, p.Limite); err != nil {
			erroFatal(fmt.Sprintf("installer patch: %s: %v", arquivoRel, err))
		}
	}

	for _, e := range a.ExecutarComando {
		if !allowCmd[e.Comando] {
			erroFatal(fmt.Sprintf("installer executar_comando: comando não permitido (allowlist): %s", e.Comando))
		}
		cwd := ""
		if strings.TrimSpace(e.Cwd) != "" {
			cwdRel, err := validarCaminhoRelativoSeguro("installer executar_comando.cwd", e.Cwd)
			if err != nil {
				erroFatal(err.Error())
			}
			cwd = cwdRel
		}
		fmt.Printf("▶️  Executando comando: %s %s\n", e.Comando, strings.Join(e.Args, " "))
		c := exec.Command(e.Comando, e.Args...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if cwd != "" {
			c.Dir = cwd
		}
		if err := c.Run(); err != nil {
			erroFatal(fmt.Sprintf("installer executar_comando: %v", err))
		}
	}
}

func validarCaminhoRelativoSeguro(label, caminho string) (string, error) {
	trim := strings.TrimSpace(caminho)
	if trim == "" {
		return "", fmt.Errorf("%s: caminho vazio", label)
	}
	clean := filepath.Clean(trim)
	if filepath.IsAbs(clean) {
		return "", fmt.Errorf("%s: caminho deve ser relativo: %s", label, caminho)
	}
	if clean == "." {
		return "", fmt.Errorf("%s: caminho inválido: %s", label, caminho)
	}
	// Bloquear traversal: qualquer segmento ".."
	for _, part := range strings.Split(clean, string(filepath.Separator)) {
		if part == ".." {
			return "", fmt.Errorf("%s: caminho não pode conter '..': %s", label, caminho)
		}
	}
	return clean, nil
}

func aplicarPatchArquivo(arquivo string, procurar string, trocar string, limite int) error {
	conteudo, err := os.ReadFile(arquivo)
	if err != nil {
		return err
	}
	texto := string(conteudo)
	if procurar == "" {
		return fmt.Errorf("patch: 'procurar' não pode ser vazio")
	}
	if limite <= 0 {
		limite = 1
	}
	cont := 0
	for {
		if cont >= limite {
			break
		}
		idx := strings.Index(texto, procurar)
		if idx < 0 {
			break
		}
		texto = texto[:idx] + trocar + texto[idx+len(procurar):]
		cont++
	}
	if cont == 0 {
		return fmt.Errorf("patch: padrão não encontrado")
	}
	return os.WriteFile(arquivo, []byte(texto), 0644)
}

func executarInstaller(pastaPacote string) {
	executarInstallerComEstado(pastaPacote, nil, map[string]bool{})
}

func instalarPacoteComEstado(spec string, pilha []string, instalados map[string]bool) {
	nome, _, _, err := parsePacoteSpec(spec)
	if err != nil {
		erroFatal(err.Error())
	}
	key := strings.ToLower(nome)
	if instalados[key] {
		return
	}
	instalados[key] = true

	// Se já estiver instalado no manifesto, não reinstalar automaticamente.
	m := carregarManifestoPacotes()
	for _, p := range m.Pacotes {
		if strings.EqualFold(p, nome) {
			fmt.Printf("✅ Dependência já instalada: %s\n", nome)
			return
		}
	}
	instalarPacote(spec)
}

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

	case "servir":
		executarServir(os.Args[2:])

	case "versao", "versão", "--version", "-v":
		fmt.Printf("Verbo v%s\n", versao)

	case "ajuda", "help", "--help", "-h":
		exibirAjuda()

	case "pacote":
		executarPacote(os.Args[2:])

	default:
		fmt.Fprintf(os.Stderr, "❌ Comando desconhecido: '%s'\n\n", comando)
		exibirAjuda()
		os.Exit(1)
	}
}

type manifestoPacotes struct {
	Pacotes []string `json:"pacotes"`
}

type lockPacote struct {
	Nome   string `json:"nome"`
	Fonte  string `json:"fonte"`
	Ref    string `json:"ref"`
	Pasta  string `json:"pasta"`
	Sha256 string `json:"sha256"`
}

type lockFile struct {
	Pacotes []lockPacote `json:"pacotes"`
}

func caminhoManifestoPacotes() string {
	return "verbo.mod.json"
}

func caminhoLockPacotes() string {
	return "verbo.lock.json"
}

func carregarManifestoPacotes() manifestoPacotes {
	caminho := caminhoManifestoPacotes()
	conteudo, err := os.ReadFile(caminho)
	if err != nil {
		return manifestoPacotes{Pacotes: []string{}}
	}
	var m manifestoPacotes
	if err := json.Unmarshal(conteudo, &m); err != nil {
		return manifestoPacotes{Pacotes: []string{}}
	}
	if m.Pacotes == nil {
		m.Pacotes = []string{}
	}
	return m
}

func salvarManifestoPacotes(m manifestoPacotes) {
	conteudo, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		erroFatal(fmt.Sprintf("erro ao serializar manifesto de pacotes: %v", err))
	}
	if err := os.WriteFile(caminhoManifestoPacotes(), conteudo, 0644); err != nil {
		erroFatal(fmt.Sprintf("erro ao escrever manifesto de pacotes: %v", err))
	}
}

func carregarLockPacotes() lockFile {
	caminho := caminhoLockPacotes()
	conteudo, err := os.ReadFile(caminho)
	if err != nil {
		return lockFile{Pacotes: []lockPacote{}}
	}
	var l lockFile
	if err := json.Unmarshal(conteudo, &l); err != nil {
		return lockFile{Pacotes: []lockPacote{}}
	}
	if l.Pacotes == nil {
		l.Pacotes = []lockPacote{}
	}
	return l
}

func salvarLockPacotes(l lockFile) {
	conteudo, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		erroFatal(fmt.Sprintf("erro ao serializar lockfile de pacotes: %v", err))
	}
	if err := os.WriteFile(caminhoLockPacotes(), conteudo, 0644); err != nil {
		erroFatal(fmt.Sprintf("erro ao escrever lockfile de pacotes: %v", err))
	}
}

func pacoteDir(nome string) string {
	return filepath.Join("pacotes", nome)
}

func garantirDiretorio(path string) {
	if err := os.MkdirAll(path, 0755); err != nil {
		erroFatal(fmt.Sprintf("erro ao criar diretório '%s': %v", path, err))
	}
}

func copiarArquivo(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

func copiarDiretorio(src, dst string) error {
	return filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0755)
		}
		return copiarArquivo(path, target)
	})
}

func unzipParaPasta(zipPath, dst string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		// Evitar ZipSlip
		cleanName := filepath.Clean(f.Name)
		if strings.Contains(cleanName, "..") {
			continue
		}
		outPath := filepath.Join(dst, cleanName)
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(outPath, 0755); err != nil {
				return err
			}
			continue
		}
		if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
			return err
		}
		in, err := f.Open()
		if err != nil {
			return err
		}
		out, err := os.Create(outPath)
		if err != nil {
			in.Close()
			return err
		}
		_, err = io.Copy(out, in)
		in.Close()
		out.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func baixarArquivo(url, dst string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

func parsePacoteSpec(spec string) (nome string, fonte string, ref string, err error) {
	// Formatos:
	// - gh:user/repo[@ref]
	// - path:C:\meu\pacote  (ou path:./pacote)
	if strings.HasPrefix(spec, "gh:") {
		fonte = "github"
		resto := strings.TrimPrefix(spec, "gh:")
		ref = "main"
		if strings.Contains(resto, "@") {
			parts := strings.SplitN(resto, "@", 2)
			resto = parts[0]
			ref = parts[1]
		}
		segs := strings.Split(resto, "/")
		if len(segs) != 2 {
			return "", "", "", fmt.Errorf("spec GitHub inválido (use gh:user/repo[@ref])")
		}
		nome = segs[1]
		return nome, spec, ref, nil
	}
	if strings.HasPrefix(spec, "path:") {
		fonte = "local"
		resto := strings.TrimPrefix(spec, "path:")
		resto = strings.TrimSpace(resto)
		if resto == "" {
			return "", "", "", fmt.Errorf("spec local inválido (use path:./pasta)")
		}
		nome = filepath.Base(resto)
		return nome, spec, "", nil
	}
	// compat: aceitar "nome" como pacote oficial/registrado (apenas registrando)
	return spec, spec, "", nil
}

func instalarPacote(spec string) {
	nome, fonteSpec, ref, err := parsePacoteSpec(spec)
	if err != nil {
		erroFatal(err.Error())
	}

	// registrar no manifesto
	m := carregarManifestoPacotes()
	for _, p := range m.Pacotes {
		if strings.EqualFold(p, nome) {
			fmt.Printf("✅ Pacote já instalado: %s\n", nome)
			return
		}
	}

	garantirDiretorio("pacotes")
	dst := pacoteDir(nome)
	_ = os.RemoveAll(dst)
	garantirDiretorio(dst)

	lock := carregarLockPacotes()

	if strings.HasPrefix(fonteSpec, "gh:") {
		repo := strings.TrimPrefix(fonteSpec, "gh:")
		repoBase := repo
		if strings.Contains(repoBase, "@") {
			repoBase = strings.SplitN(repoBase, "@", 2)[0]
		}
		zipURL := fmt.Sprintf("https://github.com/%s/archive/refs/heads/%s.zip", repoBase, ref)
		zipTmp := filepath.Join(os.TempDir(), fmt.Sprintf("verbo_%s_%s.zip", nome, strings.ReplaceAll(ref, "/", "_")))
		fmt.Printf("⬇️  Baixando %s...\n", zipURL)
		if err := baixarArquivo(zipURL, zipTmp); err != nil {
			erroFatal(fmt.Sprintf("falha ao baixar pacote: %v", err))
		}
		defer os.Remove(zipTmp)

		tmpExtract := filepath.Join(os.TempDir(), fmt.Sprintf("verbo_%s_extract", nome))
		_ = os.RemoveAll(tmpExtract)
		garantirDiretorio(tmpExtract)
		defer os.RemoveAll(tmpExtract)

		if err := unzipParaPasta(zipTmp, tmpExtract); err != nil {
			erroFatal(fmt.Sprintf("falha ao extrair pacote: %v", err))
		}

		entries, err := os.ReadDir(tmpExtract)
		if err != nil || len(entries) == 0 {
			erroFatal("pacote GitHub vazio/ilegível")
		}
		// GitHub zip normalmente vem com uma pasta raiz
		srcRoot := tmpExtract
		if len(entries) == 1 && entries[0].IsDir() {
			srcRoot = filepath.Join(tmpExtract, entries[0].Name())
		}
		if err := copiarDiretorio(srcRoot, dst); err != nil {
			erroFatal(fmt.Sprintf("falha ao copiar pacote para destino: %v", err))
		}

		lock.Pacotes = append(lock.Pacotes, lockPacote{Nome: nome, Fonte: fonteSpec, Ref: ref, Pasta: dst})
	} else if strings.HasPrefix(fonteSpec, "path:") {
		pathSrc := strings.TrimSpace(strings.TrimPrefix(fonteSpec, "path:"))
		fmt.Printf("📦 Instalando de %s...\n", pathSrc)
		if err := copiarDiretorio(pathSrc, dst); err != nil {
			erroFatal(fmt.Sprintf("falha ao copiar pacote local: %v", err))
		}
		lock.Pacotes = append(lock.Pacotes, lockPacote{Nome: nome, Fonte: fonteSpec, Ref: "", Pasta: dst})
	} else {
		// fallback: manter comportamento MVP
		fmt.Println("ℹ️  Fonte não especificada; registrando apenas no manifesto (sem baixar).")
		lock.Pacotes = append(lock.Pacotes, lockPacote{Nome: nome, Fonte: fonteSpec, Ref: "", Pasta: dst})
	}

	executarInstaller(dst)

	m.Pacotes = append(m.Pacotes, nome)
	salvarManifestoPacotes(m)
	salvarLockPacotes(lock)
	fmt.Printf("✅ Pacote instalado: %s (em %s)\n", nome, dst)
}

func executarPacote(args []string) {
	if len(args) < 1 {
		erroFatal("uso: verbo pacote <instalar|listar|remover> [nome]")
	}
	action := args[0]

	switch action {
	case "instalar":
		if len(args) < 2 {
			erroFatal("uso: verbo pacote instalar <nome>")
		}
		instalarPacote(args[1])

	case "listar":
		m := carregarManifestoPacotes()
		if len(m.Pacotes) == 0 {
			fmt.Println("(nenhum pacote instalado)")
			return
		}
		for _, p := range m.Pacotes {
			fmt.Println(p)
		}

	case "info":
		l := carregarLockPacotes()
		if len(l.Pacotes) == 0 {
			fmt.Println("(nenhum pacote instalado)")
			return
		}
		for _, p := range l.Pacotes {
			fmt.Printf("%s\n  fonte: %s\n  ref: %s\n  pasta: %s\n\n", p.Nome, p.Fonte, p.Ref, p.Pasta)
		}

	case "remover":
		if len(args) < 2 {
			erroFatal("uso: verbo pacote remover <nome>")
		}
		nome := args[1]
		m := carregarManifestoPacotes()
		var novo []string
		removido := false
		for _, p := range m.Pacotes {
			if strings.EqualFold(p, nome) {
				removido = true
				continue
			}
			novo = append(novo, p)
		}
		m.Pacotes = novo
		salvarManifestoPacotes(m)

		// remover pasta e lock
		_ = os.RemoveAll(pacoteDir(nome))
		l := carregarLockPacotes()
		var novoLock []lockPacote
		for _, p := range l.Pacotes {
			if strings.EqualFold(p.Nome, nome) {
				continue
			}
			novoLock = append(novoLock, p)
		}
		l.Pacotes = novoLock
		salvarLockPacotes(l)
		if removido {
			fmt.Printf("✅ Pacote removido: %s\n", nome)
		} else {
			fmt.Printf("ℹ️  Pacote não estava instalado: %s\n", nome)
		}

	default:
		erroFatal("uso: verbo pacote <instalar|listar|info|remover> [nome]")
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

func executarServir(args []string) {
	// Uso: verbo servir <arquivo.vrb> [--host 0.0.0.0] [--porta 5000]
	if len(args) < 1 {
		erroFatal("uso: verbo servir <arquivo.vrb> [--host <ip>] [--porta <n>]")
	}
	arquivo := args[0]
	host := "127.0.0.1"
	porta := "5000"
	for i := 1; i < len(args); i++ {
		if args[i] == "--host" && i+1 < len(args) {
			host = args[i+1]
			i++
			continue
		}
		if (args[i] == "--porta" || args[i] == "--port") && i+1 < len(args) {
			porta = args[i+1]
			i++
			continue
		}
	}
	if _, err := strconv.Atoi(porta); err != nil {
		erroFatal("porta inválida")
	}

	codigoGo := transpilar(arquivo)
	arquivoTemp := filepath.Join(os.TempDir(), "verbo_servir.go")
	if err := os.WriteFile(arquivoTemp, []byte(codigoGo), 0644); err != nil {
		erroFatal(fmt.Sprintf("erro ao escrever arquivo temporário: %v", err))
	}
	defer os.Remove(arquivoTemp)

	fmt.Printf("🚀 Servindo %s em %s:%s...\n", arquivo, host, porta)
	cmd := exec.Command("go", "run", arquivoTemp)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), "VERBO_HOST="+host, "VERBO_PORTA="+porta)
	if err := cmd.Run(); err != nil {
		erroFatal(fmt.Sprintf("erro ao executar servidor: %v", err))
	}
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
  servir <arquivo.vrb>     Executa um servidor web (estilo Flask) com flags --host/--porta
  verificar <arquivo.vrb>  Apenas verifica a sintaxe

  pacote instalar <nome>   Registra um pacote/dependência (MVP)
  pacote listar            Lista pacotes registrados
  pacote remover <nome>    Remove pacote registrado
  pacote info              Exibe info do lockfile de pacotes

  versão                   Exibe a versão do Verbo
  ajuda                    Exibe esta ajuda

Exemplos:
  verbo executar ola_mundo.vrb
  verbo compilar calculadora.vrb
  verbo verificar meu_programa.vrb

Para mais informações: https://github.com/MrJc01/crom-verbo
`, versao)
}
