# 🤝 Guia de Contribuição — Projeto Verbo

Obrigado por considerar contribuir com o projeto Verbo! Este guia explica como participar.

---

## Pré-requisitos

- **Go** 1.22+
- **Make**
- **Git**

## Setup do Ambiente

```bash
# Clonar o repositório
git clone https://github.com/juanxto/crom-verbo.git
cd crom-verbo

# Verificar que tudo compila
make build

# Rodar os testes
make test
```

## Fluxo de Trabalho

1. **Fork** o repositório
2. Crie uma **branch** descritiva: `git checkout -b feat/suporte-listas`
3. Faça suas modificações
4. Rode os testes: `make test`
5. Faça commit com mensagem clara: `git commit -m "feat: adiciona suporte a listas"`
6. Abra um **Pull Request**

## Convenções de Commit

Usamos [Conventional Commits](https://www.conventionalcommits.org/):

| Prefixo    | Descrição                    |
|------------|------------------------------|
| `feat:`    | Nova funcionalidade          |
| `fix:`     | Correção de bug              |
| `docs:`    | Documentação                 |
| `test:`    | Testes                       |
| `refactor:`| Refatoração sem mudar comportamento |
| `chore:`   | Tarefas de manutenção        |

## Estrutura do Código

- **`pkg/lexer/`** — Adicionar novos tokens? Edite `token.go` e `lexer.go`
- **`pkg/ast/`** — Novo tipo de nó? Adicione em `ast.go`
- **`pkg/parser/`** — Nova regra gramatical? Adicione método `analisarDeclaracao*`
- **`pkg/transpiler/`** — Novo mapeamento? Adicione caso no `transpilarDeclaracao`
- **`examples/`** — Todo novo recurso deve ter um exemplo `.vrb`
- **`docs/`** — Atualize a documentação quando adicionar features

## Testes

Todo novo código **deve** incluir testes:

```bash
# Rodar testes de um pacote específico
go test ./pkg/lexer/... -v

# Rodar todos os testes com detalhes
go test ./... -v -count=1
```

## Estilo de Código

- Go padrão (use `gofmt`)
- Nomes de variáveis/funções internas em **português** quando fazem parte da API da linguagem
- Comentários de documentação em **português**
- Código de infraestrutura (Go) pode usar nomes em inglês internamente

## Áreas que Precisam de Ajuda

- [ ] Suporte a estruturas (`Estrutura`)
- [ ] Suporte a listas nativas
- [ ] Mensagens de erro mais detalhadas
- [ ] Syntax highlighting para VSCode
- [ ] Language Server Protocol (LSP)
- [ ] Playground web
- [ ] Mais exemplos `.vrb`

## Código de Conduta

Seja respeitoso. A Verbo é um projeto para a comunidade lusófona de tecnologia.

---

Dúvidas? Abra uma [issue](https://github.com/juanxto/crom-verbo/issues)!
