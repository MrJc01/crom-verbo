<!DOCTYPE html>
<html lang="pt-BR" class="dark">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Verbo Lab — Playground Online</title>
    <meta name="description"
        content="Compile e execute código Verbo diretamente no navegador com syntax highlighting e exemplos pré-carregados.">

    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link
        href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700;800&family=JetBrains+Mono:wght@400;500;600&display=swap"
        rel="stylesheet">

    <script src="https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"></script>
    <style type="text/tailwindcss">
        @theme {
        --color-verde: #009739;
        --color-verde-light: #00c04b;
        --color-verde-neon: #00ff88;
        --color-verde-dark: #006b28;
        --color-amarelo: #FFDF00;
        --color-azul: #012169;
        --color-azul-light: #1a3a8a;
        --color-azul-petroleo: #0a192f;
        --color-azul-noite: #020c1b;
        --color-branco-gelo: #f0f4f8;
        --font-sans: 'Inter', system-ui, sans-serif;
        --font-mono: 'JetBrains Mono', monospace;
      }
    </style>

    <style>
        html,
        body {
            height: 100%;
            margin: 0;
            overflow: hidden;
        }

        .glass-bar {
            background: rgba(2, 12, 27, 0.85);
            backdrop-filter: blur(20px);
            -webkit-backdrop-filter: blur(20px);
            border-bottom: 1px solid rgba(0, 255, 136, 0.06);
        }

        .animate-spin {
            animation: spin 1s linear infinite;
        }

        @keyframes spin {
            to {
                transform: rotate(360deg);
            }
        }

        /* Editor/Output panels */
        #editor-container {
            width: 100%;
            height: 100%;
        }

        .output-panel {
            background: #020c1b;
        }

        /* Tab active style */
        .tab-btn {
            transition: all 0.2s;
            cursor: pointer;
        }

        .tab-btn.active {
            color: #00ff88;
            border-bottom-color: #00ff88;
        }

        .tab-btn:not(.active) {
            opacity: 0.4;
            border-bottom-color: transparent;
        }

        .tab-btn:not(.active):hover {
            opacity: 0.7;
        }

        /* Toolbar button */
        .tool-btn {
            padding: 6px 8px;
            border-radius: 8px;
            transition: all 0.2s;
            opacity: 0.4;
            display: flex;
            align-items: center;
            justify-content: center;
        }

        .tool-btn:hover {
            opacity: 1;
            background: rgba(255, 255, 255, 0.05);
        }

        /* Status colors */
        .status-ready {
            background: rgba(255, 255, 255, 0.2);
        }

        .status-ok {
            background: #00ff88;
            box-shadow: 0 0 8px rgba(0, 255, 136, 0.4);
        }

        .status-err {
            background: #ff5555;
            box-shadow: 0 0 8px rgba(255, 85, 85, 0.4);
        }

        .status-warn {
            background: #FFDF00;
            box-shadow: 0 0 8px rgba(255, 223, 0, 0.4);
        }

        /* Select override */
        .custom-select {
            background: rgba(255, 255, 255, 0.05);
            border: 1px solid rgba(0, 255, 136, 0.12);
            border-radius: 10px;
            padding: 6px 12px;
            font-size: 12px;
            color: inherit;
            cursor: pointer;
            appearance: none;
            -webkit-appearance: none;
            background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='12' height='12' viewBox='0 0 24 24' fill='none' stroke='%2300ff88' stroke-width='2'%3E%3Cpath d='M6 9l6 6 6-6'/%3E%3C/svg%3E");
            background-repeat: no-repeat;
            background-position: right 8px center;
            padding-right: 28px;
        }

        .custom-select:focus {
            outline: none;
            border-color: rgba(0, 255, 136, 0.3);
        }

        .custom-select option {
            background: #0a192f;
            color: #ccd6f6;
        }
    </style>

    <script src="https://cdn.jsdelivr.net/npm/monaco-editor@0.45.0/min/vs/loader.js"></script>
</head>

