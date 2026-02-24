#!/bin/bash
# ============================================
# Script de Teste Completo — Projeto Verbo
# ============================================
# Executa todos os testes e verificações do projeto.
# Uso: bash tests/rodar_testes.sh

set -e

COR_VERDE='\033[0;32m'
COR_VERMELHA='\033[0;31m'
COR_AMARELA='\033[0;33m'
COR_AZUL='\033[0;34m'
SEM_COR='\033[0m'

RAIZ="$(cd "$(dirname "$0")/.." && pwd)"

echo -e "${COR_AZUL}🧪 Iniciando testes do Projeto Verbo...${SEM_COR}"
echo "──────────────────────────────────────────"

# 1. Testes unitários
echo -e "\n${COR_AMARELA}📋 Fase 1: Testes Unitários${SEM_COR}"
echo "──────────────────────────────────────────"
cd "$RAIZ"
go test ./... -v -count=1

echo -e "\n${COR_VERDE}✅ Testes unitários passaram!${SEM_COR}"

# 2. Build do binário
echo -e "\n${COR_AMARELA}📋 Fase 2: Build do Binário${SEM_COR}"
echo "──────────────────────────────────────────"
make build
echo -e "${COR_VERDE}✅ Build bem-sucedido!${SEM_COR}"

# 3. Verificação de sintaxe dos exemplos
echo -e "\n${COR_AMARELA}📋 Fase 3: Verificação de Sintaxe${SEM_COR}"
echo "──────────────────────────────────────────"
for arquivo in "$RAIZ"/examples/*.vrb; do
    nome=$(basename "$arquivo")
    echo -n "  Verificando $nome... "
    if "$RAIZ/build/verbo" verificar "$arquivo" > /dev/null 2>&1; then
        echo -e "${COR_VERDE}OK${SEM_COR}"
    else
        echo -e "${COR_VERMELHA}FALHA${SEM_COR}"
        "$RAIZ/build/verbo" verificar "$arquivo"
        exit 1
    fi
done
echo -e "${COR_VERDE}✅ Todos os exemplos sintaticamente corretos!${SEM_COR}"

# 4. Execução dos exemplos
echo -e "\n${COR_AMARELA}📋 Fase 4: Execução dos Exemplos${SEM_COR}"
echo "──────────────────────────────────────────"
for arquivo in "$RAIZ"/examples/*.vrb; do
    nome=$(basename "$arquivo")
    echo -e "\n${COR_AZUL}--- $nome ---${SEM_COR}"
    "$RAIZ/build/verbo" executar "$arquivo" 2>&1 || {
        echo -e "${COR_VERMELHA}❌ Falha ao executar $nome${SEM_COR}"
    }
done

# Resumo
echo -e "\n══════════════════════════════════════════"
echo -e "${COR_VERDE}🎉 Todos os testes completados!${SEM_COR}"
echo "══════════════════════════════════════════"
