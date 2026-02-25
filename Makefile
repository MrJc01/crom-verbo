# ============================================
# Verbo — Linguagem de Programação em Português
# ============================================

BINARY_NAME=verbo
BUILD_DIR=build
CMD_DIR=cmd/verbo

.PHONY: build test run clean exemplos verificar

# Compila o binário do CLI
build:
	@echo "🔨 Compilando $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./$(CMD_DIR)
	@echo "✅ Binário gerado em $(BUILD_DIR)/$(BINARY_NAME)"

# Roda todos os testes
test:
	@echo "🧪 Executando testes..."
	go test ./... -v -count=1
	@echo "✅ Todos os testes passaram!"

# Roda um arquivo .vrb específico
# Uso: make run ARQUIVO=examples/ola_mundo.vrb
run: build
	@echo "🚀 Executando $(ARQUIVO)..."
	./$(BUILD_DIR)/$(BINARY_NAME) executar $(ARQUIVO)

# Compila um arquivo .vrb para Go
# Uso: make compilar ARQUIVO=examples/ola_mundo.vrb
compilar: build
	@echo "📝 Compilando $(ARQUIVO)..."
	./$(BUILD_DIR)/$(BINARY_NAME) compilar $(ARQUIVO)

# Verifica sintaxe de um arquivo .vrb
# Uso: make verificar ARQUIVO=examples/ola_mundo.vrb
verificar: build
	@echo "🔍 Verificando $(ARQUIVO)..."
	./$(BUILD_DIR)/$(BINARY_NAME) verificar $(ARQUIVO)

# Roda todos os exemplos
exemplos: build
	@echo "📚 Executando todos os exemplos..."
	@for f in examples/*.vrb; do \
		echo "\n--- $$f ---"; \
		./$(BUILD_DIR)/$(BINARY_NAME) executar $$f; \
	done

# Limpa artefatos de build
clean:
	@echo "🧹 Limpando..."
	@rm -rf $(BUILD_DIR)
	@rm -f output.go
	@echo "✅ Limpo!"

# Instala globalmente
install: build
	@echo "📦 Instalando $(BINARY_NAME)..."
	cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
	@echo "✅ Instalado em /usr/local/bin/$(BINARY_NAME)"

# Cross-compila binários para todas as plataformas
release:
	@echo "📦 Compilando binários para release..."
	@mkdir -p www/downloads
	GOOS=linux GOARCH=amd64 go build -o www/downloads/verbo-linux-amd64 ./$(CMD_DIR)
	GOOS=linux GOARCH=arm64 go build -o www/downloads/verbo-linux-arm64 ./$(CMD_DIR)
	GOOS=darwin GOARCH=amd64 go build -o www/downloads/verbo-darwin-amd64 ./$(CMD_DIR)
	GOOS=darwin GOARCH=arm64 go build -o www/downloads/verbo-darwin-arm64 ./$(CMD_DIR)
	GOOS=windows GOARCH=amd64 go build -o www/downloads/verbo-windows-amd64.exe ./$(CMD_DIR)
	@echo "✅ Binários gerados em www/downloads/"
	@ls -la www/downloads/