<body class="bg-azul-noite text-branco-gelo font-sans">

    <div class="h-full flex flex-col">

        <!-- ============================================
       TOOLBAR
       ============================================ -->
        <header class="glass-bar flex-shrink-0 z-20">
            <div class="flex items-center justify-between px-3 sm:px-5 h-14">
                <!-- Left -->
                <div class="flex items-center gap-3">
                    <a href="index.php" class="flex items-center gap-2 group" title="Voltar ao Portal">
                        <div
                            class="w-8 h-8 rounded-lg bg-gradient-to-br from-verde to-verde-dark flex items-center justify-center text-white font-mono font-bold text-sm shadow-lg shadow-verde/20">
                            V</div>
                        <span class="font-bold text-sm hidden sm:inline">Verbo Lab</span>
                    </a>
                    <span class="text-white/10 hidden sm:inline">│</span>
                    <a href="docs.php"
                        class="text-xs text-white/30 hover:text-verde-neon transition-colors hidden sm:inline">📖
                        Docs</a>
                </div>

                <!-- Center -->
                <div class="flex items-center gap-3">
                    <select id="exemplo-select" onchange="carregarExemplo(this.value)" class="custom-select font-mono">
                        <option value="" disabled>📂 Exemplos</option>
                        <option value="ola_mundo" selected>Olá Mundo</option>
                        <option value="fibonacci">Fibonacci</option>
                        <option value="calculadora">Calculadora</option>
                        <option value="concorrencia">Concorrência</option>
                        <option value="entidade">Entidades</option>
                        <option value="listas">Listas</option>
                        <option value="erros">Erros</option>
                        <option value="canais">Canais</option>
                    </select>

                    <button id="btn-run" onclick="executarCodigo()"
                        class="flex items-center gap-2 px-5 py-2 rounded-xl bg-verde text-white text-sm font-bold hover:bg-verde-light hover:shadow-lg hover:shadow-verde/30 active:scale-95 transition-all">
                        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                            <path d="M8 5v14l11-7z" />
                        </svg>
                        <span class="hidden sm:inline">Executar</span>
                    </button>
                </div>

                <!-- Right -->
                <div class="flex items-center gap-0.5">
                    <button onclick="copiarCodigo()" class="tool-btn" title="Copiar código">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
                        </svg>
                    </button>
                    <button onclick="baixarCodigo()" class="tool-btn" title="Baixar .vrb">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
                        </svg>
                    </button>
                    <button onclick="compartilhar()" class="tool-btn" title="Compartilhar link">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1" />
                        </svg>
                    </button>
                    <div class="w-px h-5 bg-white/5 mx-1"></div>
                    <button onclick="verboTheme.toggle()" class="tool-btn" title="Alternar tema">
                        <svg data-theme-icon="light" class="w-4 h-4 hidden" fill="none" stroke="currentColor"
                            viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
                        </svg>
                        <svg data-theme-icon="dark" class="w-4 h-4" fill="none" stroke="currentColor"
                            viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
                        </svg>
                    </button>
                </div>
            </div>
            <!-- Shortcut hint -->
            <div class="text-center text-[10px] text-white/10 pb-1 font-mono hidden sm:block">Ctrl+Enter para executar
            </div>
        </header>

        <!-- ============================================
       SPLIT LAYOUT
       ============================================ -->
        <div class="flex-1 flex flex-col lg:flex-row min-h-0">

            <!-- EDITOR -->
            <div class="flex-1 min-h-0 relative border-b lg:border-b-0 lg:border-r border-verde-neon/5"
                style="min-height:250px">
                <!-- Loading -->
                <div id="editor-loading" class="absolute inset-0 flex items-center justify-center bg-azul-noite z-10">
                    <div class="text-center">
                        <div
                            class="w-12 h-12 rounded-2xl bg-verde/10 border border-verde/20 flex items-center justify-center mx-auto mb-4">
                            <span class="text-verde-neon font-mono font-bold text-lg animate-pulse">V</span>
                        </div>
                        <p class="text-sm text-white/30">Carregando Monaco Editor...</p>
                    </div>
                </div>
                <div id="editor-container" style="display:none"></div>
            </div>

            <!-- OUTPUT -->
            <div class="flex-1 min-h-0 flex flex-col output-panel" style="min-height:200px">
                <!-- Tabs -->
                <div class="flex items-center justify-between border-b border-verde-neon/5 px-4 flex-shrink-0">
                    <div class="flex gap-0">
                        <button data-tab="console" onclick="switchTab('console')"
                            class="tab-btn active px-4 py-3 text-xs font-semibold border-b-2">
                            <span class="mr-1">⬤</span> Console
                        </button>
                        <button data-tab="go" onclick="switchTab('go')"
                            class="tab-btn px-4 py-3 text-xs font-semibold border-b-2">
                            Código Go
                        </button>
                        <button data-tab="html" onclick="switchTab('html')"
                            class="tab-btn px-4 py-3 text-xs font-semibold border-b-2">
                            Resultado HTML
                        </button>
                    </div>
                    <!-- Status -->
                    <div class="flex items-center gap-2 text-xs">
                        <div id="status-indicator" class="w-2.5 h-2.5 rounded-full status-ready transition-all"></div>
                        <span id="status-text" class="text-white/40">Pronto</span>
                        <span id="status-time" class="text-white/20 font-mono text-[10px]"></span>
                    </div>
                </div>

                <!-- Output content -->
                <div class="flex-1 min-h-0 overflow-auto">
                    <pre id="output-console" data-panel="console"
                        class="p-5 font-mono text-sm leading-relaxed h-full overflow-auto whitespace-pre-wrap text-white/50">Clique em ▶ Executar ou pressione Ctrl+Enter para compilar e rodar seu código Verbo.

📂 Exemplos disponíveis no menu acima ↑
📖 Documentação em docs.php</pre>

                    <pre id="output-go" data-panel="go" style="display:none"
                        class="p-5 font-mono text-sm leading-relaxed h-full overflow-auto whitespace-pre-wrap text-white/50">// O código Go transpilado aparecerá aqui após a execução.</pre>

                    <div data-panel="html" style="display:none" class="h-full">
                        <iframe id="output-html" class="w-full h-full border-0 bg-white"
                            sandbox="allow-scripts"></iframe>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <!-- Toast -->
    <div id="toast"
        class="fixed bottom-6 right-6 bg-azul-noite/95 border border-verde-neon/20 text-verde-neon px-5 py-3 rounded-2xl text-sm font-medium shadow-xl shadow-verde-neon/5 opacity-0 translate-y-4 transition-all duration-300 z-50 backdrop-blur-lg">
    </div>

    <script src="js/theme.js"></script>
    <script src="js/verbo-lang.js"></script>
    <script src="js/playground.js"></script>
    <script>document.addEventListener('DOMContentLoaded', initPlayground);</script>
</body>

</html>