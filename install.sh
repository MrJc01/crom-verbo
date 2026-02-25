#!/bin/bash
# ============================================================
# install.sh — Instalador do Verbo
#
# Uso:
#   curl -fsSL https://raw.githubusercontent.com/MrJc01/crom-verbo/main/install.sh | bash
#
# Suporta: Linux (amd64/arm64), macOS (amd64/arm64), Windows (via Git Bash/WSL)
# ============================================================

set -euo pipefail

# === Configuração ===
REPO="MrJc01/crom-verbo"
INSTALL_DIR="${VERBO_INSTALL_DIR:-/usr/local/bin}"
TMP_DIR=$(mktemp -d)

# Cores
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
NC='\033[0m'

cleanup() {
    rm -rf "$TMP_DIR"
}
trap cleanup EXIT

info()  { echo -e "${CYAN}ℹ${NC}  $1"; }
ok()    { echo -e "${GREEN}✅${NC} $1"; }
warn()  { echo -e "${YELLOW}⚠${NC}  $1"; }
error() { echo -e "${RED}❌${NC} $1"; exit 1; }

# === Detectar OS e Arquitetura ===
detect_platform() {
    local os arch

    case "$(uname -s)" in
        Linux*)   os="linux" ;;
        Darwin*)  os="darwin" ;;
        MINGW*|MSYS*|CYGWIN*) os="windows" ;;
        *)        error "Sistema operacional não suportado: $(uname -s)" ;;
    esac

    case "$(uname -m)" in
        x86_64|amd64)  arch="amd64" ;;
        aarch64|arm64) arch="arm64" ;;
        *)             error "Arquitetura não suportada: $(uname -m)" ;;
    esac

    echo "${os}-${arch}"
}

# === Banner ===
echo ""
echo -e "${GREEN}╔══════════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║${NC}   🇧🇷 ${CYAN}Instalador do Verbo${NC}                      ${GREEN}║${NC}"
echo -e "${GREEN}║${NC}   Linguagem de Programação em Português     ${GREEN}║${NC}"
echo -e "${GREEN}╚══════════════════════════════════════════════╝${NC}"
echo ""

# === Detectar plataforma ===
PLATFORM=$(detect_platform)
OS=$(echo "$PLATFORM" | cut -d'-' -f1)
ARCH=$(echo "$PLATFORM" | cut -d'-' -f2)

info "Plataforma detectada: ${CYAN}${OS}/${ARCH}${NC}"

# === Buscar última release ===
info "Buscando última versão..."

LATEST_URL="https://api.github.com/repos/${REPO}/releases/latest"
RELEASE_JSON=$(curl -fsSL "$LATEST_URL" 2>/dev/null || echo "")

if [ -z "$RELEASE_JSON" ] || echo "$RELEASE_JSON" | grep -q '"message"'; then
    # Sem releases — baixar do branch main como fallback
    warn "Nenhuma release encontrada. Baixando binário do branch main..."
    
    BINARY_NAME="verbo-${PLATFORM}"
    [ "$OS" = "windows" ] && BINARY_NAME="${BINARY_NAME}.exe"
    
    DOWNLOAD_URL="https://github.com/${REPO}/raw/main/downloads/${BINARY_NAME}"
else
    VERSION=$(echo "$RELEASE_JSON" | grep '"tag_name"' | head -1 | sed 's/.*"tag_name": *"\([^"]*\)".*/\1/')
    info "Versão: ${GREEN}${VERSION}${NC}"
    
    BINARY_NAME="verbo-${PLATFORM}"
    [ "$OS" = "windows" ] && BINARY_NAME="${BINARY_NAME}.exe"
    
    # Procurar asset na release
    DOWNLOAD_URL=$(echo "$RELEASE_JSON" | grep "browser_download_url" | grep "$BINARY_NAME" | head -1 | sed 's/.*"\(https[^"]*\)".*/\1/')
    
    if [ -z "$DOWNLOAD_URL" ]; then
        warn "Binário não encontrado na release. Tentando downloads/..."
        DOWNLOAD_URL="https://github.com/${REPO}/raw/main/downloads/${BINARY_NAME}"
    fi
fi

# === Download ===
info "Baixando ${CYAN}${BINARY_NAME}${NC}..."

DEST_FILE="$TMP_DIR/verbo"
[ "$OS" = "windows" ] && DEST_FILE="$TMP_DIR/verbo.exe"

HTTP_CODE=$(curl -fsSL -w '%{http_code}' -o "$DEST_FILE" "$DOWNLOAD_URL" 2>/dev/null || echo "000")

if [ "$HTTP_CODE" != "200" ] || [ ! -s "$DEST_FILE" ]; then
    error "Falha ao baixar o binário (HTTP $HTTP_CODE).
   URL: $DOWNLOAD_URL
   
   Tente manualmente:
     git clone https://github.com/${REPO}.git
     cd crom-verbo && make build
     sudo mv build/verbo /usr/local/bin/"
fi

# === Instalar ===
chmod +x "$DEST_FILE"

if [ "$OS" = "windows" ]; then
    # Windows: copiar para diretório do usuário
    WIN_DIR="$HOME/.verbo/bin"
    mkdir -p "$WIN_DIR"
    cp "$DEST_FILE" "$WIN_DIR/verbo.exe"
    ok "Binário instalado em: ${CYAN}${WIN_DIR}/verbo.exe${NC}"
    warn "Adicione ao PATH: export PATH=\"\$HOME/.verbo/bin:\$PATH\""
else
    # Linux/macOS: instalar em /usr/local/bin ou diretório custom
    if [ -w "$INSTALL_DIR" ]; then
        cp "$DEST_FILE" "$INSTALL_DIR/verbo"
    else
        info "Instalando em ${CYAN}${INSTALL_DIR}${NC} (requer sudo)..."
        sudo cp "$DEST_FILE" "$INSTALL_DIR/verbo"
        sudo chmod +x "$INSTALL_DIR/verbo"
    fi
    ok "Binário instalado em: ${CYAN}${INSTALL_DIR}/verbo${NC}"
fi

# === Verificar instalação ===
echo ""
if command -v verbo &>/dev/null; then
    INSTALLED_VERSION=$(verbo --version 2>/dev/null || echo "desconhecida")
    ok "Verbo instalado com sucesso!"
    info "Versão: ${GREEN}${INSTALLED_VERSION}${NC}"
else
    warn "Binário instalado mas 'verbo' não foi encontrado no PATH."
    warn "Reinicie o terminal ou adicione ${INSTALL_DIR} ao PATH."
fi

# === Instruções ===
echo ""
echo -e "${GREEN}╔══════════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║${NC}   🚀 ${CYAN}Comece agora!${NC}                            ${GREEN}║${NC}"
echo -e "${GREEN}╠══════════════════════════════════════════════╣${NC}"
echo -e "${GREEN}║${NC}                                              ${GREEN}║${NC}"
echo -e "${GREEN}║${NC}   echo 'Exibir com (\"Olá!\").' > ola.vrb     ${GREEN}║${NC}"
echo -e "${GREEN}║${NC}   verbo executar ola.vrb                     ${GREEN}║${NC}"
echo -e "${GREEN}║${NC}                                              ${GREEN}║${NC}"
echo -e "${GREEN}║${NC}   📖 Docs: https://crom.run/verbo/docs       ${GREEN}║${NC}"
echo -e "${GREEN}║${NC}   ⚡ Play: https://crom.run/verbo/playground  ${GREEN}║${NC}"
echo -e "${GREEN}╚══════════════════════════════════════════════╝${NC}"
echo ""
