# ✅ Checklist Master — Projeto Verbo

Use esta checklist para acompanhar o progresso do desenvolvimento.

---

## Fase 0: Planejamento & Documentação
- [x] Simular painel de 100 especialistas
- [x] Criar plano de implementação detalhado
- [x] Definir especificação da linguagem
- [x] Definir gramática formal (EBNF)
- [x] Documentar arquitetura do compilador

## Fase 1: Alicerce (Scaffold)
- [x] Inicializar módulo Go (`go.mod`)
- [x] Criar Makefile com targets
- [x] Criar README.md do projeto
- [x] Criar árvore de diretórios

## Fase 2: Lexer (Análise Léxica)
- [x] Definir todos os TokenTypes
- [x] Mapear palavras reservadas do português
- [x] Implementar scanner UTF-8
- [x] Suporte a acentuação (é, ã, ç, etc.)
- [x] Tokenizar literais (números, textos)
- [x] Tokenizar operadores e delimitadores
- [x] Ignorar comentários
- [x] Rastreamento de linha/coluna
- [x] Testes unitários do Lexer

## Fase 3: Parser + AST
- [x] Definir nós da AST
- [x] Implementar Recursive Descent Parser
- [x] Analisar declaração de variáveis (artigos)
- [x] Analisar declaração de funções
- [x] Analisar condicionais (Se/Senão)
- [x] Analisar loops (Repita/Enquanto)
- [x] Analisar expressões com precedência
- [x] Analisar chamadas de função
- [x] Testes unitários do Parser

## Fase 4: Transpiler (AST → Go)
- [x] Mapear tipos Verbo → Go
- [x] Mapear operadores Verbo → Go
- [x] Transpilar variáveis e constantes
- [x] Transpilar funções
- [x] Transpilar condicionais
- [x] Transpilar loops
- [x] Transpilar expressões
- [x] Gerar `package main` e `import`
- [x] Testes end-to-end do Transpiler

## Fase 5: CLI
- [x] Comando `compilar` (.vrb → binário)
- [x] Comando `executar` (.vrb → execução direta)
- [x] Comando `verificar` (validação sintática)
- [x] Help em português
- [x] Mensagens de erro em português

## Fase 6: Exemplos
- [x] `ola_mundo.vrb` — Hello World
- [x] `fibonacci.vrb` — Sequência de Fibonacci
- [x] `calculadora.vrb` — Funções aritméticas
- [x] `contador.vrb` — Loop com reatribuição
- [x] `saudacao.vrb` — Funções e chamadas

## Fase 7: Documentação
- [x] `docs/ESPECIFICACAO.md` — Especificação completa
- [x] `docs/GRAMATICA.md` — Gramática formal EBNF
- [x] `docs/ARQUITETURA.md` — Arquitetura do compilador
- [x] `docs/CHECKLIST.md` — Este arquivo
- [x] `docs/ESPECIALISTAS.md` — Simulação de especialistas
- [x] `docs/EXEMPLOS.md` — Galeria de exemplos
- [x] `docs/CONTRIBUINDO.md` — Guia de contribuição
- [x] `docs/ROADMAP.md` — Roadmap futuro

## Fase 8: Verificação
- [ ] Todos os testes passando (`go test ./...`)
- [ ] Build do binário (`make build`)
- [ ] Executar exemplos (`make exemplos`)
- [ ] Verificar mensagens de erro em português
