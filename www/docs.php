<!DOCTYPE html>
<html lang="pt-BR" class="dark">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Documentação — Verbo</title>
    <meta name="description" content="Especificação completa da linguagem Verbo com mini playground e modo Ajudante.">
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link
        href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700;800&family=JetBrains+Mono:wght@400;500;600&display=swap"
        rel="stylesheet">
    <script src="https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"></script>
    <style type="text/tailwindcss">
        @theme {
        --color-verde: #009739; --color-verde-neon: #00ff88; --color-verde-dark: #006b28;
        --color-amarelo: #FFDF00; --color-azul: #012169; --color-azul-light: #1a3a8a;
        --color-azul-petroleo: #0a192f; --color-azul-noite: #020c1b; --color-branco-gelo: #f0f4f8;
        --font-sans: 'Inter', system-ui, sans-serif; --font-mono: 'JetBrains Mono', monospace;
      }
    </style>
    <style>
        html {
            scroll-behavior: smooth
        }

        ::selection {
            background: rgba(0, 255, 136, .3)
        }

        .glass-nav {
            background: rgba(2, 12, 27, .85);
            backdrop-filter: blur(20px);
            border-bottom: 1px solid rgba(0, 255, 136, .06)
        }

        .glass-card {
            background: rgba(10, 25, 47, .8);
            border: 1px solid rgba(0, 255, 136, .08);
            border-radius: 16px
        }

        .text-gradient-brasil {
            background: linear-gradient(135deg, #00ff88, #FFDF00, #009739);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
            background-clip: text
        }

        .syn-kw {
            color: #00ff88;
            font-weight: 600
        }

        .syn-tp {
            color: #FFDF00;
            font-weight: 600
        }

        .syn-str {
            color: #a8e6a3
        }

        .syn-cmt {
            color: #546e7a;
            font-style: italic
        }

        .syn-num {
            color: #ffd700
        }

        .syn-art {
            color: #80cbc4
        }

        .syn-op {
            color: #82b1ff
        }

        .syn-fn {
            color: #00c04b
        }

        .syn-id {
            color: #ccd6f6
        }

        .code-window {
            background: #020c1b;
            border: 1px solid rgba(0, 255, 136, .1);
            border-radius: 16px;
            overflow: hidden;
            position: relative
        }

        .code-window pre {
            padding: 20px 24px;
            font-family: 'JetBrains Mono', monospace;
            font-size: 13px;
            line-height: 1.8;
            overflow-x: auto;
            margin: 0
        }

        .code-window .copy-btn {
            position: absolute;
            top: 10px;
            right: 10px;
            padding: 4px 8px;
            border-radius: 8px;
            font-size: 11px;
            background: rgba(255, 255, 255, .05);
            color: #00ff88;
            cursor: pointer;
            opacity: 0;
            transition: opacity .2s;
            border: none
        }

        .code-window:hover .copy-btn {
            opacity: 1
        }

        .code-window .copy-btn:hover {
            background: rgba(0, 255, 136, .1)
        }

        .sidebar-link {
            display: block;
            padding: 6px 14px;
            font-size: 13px;
            border-left: 2px solid transparent;
            transition: all .2s;
            opacity: .45
        }

        .sidebar-link:hover {
            opacity: 1;
            color: #00ff88
        }

        .sidebar-link.active {
            opacity: 1;
            color: #00ff88;
            font-weight: 600;
            border-left-color: #00ff88;
            background: rgba(0, 255, 136, .03)
        }

        .section-label {
            font-family: 'JetBrains Mono', monospace;
            font-size: 10px;
            letter-spacing: .2em;
            text-transform: uppercase;
            opacity: .4;
            margin-bottom: 8px
        }

        .doc-table {
            width: 100%;
            font-size: 14px;
            border-collapse: collapse
        }

        .doc-table thead {
            background: rgba(0, 255, 136, .03)
        }

        .doc-table th {
            padding: 12px 16px;
            text-align: left;
            font-weight: 600;
            font-size: 12px;
            color: #00ff88
        }

        .doc-table td {
            padding: 10px 16px;
            border-top: 1px solid rgba(255, 255, 255, .04)
        }

        .doc-table tbody tr:hover {
            background: rgba(0, 255, 136, .02)
        }

        #scroll-indicator {
            position: fixed;
            top: 0;
            left: 0;
            height: 3px;
            z-index: 200;
            background: linear-gradient(90deg, #009739, #00ff88, #FFDF00);
            transition: width .15s linear
        }

        .sidebar-overlay {
            display: none;
            position: fixed;
            inset: 0;
            z-index: 40;
            background: rgba(0, 0, 0, .5)
        }

        .sidebar-overlay.open {
            display: block
        }

        /* Ajudante mode */
        .ajudante-tip {
            display: none;
            margin-top: 12px;
            padding: 14px 18px;
            border-radius: 14px;
            background: rgba(255, 223, 0, .04);
            border: 1px solid rgba(255, 223, 0, .12);
            font-size: 13px;
            color: rgba(255, 255, 255, .6);
            line-height: 1.7
        }

        .ajudante-tip strong {
            color: #FFDF00
        }

        body.ajudante-on .ajudante-tip {
            display: block
        }

        .ajudante-toggle {
            transition: all .3s
        }

        .ajudante-toggle.active {
            background: rgba(255, 223, 0, .15) !important;
            border-color: rgba(255, 223, 0, .3) !important;
            color: #FFDF00 !important
        }

        /* Mini playground */
        .mini-pg {
            border: 1px solid rgba(0, 255, 136, .1);
            border-radius: 14px;
            overflow: hidden;
            background: #020c1b;
            margin-top: 12px
        }

        .mini-pg textarea {
            width: 100%;
            background: transparent;
            color: #ccd6f6;
            font-family: 'JetBrains Mono', monospace;
            font-size: 13px;
            line-height: 1.7;
            padding: 14px 18px;
            border: none;
            resize: vertical;
            min-height: 80px;
            outline: none
        }

        .mini-pg-bar {
            display: flex;
            align-items: center;
            justify-content: space-between;
            padding: 8px 14px;
            border-top: 1px solid rgba(0, 255, 136, .06);
            background: rgba(0, 255, 136, .02)
        }

        .mini-pg-bar button {
            padding: 4px 14px;
            border-radius: 8px;
            font-size: 12px;
            font-weight: 600;
            cursor: pointer;
            transition: all .2s;
            border: none
        }

        .mini-pg-output {
            padding: 10px 18px;
            font-family: 'JetBrains Mono', monospace;
            font-size: 12px;
            color: #00ff88;
            border-top: 1px solid rgba(0, 255, 136, .06);
            white-space: pre-wrap;
            min-height: 28px;
            max-height: 120px;
            overflow-y: auto
        }
    </style>
</head>

<body class="bg-azul-petroleo text-branco-gelo font-sans">
    <div id="scroll-indicator" style="width:0%"></div>

    <!-- NAVBAR -->
    <nav class="fixed top-0 left-0 right-0 z-50 glass-nav">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 flex items-center justify-between h-14">
            <div class="flex items-center gap-3">
                <button onclick="toggleSidebar()" class="lg:hidden p-2 rounded-lg hover:bg-white/5 -ml-2">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M4 6h16M4 12h16M4 18h16" />
                    </svg>
                </button>
                <a href="index.php" class="flex items-center gap-2">
                    <div
                        class="w-7 h-7 rounded-lg bg-gradient-to-br from-verde to-verde-dark flex items-center justify-center text-white font-mono font-bold text-sm shadow-lg shadow-verde/20">
                        V</div>
                    <span class="font-bold text-sm">Documentação</span>
                </a>
            </div>
            <div class="flex items-center gap-3 text-sm">
                <!-- AJUDANTE TOGGLE -->
                <button id="ajudante-btn" onclick="toggleAjudante()"
                    class="ajudante-toggle flex items-center gap-1.5 px-3 py-1.5 rounded-xl border border-white/10 text-xs font-semibold text-white/40 hover:text-amarelo hover:border-amarelo/20"
                    title="Ativar explicações detalhadas para iniciantes">
                    🤝 Ajudante
                </button>
                <a href="playground.php" class="text-xs text-white/30 hover:text-verde-neon transition-colors">⚡
                    Playground</a>
                <button onclick="verboTheme.toggle()"
                    class="p-2 rounded-lg hover:bg-white/5 opacity-40 hover:opacity-100">
                    <svg data-theme-icon="light" class="w-4 h-4 hidden" fill="none" stroke="currentColor"
                        viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
                    </svg>
                    <svg data-theme-icon="dark" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
                    </svg>
                </button>
            </div>
        </div>
    </nav>

    <div id="sidebar-overlay" class="sidebar-overlay" onclick="toggleSidebar()"></div>

    <!-- LAYOUT -->
    <div class="flex pt-14">
        <!-- SIDEBAR -->
        <aside id="sidebar"
            class="fixed top-14 left-0 bottom-0 w-64 border-r border-verde-neon/5 overflow-y-auto p-5 bg-azul-petroleo z-45 transform -translate-x-full lg:translate-x-0 transition-transform duration-300">
            <div class="mb-5">
                <input id="docs-search" type="text" placeholder="🔍 Buscar..."
                    class="w-full bg-white/5 border border-white/8 rounded-xl px-3 py-2.5 text-sm focus:outline-none focus:border-verde-neon/20 placeholder:text-white/20"
                    oninput="filterDocs(this.value)">
            </div>
            <nav id="sidebar-nav">
                <p class="section-label text-verde-neon">Especificação</p>
                <a href="#visao-geral" class="sidebar-link">Visão Geral</a>
                <a href="#tipos" class="sidebar-link">Tipos de Dados</a>
                <a href="#variaveis" class="sidebar-link">Variáveis</a>
                <a href="#funcoes" class="sidebar-link">Funções</a>
                <a href="#controle" class="sidebar-link">Controle de Fluxo</a>
                <a href="#loops" class="sidebar-link">Loops</a>
                <a href="#operadores" class="sidebar-link">Operadores</a>
                <p class="section-label text-amarelo mt-6">Verbo 2.0</p>
                <a href="#entidades" class="sidebar-link">Entidades</a>
                <a href="#listas" class="sidebar-link">Listas</a>
                <a href="#concorrencia" class="sidebar-link">Concorrência</a>
                <a href="#erros" class="sidebar-link">Erros</a>
                <a href="#stdlib" class="sidebar-link">BibVerbo</a>
                <p class="section-label mt-6">Referência</p>
                <a href="#palavras" class="sidebar-link">Palavras Reservadas</a>
                <a href="#arquitetura" class="sidebar-link">Arquitetura</a>
            </nav>
        </aside>

        <!-- CONTENT -->
        <main class="flex-1 lg:ml-64 max-w-4xl mx-auto px-5 sm:px-8 py-10">
            <div class="mb-14">
                <h1 class="text-3xl sm:text-4xl font-black mb-3">Especificação da Linguagem <span
                        class="text-gradient-brasil">Verbo</span></h1>
                <p class="text-white/30 text-sm">Versão 2.0 — Em Desenvolvimento — Autor: Juan / Projeto Crom</p>
            </div>

            <!-- VISÃO GERAL -->
            <section id="visao-geral" class="mb-20 doc-section">
                <h2 class="text-2xl font-bold mb-5 flex items-center gap-2"><span
                        class="text-verde-neon opacity-40">#</span> Visão Geral</h2>
                <p class="text-white/60 leading-relaxed mb-5"><strong class="text-white">Verbo</strong> é uma linguagem
                    de programação transpilada que utiliza a gramática da norma culta do Português Brasileiro como
                    sintaxe. Transpilada para Go, garante performance nativa.</p>
                <div class="ajudante-tip">
                    🤝 <strong>O que é uma linguagem de programação?</strong> É um idioma que você usa para "conversar"
                    com o computador. Em vez de usar inglês como a maioria das linguagens, o Verbo usa português
                    brasileiro! <strong>Transpilar</strong> significa traduzir o código de um idioma para outro — no
                    caso, de Verbo para Go (outra linguagem muito rápida).
                </div>
                <div class="glass-card p-6 mt-5">
                    <h4 class="text-sm font-bold text-verde-neon mb-3">Princípios de Design</h4>
                    <ul class="space-y-2 text-sm text-white/50">
                        <li>✦ <strong class="text-white/70">Legibilidade máxima</strong> — Código deve parecer prosa
                            técnica</li>
                        <li>✦ <strong class="text-white/70">Gramática como semântica</strong> — Artigos, verbos e
                            preposições têm significado</li>
                        <li>✦ <strong class="text-white/70">Tipagem forte inferida</strong> — O sistema de tipos é
                            inferido pela gramática</li>
                        <li>✦ <strong class="text-white/70">Sem ambiguidade</strong> — Ordem SVO (Sujeito-Verbo-Objeto)
                            estrita</li>
                    </ul>
                </div>
            </section>

            <!-- TIPOS -->
            <section id="tipos" class="mb-20 doc-section">
                <h2 class="text-2xl font-bold mb-5 flex items-center gap-2"><span
                        class="text-verde-neon opacity-40">#</span> Tipos de Dados</h2>
                <div class="ajudante-tip">
                    🤝 <strong>O que são tipos de dados?</strong> Imagine que cada dado que você guarda no computador
                    tem uma "etiqueta". Um número inteiro (1, 2, 3) é diferente de um texto ("Olá"). Os tipos ajudam o
                    computador a saber como tratar cada informação. <strong>Texto</strong> = palavras e frases.
                    <strong>Inteiro</strong> = números sem vírgula. <strong>Decimal</strong> = números com vírgula.
                    <strong>Lógico</strong> = Verdadeiro ou Falso (sim/não). <strong>Lista</strong> = uma coleção de
                    coisas.
                </div>
                <div class="overflow-x-auto rounded-2xl border border-verde-neon/8">
                    <table class="doc-table">
                        <thead>
                            <tr>
                                <th>Tipo</th>
                                <th>Descrição</th>
                                <th>Exemplo</th>
                            </tr>
                        </thead>
                        <tbody class="text-white/50">
                            <tr>
                                <td class="font-mono"><span class="syn-tp">Texto</span></td>
                                <td>Cadeia de caracteres</td>
                                <td class="font-mono"><span class="syn-str">"Olá"</span></td>
                            </tr>
                            <tr>
                                <td class="font-mono"><span class="syn-tp">Inteiro</span></td>
                                <td>Número inteiro</td>
                                <td class="font-mono"><span class="syn-num">42</span></td>
                            </tr>
                            <tr>
                                <td class="font-mono"><span class="syn-tp">Decimal</span></td>
                                <td>Ponto flutuante</td>
                                <td class="font-mono"><span class="syn-num">3.14</span></td>
                            </tr>
                            <tr>
                                <td class="font-mono"><span class="syn-tp">Lógico</span></td>
                                <td>Verdadeiro/Falso</td>
                                <td class="font-mono"><span class="syn-kw">Verdadeiro</span></td>
                            </tr>
                            <tr>
                                <td class="font-mono"><span class="syn-tp">Lista</span></td>
                                <td>Coleção ordenada</td>
                                <td class="font-mono">[<span class="syn-num">1</span>, <span class="syn-num">2</span>,
                                    <span class="syn-num">3</span>]
                                </td>
                            </tr>
                            <tr>
                                <td class="font-mono"><span class="syn-tp">Nulo</span></td>
                                <td>Ausência de valor</td>
                                <td class="font-mono"><span class="syn-kw">Nulo</span></td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </section>

            <!-- VARIÁVEIS -->
            <section id="variaveis" class="mb-20 doc-section">
                <h2 class="text-2xl font-bold mb-5 flex items-center gap-2"><span
                        class="text-verde-neon opacity-40">#</span> Variáveis e Constantes</h2>
                <div class="ajudante-tip">
                    🤝 <strong>O que é uma variável?</strong> Pense numa caixa etiquetada onde você guarda um valor. O
                    nome da caixa é o nome da variável, e dentro dela está o valor. Em Verbo, usamos artigos do
                    português: <strong>"O"/"A"</strong> criam caixas que <em>nunca mudam</em> (constantes), e
                    <strong>"Um"/"Uma"</strong> criam caixas que <em>podem mudar</em> (variáveis). <strong>"é"</strong>
                    define algo permanente (como "A água é molhada") e <strong>"está"</strong> algo temporário (como "O
                    tempo está chuvoso").
                </div>
                <h3 class="text-lg font-bold mt-8 mb-3 text-amarelo">Constantes (Artigo Definido)</h3>
                <div class="code-window">
                    <pre><span class="syn-art">O</span> <span class="syn-id">limite</span> <span class="syn-op">é</span> <span class="syn-num">100</span>.
<span class="syn-art">A</span> <span class="syn-id">mensagem</span> <span class="syn-op">é</span> <span class="syn-str">"Olá, Mundo!"</span>.</pre>
                    <button onclick="copiarBloco(this)" class="copy-btn">📋</button>
                </div>
                <!-- MINI PLAYGROUND -->
                <div class="mini-pg mt-3">
                    <textarea id="mp-variaveis" spellcheck="false">A saudacao é "Olá, eu sou uma constante!".
Exibir com (saudacao).</textarea>
                    <div class="mini-pg-bar">
                        <span class="text-[10px] text-white/20 font-mono">Mini Playground</span>
                        <button onclick="runMini('mp-variaveis','mp-variaveis-out')"
                            class="bg-verde text-white hover:bg-verde-neon hover:text-azul-noite">▶ Testar</button>
                    </div>
                    <div id="mp-variaveis-out" class="mini-pg-output"></div>
                </div>
                <h3 class="text-lg font-bold mt-8 mb-3 text-amarelo">Variáveis (Artigo Indefinido)</h3>
                <div class="code-window">
                    <pre><span class="syn-art">Um</span> <span class="syn-id">contador</span> <span class="syn-op">está</span> <span class="syn-num">0</span>.
<span class="syn-art">Uma</span> <span class="syn-id">taxa</span> <span class="syn-op">está</span> <span class="syn-num">0.15</span>.</pre>
                    <button onclick="copiarBloco(this)" class="copy-btn">📋</button>
                </div>
                <div class="glass-card p-5 mt-6">
                    <h4 class="text-sm font-bold text-verde-neon mb-2">Semântica de Estado</h4>
                    <ul class="text-sm text-white/50 space-y-1">
                        <li>• <code class="font-mono text-amarelo">é</code> → Atribuição estática (natureza) — imutável
                        </li>
                        <li>• <code class="font-mono text-amarelo">está</code> → Atribuição de estado (temporário) —
                            mutável</li>
                    </ul>
                </div>
            </section>

            <!-- FUNÇÕES -->
            <section id="funcoes" class="mb-20 doc-section">
                <h2 class="text-2xl font-bold mb-5 flex items-center gap-2"><span
                        class="text-verde-neon opacity-40">#</span> Funções</h2>
                <div class="ajudante-tip">
                    🤝 <strong>O que é uma função?</strong> É como uma receita de bolo. Você dá um nome para ela
                    ("Calcular"), define os ingredientes ("valor: Inteiro") e descreve o passo a passo. Depois, pode
                    "chamar" essa receita quantas vezes quiser com ingredientes diferentes. A palavra
                    <strong>"Para"</strong> inicia a receita, <strong>"usando"</strong> define os ingredientes, e
                    <strong>"Retorne"</strong> devolve o resultado final.
                </div>
                <div class="code-window">
                    <pre><span class="syn-kw">Para</span> <span class="syn-id">Calcular</span> <span class="syn-op">usando</span> (<span class="syn-id">valor</span>: <span class="syn-tp">Inteiro</span>):
    <span class="syn-kw">Retorne</span> <span class="syn-id">valor</span> + <span class="syn-num">10</span>.

<span class="syn-art">O</span> <span class="syn-id">resultado</span> <span class="syn-op">é</span> <span class="syn-id">Calcular</span> <span class="syn-op">com</span> (<span class="syn-num">5</span>).
<span class="syn-fn">Exibir</span> <span class="syn-op">com</span> (<span class="syn-id">resultado</span>).</pre>
                    <button onclick="copiarBloco(this)" class="copy-btn">📋</button>
                </div>
                <div class="mini-pg mt-3">
                    <textarea id="mp-funcoes" spellcheck="false">Para Dobrar usando (n: Inteiro):
    Retorne n * 2.

O resultado é Dobrar com (21).
Exibir com (resultado).</textarea>
                    <div class="mini-pg-bar">
                        <span class="text-[10px] text-white/20 font-mono">Mini Playground</span>
                        <button onclick="runMini('mp-funcoes','mp-funcoes-out')"
                            class="bg-verde text-white hover:bg-verde-neon hover:text-azul-noite">▶ Testar</button>
                    </div>
                    <div id="mp-funcoes-out" class="mini-pg-output"></div>
                </div>
            </section>

            <!-- CONTROLE DE FLUXO -->
            <section id="controle" class="mb-20 doc-section">
                <h2 class="text-2xl font-bold mb-5 flex items-center gap-2"><span
                        class="text-verde-neon opacity-40">#</span> Controle de Fluxo</h2>
                <div class="ajudante-tip">
                    🤝 <strong>O que é controle de fluxo?</strong> É como uma placa de trânsito no código.
                    <strong>"Se"</strong> é uma pergunta: "Se isso for verdade, faça isso. <strong>Senão</strong>, faça
                    aquilo." Exemplo do dia a dia: "Se estiver chovendo, leve guarda-chuva. Senão, leve óculos de sol."
                    O computador segue a mesma lógica!
                </div>
                <div class="code-window">
                    <pre><span class="syn-kw">Se</span> <span class="syn-art">a</span> <span class="syn-id">idade</span> <span class="syn-op">for menor que</span> <span class="syn-num">18</span>, <span class="syn-op">então</span>:
    <span class="syn-fn">Exibir</span> <span class="syn-op">com</span> (<span class="syn-str">"Menor de idade"</span>).
<span class="syn-kw">Senão</span>:
    <span class="syn-fn">Exibir</span> <span class="syn-op">com</span> (<span class="syn-str">"Maior de idade"</span>).</pre>
                    <button onclick="copiarBloco(this)" class="copy-btn">📋</button>
                </div>
                <div class="mini-pg mt-3">
                    <textarea id="mp-controle" spellcheck="false">A idade é 25.
Se a idade for menor que 18, então:
    Exibir com ("Menor de idade").
Senão:
    Exibir com ("Maior de idade").</textarea>
                    <div class="mini-pg-bar">
                        <span class="text-[10px] text-white/20 font-mono">Mini Playground</span>
                        <button onclick="runMini('mp-controle','mp-controle-out')"
                            class="bg-verde text-white hover:bg-verde-neon hover:text-azul-noite">▶ Testar</button>
                    </div>
                    <div id="mp-controle-out" class="mini-pg-output"></div>
                </div>
            </section>

            <!-- LOOPS -->
            <section id="loops" class="mb-20 doc-section">
                <h2 class="text-2xl font-bold mb-5 flex items-center gap-2"><span
                        class="text-verde-neon opacity-40">#</span> Loops</h2>
                <div class="ajudante-tip">
                    🤝 <strong>O que é um loop?</strong> É quando você quer repetir algo várias vezes sem ter que
                    escrever a mesma coisa toda hora. <strong>"Repita 10 vezes"</strong> = faz algo 10 vezes.
                    <strong>"Enquanto"</strong> = continua fazendo enquanto uma condição for verdadeira (como "enquanto
                    tiver fome, coma"). <strong>"Para cada"</strong> = faz algo para cada item de uma lista.
                </div>
                <div class="code-window">
                    <pre><span class="syn-kw">Repita</span> <span class="syn-num">5</span> <span class="syn-op">vezes</span>:
    <span class="syn-fn">Exibir</span> <span class="syn-op">com</span> (<span class="syn-str">"Repetindo!"</span>).</pre>
                    <button onclick="copiarBloco(this)" class="copy-btn">📋</button>
                </div>
                <div class="mini-pg mt-3">
                    <textarea id="mp-loops" spellcheck="false">Repita 5 vezes:
    Exibir com ("Vez número: iteração").</textarea>
                    <div class="mini-pg-bar">
                        <span class="text-[10px] text-white/20 font-mono">Mini Playground</span>
                        <button onclick="runMini('mp-loops','mp-loops-out')"
                            class="bg-verde text-white hover:bg-verde-neon hover:text-azul-noite">▶ Testar</button>
                    </div>
                    <div id="mp-loops-out" class="mini-pg-output"></div>
                </div>
            </section>

            <!-- OPERADORES -->
            <section id="operadores" class="mb-20 doc-section">
                <h2 class="text-2xl font-bold mb-5 flex items-center gap-2"><span
                        class="text-verde-neon opacity-40">#</span> Operadores</h2>
                <div class="ajudante-tip">
                    🤝 <strong>O que são operadores?</strong> São os símbolos de matemática que você já conhece!
                    <strong>+</strong> soma, <strong>-</strong> subtrai, <strong>*</strong> multiplica,
                    <strong>/</strong> divide. Simples assim. Para juntar dois textos, usamos <strong>"e"</strong> —
                    como na frase "João <em>e</em> Maria".
                </div>
                <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                    <div class="glass-card p-5">
                        <h4 class="text-sm font-bold text-verde-neon mb-3">Aritméticos</h4>
                        <div class="space-y-1 text-sm font-mono text-white/50">
                            <div>+ Adição</div>
                            <div>- Subtração</div>
                            <div>* Multiplicação</div>
                            <div>/ Divisão</div>
                        </div>
                    </div>
                    <div class="glass-card p-5">
                        <h4 class="text-sm font-bold text-amarelo mb-3">Comparação em Português</h4>
                        <div class="space-y-1 text-sm font-mono text-white/50">
                            <div><span class="syn-op">menor que</span> → &lt;</div>
                            <div><span class="syn-op">maior que</span> → &gt;</div>
                            <div><span class="syn-op">igual</span> → ==</div>
                        </div>
                    </div>
                </div>
            </section>

            <!-- ENTIDADES v2 -->
            <section id="entidades" class="mb-20 doc-section">
                <h2 class="text-2xl font-bold mb-5 flex items-center gap-2"><span
                        class="text-amarelo opacity-40">#</span> Entidades <span
                        class="text-[10px] font-mono bg-amarelo/10 text-amarelo px-2 py-0.5 rounded-full">v2.0</span>
                </h2>
                <div class="ajudante-tip">
                    🤝 <strong>O que é uma Entidade?</strong> Imagine uma ficha cadastral. Uma <strong>Entidade
                        "Pessoa"</strong> é um modelo de ficha com campos como "Nome" e "Idade". Quando você cria "um
                    novo Pessoa", está preenchendo uma ficha real com dados. É como a diferença entre o <em>formulário
                        em branco</em> (a Entidade) e o <em>formulário preenchido</em> (a instância).
                </div>
                <div class="code-window">
                    <pre><span class="syn-art">A</span> <span class="syn-kw">Entidade</span> <span class="syn-tp">Pessoa</span> <span class="syn-op">contendo</span> (<span class="syn-id">Nome</span>: <span class="syn-tp">Texto</span>, <span class="syn-id">Idade</span>: <span class="syn-tp">Inteiro</span>).

<span class="syn-art">Um</span> <span class="syn-id">cliente</span> <span class="syn-op">é</span> <span class="syn-tp">Pessoa</span> <span class="syn-op">com</span> (<span class="syn-str">"Ada"</span>, <span class="syn-num">36</span>).
<span class="syn-fn">Exibir</span> <span class="syn-op">com</span> (<span class="syn-id">cliente</span>).</pre>
                    <button onclick="copiarBloco(this)" class="copy-btn">📋</button>
                </div>
                <div class="mini-pg mt-3">
                    <textarea id="mp-entidade" spellcheck="false">A Entidade Produto contendo (Nome: Texto, Preco: Inteiro).
Um item é Produto com ("Café", 15).
Exibir com (item).</textarea>
                    <div class="mini-pg-bar">
                        <span class="text-[10px] text-white/20 font-mono">Mini Playground</span>
                        <button onclick="runMini('mp-entidade','mp-entidade-out')"
                            class="bg-verde text-white hover:bg-verde-neon hover:text-azul-noite">▶ Testar</button>
                    </div>
                    <div id="mp-entidade-out" class="mini-pg-output"></div>
                </div>
            </section>

            <!-- LISTAS v2 -->
            <section id="listas" class="mb-20 doc-section">
                <h2 class="text-2xl font-bold mb-5 flex items-center gap-2"><span
                        class="text-amarelo opacity-40">#</span> Listas <span
                        class="text-[10px] font-mono bg-amarelo/10 text-amarelo px-2 py-0.5 rounded-full">v2.0</span>
                </h2>
                <div class="ajudante-tip">
                    🤝 <strong>O que é uma Lista?</strong> É como uma fila de supermercado. Cada pessoa (ou item) tem
                    uma posição. A posição começa em <strong>0</strong> (sim, programadores contam a partir do zero!).
                    <code>[10, 20, 30]</code> é uma lista com 3 números. O primeiro item é <code>lista[0]</code> = 10.
                </div>
                <div class="code-window">
                    <pre><span class="syn-art">Uma</span> <span class="syn-id">precos</span> <span class="syn-op">é</span> [<span class="syn-num">10</span>, <span class="syn-num">20</span>, <span class="syn-num">30</span>].
<span class="syn-art">O</span> <span class="syn-id">primeiro</span> <span class="syn-op">é</span> <span class="syn-id">precos</span>[<span class="syn-num">0</span>].
<span class="syn-fn">Exibir</span> <span class="syn-op">com</span> (<span class="syn-id">primeiro</span>).</pre>
                    <button onclick="copiarBloco(this)" class="copy-btn">📋</button>
                </div>
            </section>

            <!-- CONCORRÊNCIA v2 -->
            <section id="concorrencia" class="mb-20 doc-section">
                <h2 class="text-2xl font-bold mb-5 flex items-center gap-2"><span
                        class="text-amarelo opacity-40">#</span> Concorrência <span
                        class="text-[10px] font-mono bg-amarelo/10 text-amarelo px-2 py-0.5 rounded-full">v2.0</span>
                </h2>
                <div class="ajudante-tip">
                    🤝 <strong>O que é Concorrência?</strong> Imagine que você tem 3 tarefas: lavar roupa, cozinhar e
                    limpar a casa. Normalmente faria uma de cada vez. Com concorrência, você inicia as três ao mesmo
                    tempo! <strong>"Simultaneamente"</strong> faz exatamente isso — cada instrução dentro roda em
                    paralelo, como ter 3 funcionários trabalhando ao mesmo tempo.
                </div>
                <div class="code-window">
                    <pre><span class="syn-kw">Simultaneamente</span>:
    <span class="syn-fn">Exibir</span> <span class="syn-op">com</span> (<span class="syn-str">"Tarefa 1"</span>).
    <span class="syn-fn">Exibir</span> <span class="syn-op">com</span> (<span class="syn-str">"Tarefa 2"</span>).</pre>
                    <button onclick="copiarBloco(this)" class="copy-btn">📋</button>
                </div>
                <div class="mini-pg mt-3">
                    <textarea id="mp-conc" spellcheck="false">Simultaneamente:
    Exibir com ("Processando dados...").
    Exibir com ("Enviando email...").
    Exibir com ("Gerando relatório...").

Exibir com ("Tudo pronto!").</textarea>
                    <div class="mini-pg-bar">
                        <span class="text-[10px] text-white/20 font-mono">Mini Playground</span>
                        <button onclick="runMini('mp-conc','mp-conc-out')"
                            class="bg-verde text-white hover:bg-verde-neon hover:text-azul-noite">▶ Testar</button>
                    </div>
                    <div id="mp-conc-out" class="mini-pg-output"></div>
                </div>
            </section>

            <!-- ERROS v2 -->
            <section id="erros" class="mb-20 doc-section">
                <h2 class="text-2xl font-bold mb-5 flex items-center gap-2"><span
                        class="text-amarelo opacity-40">#</span> Tratamento de Erros <span
                        class="text-[10px] font-mono bg-amarelo/10 text-amarelo px-2 py-0.5 rounded-full">v2.0</span>
                </h2>
                <div class="ajudante-tip">
                    🤝 <strong>O que é tratamento de erros?</strong> Imagine que você está tentando abrir uma porta. Se
                    ela estiver trancada, você não quer que seu programa "quebre" — você quer lidar com a situação.
                    <strong>"Tente"</strong> = "tente fazer isso". <strong>"Capture"</strong> = "se der errado, faça
                    isso em vez". <strong>"Sinalize"</strong> = "avise que algo deu errado!". É como ter um plano B.
                </div>
                <div class="code-window">
                    <pre><span class="syn-kw">Tente</span>:
    <span class="syn-kw">Sinalize</span> <span class="syn-op">com</span> (<span class="syn-str">"Saldo insuficiente!"</span>).
<span class="syn-kw">Capture</span> <span class="syn-id">erro</span>:
    <span class="syn-fn">Exibir</span> <span class="syn-op">com</span> (<span class="syn-id">erro</span>).</pre>
                    <button onclick="copiarBloco(this)" class="copy-btn">📋</button>
                </div>
                <div class="mini-pg mt-3">
                    <textarea id="mp-erros" spellcheck="false">Tente:
    Exibir com ("Tentando...").
    Sinalize com ("Ops! Algo falhou.").
Capture erro:
    Exibir com ("Erro capturado:").
    Exibir com (erro).

Exibir com ("Programa continua!").</textarea>
                    <div class="mini-pg-bar">
                        <span class="text-[10px] text-white/20 font-mono">Mini Playground</span>
                        <button onclick="runMini('mp-erros','mp-erros-out')"
                            class="bg-verde text-white hover:bg-verde-neon hover:text-azul-noite">▶ Testar</button>
                    </div>
                    <div id="mp-erros-out" class="mini-pg-output"></div>
                </div>
            </section>

            <!-- STDLIB -->
            <section id="stdlib" class="mb-20 doc-section">
                <h2 class="text-2xl font-bold mb-5 flex items-center gap-2"><span
                        class="text-amarelo opacity-40">#</span> BibVerbo (Biblioteca Padrão)</h2>
                <div class="ajudante-tip">
                    🤝 <strong>O que é uma biblioteca?</strong> São "poderes extras" prontos para usar. Em vez de
                    inventar tudo do zero, você <strong>inclui</strong> uma biblioteca e ganha acesso a funções que
                    alguém já criou. <strong>Matematica</strong> = funções de cálculo. <strong>Texto</strong> =
                    manipulação de palavras. <strong>Html</strong> = criar páginas web!
                </div>
                <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 mb-5">
                    <div class="glass-card p-5">
                        <h4 class="text-sm font-bold text-verde-neon mb-2">Matematica</h4>
                        <ul class="text-xs text-white/40 space-y-1 font-mono">
                            <li>Absoluto(x)</li>
                            <li>Teto(x)</li>
                            <li>Piso(x)</li>
                            <li>Maximo(a,b)</li>
                            <li>Raiz(x)</li>
                        </ul>
                    </div>
                    <div class="glass-card p-5">
                        <h4 class="text-sm font-bold text-amarelo mb-2">Texto</h4>
                        <ul class="text-xs text-white/40 space-y-1 font-mono">
                            <li>Maiusculas(t)</li>
                            <li>Minusculas(t)</li>
                            <li>Contem(t,s)</li>
                            <li>Tamanho(t)</li>
                        </ul>
                    </div>
                    <div class="glass-card p-5">
                        <h4 class="text-sm font-bold text-[#64b5f6] mb-2">Html</h4>
                        <ul class="text-xs text-white/40 space-y-1 font-mono">
                            <li>CriarElemento(t,c)</li>
                            <li>CriarPagina(t,c)</li>
                            <li>CriarLink(u,t)</li>
                            <li>CriarLista(i...)</li>
                        </ul>
                    </div>
                </div>
            </section>

            <!-- PALAVRAS RESERVADAS -->
            <section id="palavras" class="mb-20 doc-section">
                <h2 class="text-2xl font-bold mb-5 flex items-center gap-2"><span
                        class="text-verde-neon opacity-40">#</span> Palavras Reservadas</h2>
                <div class="overflow-x-auto rounded-2xl border border-verde-neon/8">
                    <table class="doc-table">
                        <thead>
                            <tr>
                                <th>Palavra</th>
                                <th>Função</th>
                            </tr>
                        </thead>
                        <tbody class="text-white/50">
                            <tr>
                                <td class="font-mono"><span class="syn-art">O/A</span></td>
                                <td>Artigo definido (constante)</td>
                            </tr>
                            <tr>
                                <td class="font-mono"><span class="syn-art">Um/Uma</span></td>
                                <td>Artigo indefinido (variável)</td>
                            </tr>
                            <tr>
                                <td class="font-mono"><span class="syn-op">é</span></td>
                                <td>Atribuição permanente</td>
                            </tr>
                            <tr>
                                <td class="font-mono"><span class="syn-op">está</span></td>
                                <td>Atribuição temporária</td>
                            </tr>
                            <tr>
                                <td class="font-mono"><span class="syn-kw">Para</span></td>
                                <td>Declaração de função</td>
                            </tr>
                            <tr>
                                <td class="font-mono"><span class="syn-kw">Se / Senão</span></td>
                                <td>Condicional</td>
                            </tr>
                            <tr>
                                <td class="font-mono"><span class="syn-kw">Repita / Enquanto</span></td>
                                <td>Loops</td>
                            </tr>
                            <tr>
                                <td class="font-mono"><span class="syn-kw">Simultaneamente</span></td>
                                <td>Concorrência</td>
                            </tr>
                            <tr>
                                <td class="font-mono"><span class="syn-kw">Tente / Capture</span></td>
                                <td>Tratamento de erros</td>
                            </tr>
                            <tr>
                                <td class="font-mono"><span class="syn-fn">Exibir</span></td>
                                <td>Saída padrão</td>
                            </tr>
                            <tr>
                                <td class="font-mono"><span class="syn-kw">Retorne</span></td>
                                <td>Retorno de função</td>
                            </tr>
                            <tr>
                                <td class="font-mono"><span class="syn-fn">Incluir</span></td>
                                <td>Importar pacote</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </section>

            <!-- ARQUITETURA -->
            <section id="arquitetura" class="mb-20 doc-section">
                <h2 class="text-2xl font-bold mb-5 flex items-center gap-2"><span
                        class="text-verde-neon opacity-40">#</span> Arquitetura do Compilador</h2>
                <div class="ajudante-tip">
                    🤝 <strong>Como o Verbo funciona por dentro?</strong> Seu código passa por 4 etapas: (1) O
                    <strong>Lexer</strong> lê seu texto e separa em "palavras" (tokens). (2) O <strong>Parser</strong>
                    organiza essas palavras numa árvore lógica (AST). (3) O <strong>Transpiler</strong> converte essa
                    árvore em código Go. (4) O Go compila num <strong>binário</strong> super rápido. É como traduzir um
                    livro: primeiro você entende as palavras, depois as frases, depois reescreve em outro idioma!
                </div>
                <div class="glass-card p-6 mb-6">
                    <div class="flex items-center justify-center gap-2 sm:gap-3 flex-wrap font-mono text-sm">
                        <span
                            class="px-3 py-2 rounded-xl bg-verde/10 text-verde-neon border border-verde/20">.vrb</span>
                        <span class="text-white/15">→</span>
                        <span
                            class="px-3 py-2 rounded-xl bg-amarelo/10 text-amarelo border border-amarelo/10">Lexer</span>
                        <span class="text-white/15">→</span>
                        <span
                            class="px-3 py-2 rounded-xl bg-amarelo/10 text-amarelo border border-amarelo/10">Parser</span>
                        <span class="text-white/15">→</span>
                        <span
                            class="px-3 py-2 rounded-xl bg-amarelo/10 text-amarelo border border-amarelo/10">AST</span>
                        <span class="text-white/15">→</span>
                        <span
                            class="px-3 py-2 rounded-xl bg-verde/10 text-verde-neon border border-verde/20">Transpiler</span>
                        <span class="text-white/15">→</span>
                        <span class="px-3 py-2 rounded-xl bg-verde/10 text-verde-neon border border-verde/20">.go</span>
                        <span class="text-white/15">→</span>
                        <span
                            class="px-3 py-2 rounded-xl bg-white/5 text-[#64b5f6] border border-white/10">Binário</span>
                    </div>
                </div>
                <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                    <div class="glass-card p-5">
                        <h4 class="text-sm font-bold text-verde-neon mb-1">Lexer <span
                                class="text-white/15 font-normal">pkg/lexer/</span></h4>
                        <p class="text-xs text-white/40">Scanner UTF-8, tokenização de keywords, literais, operadores.
                        </p>
                    </div>
                    <div class="glass-card p-5">
                        <h4 class="text-sm font-bold text-amarelo mb-1">Parser <span
                                class="text-white/15 font-normal">pkg/parser/</span></h4>
                        <p class="text-xs text-white/40">Recursive Descent com precedência de expressões.</p>
                    </div>
                    <div class="glass-card p-5">
                        <h4 class="text-sm font-bold text-verde-neon mb-1">AST <span
                                class="text-white/15 font-normal">pkg/ast/</span></h4>
                        <p class="text-xs text-white/40">Nós tipados: Programa, Variáveis, Funções, Se, Repita…</p>
                    </div>
                    <div class="glass-card p-5">
                        <h4 class="text-sm font-bold text-amarelo mb-1">Transpiler <span
                                class="text-white/15 font-normal">pkg/transpiler/</span></h4>
                        <p class="text-xs text-white/40">Visitor pattern para gerar código Go compilável.</p>
                    </div>
                </div>
            </section>

            <div class="text-center mt-16 opacity-20">
                <a href="#" class="hover:text-verde-neon transition-colors text-sm">↑ Voltar ao topo</a>
            </div>
        </main>
    </div>

    <script src="js/theme.js"></script>
    <script>
        function copiarBloco(btn) { const p = btn.previousElementSibling; navigator.clipboard.writeText(p.textContent); btn.textContent = '✓'; setTimeout(() => btn.textContent = '📋', 1500) }

        // Scroll indicator
        window.addEventListener('scroll', () => { const p = (window.scrollY / (document.body.scrollHeight - window.innerHeight)) * 100; document.getElementById('scroll-indicator').style.width = p + '%' });

        // Scroll spy
        const sections = document.querySelectorAll('.doc-section');
        const navLinks = document.querySelectorAll('.sidebar-link');
        const spy = new IntersectionObserver(e => { e.forEach(en => { if (en.isIntersecting) navLinks.forEach(l => l.classList.toggle('active', l.getAttribute('href') === '#' + en.target.id)) }) }, { rootMargin: '-20% 0px -70% 0px' });
        sections.forEach(s => spy.observe(s));

        // Search
        function filterDocs(q) { const s = q.toLowerCase(); sections.forEach(sec => { sec.style.display = sec.textContent.toLowerCase().includes(s) ? '' : 'none' }) }

        // Mobile sidebar
        function toggleSidebar() { document.getElementById('sidebar').classList.toggle('-translate-x-full'); document.getElementById('sidebar-overlay').classList.toggle('open') }
        navLinks.forEach(l => l.addEventListener('click', () => { if (window.innerWidth < 1024) toggleSidebar() }));

        // AJUDANTE MODE
        function toggleAjudante() {
            document.body.classList.toggle('ajudante-on');
            document.getElementById('ajudante-btn').classList.toggle('active');
            localStorage.setItem('ajudante', document.body.classList.contains('ajudante-on') ? '1' : '0');
        }
        if (localStorage.getItem('ajudante') === '1') { document.body.classList.add('ajudante-on'); document.getElementById('ajudante-btn').classList.add('active') }

        // MINI PLAYGROUND
        async function runMini(textareaId, outputId) {
            const code = document.getElementById(textareaId).value;
            const out = document.getElementById(outputId);
            out.textContent = '⏳ Executando...';
            out.style.color = '#FFDF00';
            try {
                const r = await fetch('api/compilar.php', { method: 'POST', headers: { 'Content-Type': 'application/json' }, body: JSON.stringify({ codigo: code }) });
                const d = await r.json();
                if (d.erro) { out.textContent = '❌ ' + d.erro; out.style.color = '#ff5555' }
                else { out.textContent = d.output || '(sem saída)'; out.style.color = '#00ff88' }
            } catch (e) { out.textContent = '⚠ Servidor não disponível. Use o Playground completo.'; out.style.color = '#FFDF00' }
        }
    </script>
</body>

</html>