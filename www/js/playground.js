/**
 * Verbo Lab — Playground Logic
 * Handles Monaco initialization, examples, compilation, and sharing.
 */

// ============================================
// Example Programs
// ============================================
const EXEMPLOS = {
    html_perfil: {
        nome: 'HTML: Cartão de Perfil',
        codigo: `// Gerando um Cartão de Perfil usando a Biblioteca Html
Incluir Html.

A estilo é "
    .card { background: #1e293b; color: white; padding: 2rem; border-radius: 1rem; width: 300px; text-align: center; font-family: sans-serif; box-shadow: 0 10px 15px -3px rgba(0,0,0,0.5); border: 1px solid #334155; }
    .avatar { width: 100px; height: 100px; border-radius: 50%; border: 3px solid #10b981; margin-bottom: 1rem; }
    .badge { background: #10b981; color: #022c22; padding: 4px 12px; border-radius: 999px; font-size: 12px; font-weight: bold; }
".

Um avatar é CriarImagem de Html com ("https://i.pravatar.cc/150?u=dev", "Avatar").
Um avatarTag é CriarElementoComAtributos de Html com ("div", Atributo de Html com ("class", "avatar-wrapper"), avatar).
Um nome é CriarElemento de Html com ("h2", "Dev Verbo").
Um badge é CriarElementoComAtributos de Html com ("span", Atributo de Html com ("class", "badge"), "Engenheiro de Software").
Um desc é CriarElemento de Html com ("p", "Apaixonado por código legível e arquitetura limpa.").

// Agrupando Elementos
Um cartaoConteudo é ListaElementos de Html com (avatarTag, nome, badge, desc).
Um cartao é CriarElementoComAtributos de Html com ("div", Atributo de Html com ("class", "card"), cartaoConteudo).

Um site é CriarPaginaComEstilo de Html com ("Perfil Web", estilo, cartao).
Exibir com (site).`
    },
    html_tarefas: {
        nome: 'HTML: Lista de Tarefas',
        codigo: `// Criando uma lista dinâmica de tarefas com a Stdlib Html
Incluir Html.

A estilo é "
    .todo-list { background: #0f172a; padding: 1.5rem; border-radius: 12px; width: 350px; font-family: sans-serif; color: white; border: 1px solid #1e293b; }
    .title { color: #facc15; font-size: 1.2rem; margin-top: 0; }
    ul { list-style: none; padding: 0; margin: 0; }
    li { background: #1e293b; padding: 12px; margin-bottom: 8px; border-radius: 6px; }
".

Um titulo é CriarElementoComAtributos de Html com ("h3", Atributo de Html com ("class", "title"), "Minhas Tarefas").

Um i1 é "Estudar Verbo 2.0".
Um i2 é "Criar exemplos HTML".
Um i3 é "Atualizar documentação".
Um i4 é "Tomar café".

Um items é CriarLista de Html com (i1, i2, i3, i4).

Um bloco é ListaElementos de Html com (titulo, items).
Um cartao é CriarElementoComAtributos de Html com ("div", Atributo de Html com ("class", "todo-list"), bloco).

Um site é CriarPaginaComEstilo de Html com ("Minhas Tarefas", estilo, cartao).
Exibir com (site).`
    },
    html_galeria: {
        nome: 'HTML: Galeria de Imagens',
        codigo: `// Galeria de imagens com CSS Grid
Incluir Html.

A estilo é "
    .gallery { display: grid; grid-template-columns: repeat(3, 1fr); gap: 10px; width: 100%; max-width: 500px; padding: 10px; background: #020c1b; border-radius: 8px; }
    .img-box { width: 100%; aspect-ratio: 1; border-radius: 6px; overflow: hidden; border: 1px solid #00ff88; transition: transform 0.2s; }
    .img-box:hover { transform: scale(1.05); }
    .img-box img { width: 100%; height: 100%; object-fit: cover; }
    h3 { color: #00ff88; font-family: monospace; text-align: center; }
".

Um htmlLista está "".
Um contador está 0.

Repita 6 vezes:
    contador está contador + 1.
    Uma url é "https://picsum.photos/200?random=".
    Uma foto é CriarImagem de Html com (url, "Foto").
    Um box é CriarElementoComAtributos de Html com ("div", Atributo de Html com ("class", "img-box"), foto).
    htmlLista está htmlLista + box.
.

Um titulo é CriarElemento de Html com ("h3", "Fotos do Espaço").
Um grid é CriarElementoComAtributos de Html com ("div", Atributo de Html com ("class", "gallery"), htmlLista).
Um cartao é ListaElementos de Html com (titulo, grid).

Um site é CriarPaginaComEstilo de Html com ("Galeria", estilo, cartao).
Exibir com (site).`
    },
    html_tabela: {
        nome: 'HTML: Tabela de Preços',
        codigo: `// Construção nativa de tabelas com módulos do Verbo
Incluir Html.

A estilo é "
    body { background: #0f172a; color: white; font-family: sans-serif; padding: 2rem; }
    table { width: 100%; border-collapse: collapse; max-width: 600px; background: #1e293b; border-radius: 8px; overflow: hidden;}
    th { background: #334155; color: #00ff88; padding: 12px; text-align: left; }
    td { padding: 12px; border-bottom: 1px solid #334155; }
    tr:hover { background: #2a374a; }
".

Uma colunas é ["Plano", "Preço", "Benefício Principal"].
Uma linha1 é ["Starter", "$0", "Para iniciantes"].
Uma linha2 é ["Pro", "$15", "Ferramentas Avançadas"].
Uma linha3 é ["Enterprise", "Consulte", "Suporte 24/7"].

Uma matriz é [linha1, linha2, linha3].
Uma tabelaHtml é CriarTabela de Html com (colunas, matriz).

Um titulo é CriarElemento de Html com ("h2", "Planos de Assinatura").
Um bloco é ListaElementos de Html com (titulo, tabelaHtml).

Um site é CriarPaginaComEstilo de Html com ("Preços", estilo, bloco).
Exibir com (site).`
    },
    html_login: {
        nome: 'HTML: Tela de Login',
        codigo: `// Formulário interativo
Incluir Html.

A estilo é "
    .login-box { width: 320px; background: #0f172a; border: 1px solid #1e293b; border-radius: 12px; padding: 2rem; font-family: sans-serif; color: white; }
    .login-box h2 { text-align: center; margin-top: 0; color: #00ff88; }
    .grupo { display: flex; flex-direction: column; gap: 6px; margin-bottom: 1rem; }
    .grupo label { font-size: 13px; color: #94a3b8; }
    .grupo input { background: #1e293b; border: 1px solid #334155; padding: 10px; border-radius: 6px; color: white; outline: none; }
    form button { width: 100%; background: #00ff88; color: #0f172a; font-weight: bold; border: none; padding: 12px; border-radius: 6px; cursor: pointer; margin-top: 10px; }
".

O emailLabel é CriarElemento de Html com ("label", "Email").
O emailInput é CriarElementoComAtributos de Html com ("input", Atributo de Html com ("type", "email") + " placeholder='seu@email.com'", "").
O grpEmail é CriarElementoComAtributos de Html com ("div", Atributo de Html com ("class", "grupo"), emailLabel + emailInput).

A senhaLabel é CriarElemento de Html com ("label", "Senha").
A senhaInput é CriarElementoComAtributos de Html com ("input", Atributo de Html com ("type", "password") + " placeholder='********'", "").
O grpSenha é CriarElementoComAtributos de Html com ("div", Atributo de Html com ("class", "grupo"), senhaLabel + senhaInput).

O bot é CriarElemento de Html com ("button", "Acessar Portal").
A formu é CriarElemento de Html com ("form", grpEmail + grpSenha + bot).
O titu é CriarElemento de Html com ("h2", "Entrar").

O formParams é ListaElementos de Html com (titu, formu).
Um cartao é CriarElementoComAtributos de Html com ("div", Atributo de Html com ("class", "login-box"), formParams).

Um site é CriarPaginaComEstilo de Html com ("Acesso Seguro", estilo, cartao).
Exibir com (site).`
    },
    logica_fizzbuzz: {
        nome: 'Lógica: FizzBuzz',
        codigo: `// O clássico problema da entrevista de programação!
// Conta de 1 a 20: 
// - Múltiplos de 3: Fizz
// - Múltiplos de 5: Buzz
// - Ambos: FizzBuzz

Exibir com ("Iniciando FizzBuzz...").
Um numero está 0.

Repita 20 vezes:
    numero está numero + 1.
    
    // Como Verbo não tem o operador %, usamos matemática de inteiros
    // resto = numero - (numero / divisor * divisor)
    Um divTres está numero / 3.
    Um multTres está divTres * 3.
    Um restoTres está numero - multTres.
    
    Um divCinco está numero / 5.
    Um multCinco está divCinco * 5.
    Um restoCinco está numero - multCinco.
    
    Um escrito está 0.
    Se restoTres == 0 e restoCinco == 0, então:
        Exibir com ("FizzBuzz").
        escrito está 1.
    .
    
    Se escrito == 0, então:
        Se restoTres == 0, então:
            Exibir com ("Fizz").
            escrito está 1.
        .
    .
    
    Se escrito == 0, então:
        Se restoCinco == 0, então:
            Exibir com ("Buzz").
            escrito está 1.
        .
    .
    
    Se escrito == 0, então:
        Exibir com (numero).
    .
.

Exibir com ("Fim!").`
    },
    logica_fibonacci: {
        nome: 'Lógica: Fibonacci',
        codigo: `// Gera a sequência matemática de Fibonacci iterativamente
// Onde cada número é a soma dos dois anteriores

Um termoA está 0.
Um termoB está 1.

Exibir com ("Sequência de Fibonacci:").
Exibir com (termoA).
Exibir com (termoB).

Repita 8 vezes: // Gera os próximos 8 itens da série
    Um proximo está termoA + termoB.
    Exibir com (proximo).
    
    // Atualiza os ponteiros para a próxima iteração
    termoA está termoB.
    termoB está proximo.
.

Exibir com ("-- Concluído --").`
    },
    ola_mundo: {
        nome: 'Olá Mundo',
        codigo: `// Meu primeiro programa em Verbo
A saudacao é "Olá, Mundo!".
Exibir com (saudacao).`
    },
    variaveis: {
        nome: 'Variáveis',
        codigo: `// Constantes (imutáveis) e Variáveis (mutáveis)
A mensagem é "Bem-vindo ao Verbo!".
Um contador está 0.

Exibir com (mensagem).

Repita 5 vezes:
    contador está contador + 1.
    Exibir com (contador).
.`
    },
    funcoes: {
        nome: 'Funções',
        codigo: `// Declaração e chamada de funções
Para Saudar usando (nome: Texto):
    Exibir com ("Olá, ").
    Exibir com (nome).
.

Para Dobrar usando (x: Inteiro):
    Retorne x * 2.
.

Saudar com ("Brasil").

O resultado é Dobrar com (21).
Exibir com ("O dobro de 21 é:").
Exibir com (resultado).`
    },
    calculadora: {
        nome: 'Calculadora',
        codigo: `// Calculadora básica
Para Somar usando (x: Inteiro, y: Inteiro):
    Retorne x + y.
.

Para Multiplicar usando (x: Inteiro, y: Inteiro):
    Retorne x * y.
.

O resultado é Somar com (10, 20).
Exibir com ("Soma: ").
Exibir com (resultado).

O produto é Multiplicar com (5, 7).
Exibir com ("Produto: ").
Exibir com (produto).`
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
    Exibir com (preco).
.`
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
.

Exibir com ("Programa continua normalmente.").`
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
            value: EXEMPLOS.html_perfil.codigo,
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
            console.error("[Crom-Verbo API] Compilation Error:", data.erro);
        } else {
            statusIndicator.className = 'w-2.5 h-2.5 rounded-full status-ok transition-all';
            statusText.textContent = 'Sucesso';

            console.log("[Crom-Verbo API] Raw Output Received:\n", data.output);

            // Removing Verbo CLI decorative headers & footers for strict rendering
            let cleanOutput = data.output || '';
            cleanOutput = cleanOutput.replace(/🚀 Executando programa Verbo\.\.\.\r?\n(─+)\r?\n/g, '');
            cleanOutput = cleanOutput.replace(/\r?\n(─+)\r?\n✅ Programa finalizado com sucesso\.\r?\n?/g, '');
            cleanOutput = cleanOutput.trim();

            console.log("[Crom-Verbo API] Cleaned Output:\n", cleanOutput);

            consoleEl.textContent = cleanOutput || '(sem saída)';

            // HTML output detection
            if (cleanOutput && cleanOutput.startsWith('<')) {
                console.log("[Crom-Verbo UI] HTML Tag Detected! Switching to Web View.");
                document.getElementById('output-html').srcdoc = cleanOutput;
                switchTab('html');
            } else {
                console.log("[Crom-Verbo UI] Renderizing as standard plain text console.");
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
