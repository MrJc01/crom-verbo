/**
 * Verbo Lab — Playground Logic
 * Handles Monaco initialization, examples, compilation, and sharing.
 */

// ============================================
// Example Programs
// ============================================
const EXEMPLOS = {
    ola_mundo: {
        nome: 'Olá Mundo',
        codigo: `// Meu primeiro programa em Verbo
A saudacao é "Olá, Mundo!".
Exibir com (saudacao).`
    },
    fibonacci: {
        nome: 'Fibonacci',
        codigo: `// Sequência de Fibonacci
Para Fibonacci usando (n: Inteiro):
    Se n for menor que 2, então:
        Retorne n.
    Retorne Fibonacci com (n - 1) + Fibonacci com (n - 2).

Repita 10 vezes:
    Exibir com (Fibonacci com (i)).`
    },
    calculadora: {
        nome: 'Calculadora',
        codigo: `// Calculadora básica
Para Somar usando (a: Inteiro, b: Inteiro):
    Retorne a + b.

Para Multiplicar usando (a: Inteiro, b: Inteiro):
    Retorne a * b.

O resultado é Somar com (10, 20).
Exibir com ("Soma: ").
Exibir com (resultado).

O produto é Multiplicar com (5, 7).
Exibir com ("Produto: ").
Exibir com (produto).`
    },
    concorrencia: {
        nome: 'Concorrência',
        codigo: `// Execução simultânea com goroutines
Simultaneamente:
    Exibir com ("Tarefa 1: Processando dados...").
    Exibir com ("Tarefa 2: Enviando relatório...").
    Exibir com ("Tarefa 3: Atualizando cache...").

Exibir com ("Todas as tarefas completadas!").`
    },
    entidade: {
        nome: 'Entidades',
        codigo: `// Entidades (structs) em Verbo
A Entidade Pessoa contendo (Nome: Texto, Idade: Inteiro, Saldo: Decimal).

Um cliente é Pessoa com ("Ada Lovelace", 36, 1500.50).
Exibir com ("Cliente cadastrado:").
Exibir com (cliente).`
    },
    listas: {
        nome: 'Listas',
        codigo: `// Listas literais e iteração
Uma precos é [10, 20, 30, 40, 50].
O primeiro_preco é precos[0].
Exibir com ("Primeiro preço:").
Exibir com (primeiro_preco).

Exibir com ("Todos os preços:").
Repita para cada preco em precos:
    Exibir com (preco).`
    },
    erros: {
        nome: 'Tratamento de Erros',
        codigo: `// Tratamento de erros com Tente/Capture
Tente:
    Exibir com ("Tentando operação arriscada...").
    Sinalize com ("Saldo insuficiente!").
Capture erro:
    Exibir com ("Operação falhou. Motivo:").
    Exibir com (erro).

Exibir com ("Programa continua normalmente.").`
    },
    canais: {
        nome: 'Canais',
        codigo: `// Canais para comunicação entre goroutines
Uma via é um Canal de Inteiros.
Enviar 42 para via.
O valor é Receber de via.
Exibir com ("Valor recebido do canal:").
Exibir com (valor).`
    }
};

// ============================================
// Global State
// ============================================
let editor = null;
let currentTab = 'console';

// ============================================
// Initialization
// ============================================
function initPlayground() {
    require.config({ paths: { vs: 'https://cdn.jsdelivr.net/npm/monaco-editor@0.45.0/min/vs' } });

    require(['vs/editor/editor.main'], function () {
        registerVerboLanguage(monaco);
        registerVerboThemes(monaco);

        const isDark = document.documentElement.classList.contains('dark');

        editor = monaco.editor.create(document.getElementById('editor-container'), {
            value: EXEMPLOS.ola_mundo.codigo,
            language: 'verbo',
            theme: isDark ? 'verbo-dark' : 'verbo-light',
            fontFamily: "'JetBrains Mono', 'Fira Code', monospace",
            fontSize: 14,
            lineHeight: 24,
            minimap: { enabled: false },
            padding: { top: 16, bottom: 16 },
            scrollBeyondLastLine: false,
            automaticLayout: true,
            tabSize: 4,
            renderLineHighlight: 'all',
            bracketPairColorization: { enabled: true },
            suggestOnTriggerCharacters: true,
            wordWrap: 'on',
            cursorBlinking: 'smooth',
            cursorSmoothCaretAnimation: 'on',
            smoothScrolling: true,
            lineNumbers: 'on',
            glyphMargin: false,
            folding: true,
            links: false,
            overviewRulerBorder: false,
            scrollbar: {
                verticalScrollbarSize: 8,
                horizontalScrollbarSize: 8,
            },
        });

        // Ctrl+Enter to run
        editor.addCommand(monaco.KeyMod.CtrlCmd | monaco.KeyCode.Enter, executarCodigo);

        // Theme observer
        const observer = new MutationObserver(() => {
            const dark = document.documentElement.classList.contains('dark');
            monaco.editor.setTheme(dark ? 'verbo-dark' : 'verbo-light');
        });
        observer.observe(document.documentElement, { attributes: true, attributeFilter: ['class'] });

        loadFromURL();

        // Show editor
        document.getElementById('editor-loading').style.display = 'none';
        document.getElementById('editor-container').style.display = 'block';
    });
}

