# 🗺️ Roadmap — Linguagem Verbo

## Estágio 1: MVP ✅ (Atual)

O alicerce funcional da linguagem.

- [x] Lexer com suporte UTF-8
- [x] Parser Recursive Descent
- [x] AST tipada
- [x] Transpiler AST → Go
- [x] CLI (`compilar`, `executar`, `verificar`)
- [x] Tipos: Texto, Inteiro, Decimal, Lógico, Nulo
- [x] Variáveis/Constantes via artigos
- [x] Funções com parâmetros tipados
- [x] Condicionais (Se/Senão)
- [x] Loops (Repita/Enquanto)
- [x] Exemplos e documentação

---

## Estágio 2: Semântica Avançada

Tornar a linguagem mais expressiva.

- [ ] **Estruturas** — `Usuário É uma Estrutura com (nome: Texto, idade: Inteiro)`
- [ ] **Listas nativas** — `Uma frutas é ["maçã", "banana", "uva"]`
- [ ] **Acesso com preposição** — `nome de usuario` (em vez de `usuario.nome`)
- [ ] **Guard Clauses** — `Dado que x é maior que 0.`
- [ ] **Ser vs Estar completo** — Ownership/empréstimo baseado em artigos
- [ ] **Importação de módulos** — `Incluir "matematica".`
- [ ] **Standard Library (BibVerbo)** — Matemática, Strings, Arquivos
- [ ] **Tratamento de erros** — `Tente: ... Se falhar:`

---

## Estágio 3: Ecossistema

Ferramentas de desenvolvimento.

- [ ] **Extensão VSCode** — Syntax highlighting para `.vrb`
- [ ] **LSP (Language Server Protocol)** — Autocompletar, go-to-definition
- [ ] **Playground Web** — Testar Verbo no navegador
- [ ] **REPL interativo** — `verbo repl` para experimentação
- [ ] **Formatter** — `verbo formatar` para padronizar código
- [ ] **Linter** — Sugestões de melhoria de código
- [ ] **Package Manager** — Gerenciador de pacotes Verbo

---

## Estágio 4: Integração Sistêmica

Aplicações práticas.

- [ ] **Concorrência** — `Simultaneamente:` e `Enquanto isso:`
- [ ] **Bindings nativos** — FFI com C/Go
- [ ] **WebAssembly** — Compilar Verbo para Wasm
- [ ] **Contratos Inteligentes** — DSL para smart contracts legíveis
- [ ] **Scripts de sistema** — Automação de tarefas em português
- [ ] **Integração com Crom** — Linguagem de script para o ecossistema Crom

---

## Visão de Longo Prazo

```
2025 Q1: MVP funcional (Estágio 1) ← Estamos aqui
2025 Q2: Semântica avançada (Estágio 2)
2025 Q3: Ecossistema (Estágio 3)
2025 Q4: Integração sistêmica (Estágio 4)
2026+:   Comunidade e adoção
```

---

## Métricas de Sucesso

| Métrica | Meta |
|---------|------|
| Testes passando | 100% |
| Exemplos executáveis | 5+ |
| Documentação | Completa em PT-BR |
| Tempo de compilação | < 1s para arquivos simples |
| Tamanho do binário CLI | < 10MB |
