# 🇧🇷 Verbo — Linguagem de Programação em Português

**Verbo** é uma linguagem de programação transpilada que utiliza a gramática da norma culta do Português Brasileiro como sintaxe lógica. Inspirada pelo [Wenyan-lang](https://wy-lang.org/) (chinês clássico) e pela [易语言](https://www.eyuyan.com/) (Yi), a Verbo prova que programação pode ser escrita na nossa língua materna sem perder rigor técnico.

## ✨ Diferencial

- **Artigos como mutabilidade**: `O` / `A` = constante, `Um` / `Uma` = variável mutável
- **Verbos como funções**: Toda função é um verbo no infinitivo
- **Semântica de estado**: `É` (estático) vs `Está` (temporário)
- **Preposições como acesso**: `nome de usuario` em vez de `usuario.nome`
- **100% em Português**: CLI, erros, docs — tudo na nossa língua

## 🚀 Início Rápido

```bash
# Compilar o CLI
make build

# Executar um programa Verbo
make run ARQUIVO=examples/ola_mundo.vrb

# Verificar sintaxe
make verificar ARQUIVO=examples/ola_mundo.vrb

# Rodar todos os testes
make test

# Rodar um servidor web (estilo Flask)
make build
./build/verbo servir examples/web_index.vrb --host 0.0.0.0 --porta 5000
```

## 📝 Exemplo

```
A mensagem é "Olá, Mundo!".
Um contador está 0.

Para Saudar usando (nome: Texto):
    Exibir com ("Bem-vindo, " e nome e "!").

Saudar com ("Brasil").

Repita 5 vezes:
    contador está contador + 1.
    Exibir com (contador).
```

## 📁 Estrutura do Projeto

```
crom-verbo/
├── cmd/verbo/          # CLI principal
├── pkg/
│   ├── lexer/          # Análise léxica (tokenização)
│   ├── parser/         # Análise sintática (AST)
│   ├── ast/            # Árvore de Sintaxe Abstrata
│   └── transpiler/     # Geração de código Go
├── examples/           # Programas de exemplo .vrb
├── tests/              # Testes de integração
└── docs/               # Documentação completa
```

## 📚 Documentação

- [Especificação da Linguagem](docs/ESPECIFICACAO.md)
- [Gramática Formal (EBNF)](docs/GRAMATICA.md)
- [Arquitetura do Compilador](docs/ARQUITETURA.md)
- [Exemplos Comentados](docs/EXEMPLOS.md)
- [Guia de Contribuição](docs/CONTRIBUINDO.md)
- [Roadmap](docs/ROADMAP.md)

## 🌐 Servidor Web (estilo Flask)

O Verbo tem um modo prático para rodar aplicações web, inspirado no fluxo de desenvolvimento do Flask.

```bash
./build/verbo servir examples/web_index.vrb --host 127.0.0.1 --porta 8099
```

Esse comando transpila o `.vrb` e executa o servidor, permitindo override de host/porta via variáveis de ambiente internas.

No código, você pode iniciar o servidor com `app iniciar.` ou usando o açúcar sintático `app rodar.`.

## 📦 Pacotes (manifesto, lockfile e installer)

- **Manifesto**: `verbo.mod.json`
- **Lockfile**: `verbo.lock.json`
- **Pasta de pacotes**: `./pacotes/<nome>`

Instalação:

```bash
# GitHub (baixa zip do repositório)
./build/verbo pacote instalar gh:user/repo@main

# Local (copia uma pasta local)
./build/verbo pacote instalar path:./meu_pacote

./build/verbo pacote listar
./build/verbo pacote info
./build/verbo pacote remover meu_pacote
```

Se o pacote contiver um `installer.vrb`, ele é executado após a instalação e pode emitir um JSON com ações (ex: copiar arquivos).

## 🛠️ Requisitos

- Go 1.22+
- Make

## 📜 Licença

MIT License — Feito com ❤️ no Brasil.