// ============================================
// Load from URL
// ============================================
function loadFromURL() {
    const params = new URLSearchParams(window.location.search);

    const code = params.get('code');
    if (code && editor) {
        try { editor.setValue(atob(code)); } catch (e) { }
        return;
    }

    const exemplo = params.get('exemplo');
    if (exemplo && EXEMPLOS[exemplo]) {
        editor.setValue(EXEMPLOS[exemplo].codigo);
        document.getElementById('exemplo-select').value = exemplo;
    }
}

// ============================================
// Example loading
// ============================================
function carregarExemplo(key) {
    if (EXEMPLOS[key] && editor) {
        editor.setValue(EXEMPLOS[key].codigo);
        editor.focus();
        // Update URL without reload
        const url = new URL(window.location);
        url.searchParams.set('exemplo', key);
        history.replaceState(null, '', url);
    }
}

// ============================================
// Tab switching
// ============================================
function switchTab(tab) {
    currentTab = tab;
    document.querySelectorAll('[data-tab]').forEach(el => {
        el.classList.toggle('active', el.dataset.tab === tab);
    });
    document.querySelectorAll('[data-panel]').forEach(el => {
        el.style.display = el.dataset.panel === tab ? 'block' : 'none';
    });
}

// ============================================
// Code Execution
// ============================================
async function executarCodigo() {
    if (!editor) return;

    const codigo = editor.getValue();
    const btnRun = document.getElementById('btn-run');
    const consoleEl = document.getElementById('output-console');
    const goPanel = document.getElementById('output-go');
    const statusIndicator = document.getElementById('status-indicator');
    const statusText = document.getElementById('status-text');
    const statusTime = document.getElementById('status-time');

    // Loading state
    btnRun.disabled = true;
    btnRun.innerHTML = '<span class="animate-spin inline-block w-4 h-4 border-2 border-white/20 border-t-white rounded-full"></span>';
    consoleEl.textContent = '⏳ Compilando e executando...';
    statusIndicator.className = 'w-2.5 h-2.5 rounded-full status-warn transition-all';
    statusText.textContent = 'Compilando';
    statusTime.textContent = '';

    const startTime = performance.now();

    try {
        // 1. Execute code
        const response = await fetch('api/compilar.php', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ codigo }),
        });

        const data = await response.json();
        const elapsed = ((performance.now() - startTime) / 1000).toFixed(2);
        statusTime.textContent = `${elapsed}s`;

        if (data.erro) {
            statusIndicator.className = 'w-2.5 h-2.5 rounded-full status-err transition-all';
            statusText.textContent = 'Erro';
            consoleEl.innerHTML = `<span style="color:#ff5555">❌ ${escapeHTML(data.erro)}</span>`;
        } else {
            statusIndicator.className = 'w-2.5 h-2.5 rounded-full status-ok transition-all';
            statusText.textContent = 'Sucesso';
            consoleEl.textContent = data.output || '(sem saída)';

            // HTML output detection
            if (data.output && data.output.trim().startsWith('<')) {
                document.getElementById('output-html').srcdoc = data.output;
                switchTab('html');
            }
        }

        // 2. Get transpiled Go code
        try {
            const goResponse = await fetch('api/transpilar.php', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ codigo }),
            });
            const goData = await goResponse.json();
            goPanel.textContent = goData.go_code || goData.erro || '// Transpilação indisponível';
        } catch (e) {
            goPanel.textContent = '// Transpilação indisponível: ' + e.message;
        }

    } catch (e) {
        statusIndicator.className = 'w-2.5 h-2.5 rounded-full status-warn transition-all';
        statusText.textContent = 'Offline';
        consoleEl.innerHTML = `<span style="color:#FFDF00">⚠ Servidor indisponível.

O Playground requer o binário 'verbo' instalado no servidor.

Para testar localmente:
  1. make build
  2. ./build/verbo executar seu_arquivo.vrb

Ou acesse o Playground online quando disponível.</span>`;
    }

    // Reset button
    btnRun.disabled = false;
    btnRun.innerHTML = '<svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24"><path d="M8 5v14l11-7z"/></svg><span class="hidden sm:inline"> Executar</span>';

    if (currentTab !== 'html') switchTab('console');
}

// ============================================
// Utilities
// ============================================
function copiarCodigo() {
    if (!editor) return;
    navigator.clipboard.writeText(editor.getValue());
    showToast('📋 Código copiado!');
}

function baixarCodigo() {
    if (!editor) return;
    const blob = new Blob([editor.getValue()], { type: 'text/plain' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = 'programa.vrb';
    a.click();
    URL.revokeObjectURL(url);
    showToast('⬇ Arquivo baixado!');
}

function compartilhar() {
    if (!editor) return;
    const encoded = btoa(unescape(encodeURIComponent(editor.getValue())));
    const url = `${window.location.origin}${window.location.pathname}?code=${encoded}`;
    navigator.clipboard.writeText(url);
    showToast('🔗 Link copiado!');
}

function showToast(msg) {
    const toast = document.getElementById('toast');
    toast.textContent = msg;
    toast.classList.remove('opacity-0', 'translate-y-4');
    toast.classList.add('opacity-100', 'translate-y-0');
    setTimeout(() => {
        toast.classList.remove('opacity-100', 'translate-y-0');
        toast.classList.add('opacity-0', 'translate-y-4');
    }, 2500);
}

function escapeHTML(str) {
    return str.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
}
