<!DOCTYPE html>
<html lang="pt-BR" class="dark">

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Verbo — Programe com a alma de um brasileiro</title>
  <meta name="description"
    content="Verbo é a primeira linguagem de programação de alto desempenho baseada na gramática brasileira. Escreva como prosa; execute com a performance do Go.">
  <meta property="og:title" content="Verbo — Linguagem de Programação em Português">
  <meta property="og:description" content="Performance do Go com a clareza da língua portuguesa.">
  <meta property="og:type" content="website">
  <meta name="twitter:card" content="summary_large_image">

  <!-- Fonts -->
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
  <link
    href="https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700;800;900&family=JetBrains+Mono:wght@400;500;600;700&display=swap"
    rel="stylesheet">

  <!-- Tailwind v4 Browser CDN -->
  <script src="https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"></script>
  <style type="text/tailwindcss">
    @theme {
        --color-verde: #009739;
        --color-verde-light: #00c04b;
        --color-verde-neon: #00ff88;
        --color-verde-dark: #006b28;
        --color-amarelo: #FFDF00;
        --color-amarelo-light: #ffe94d;
        --color-amarelo-dark: #ccb200;
        --color-azul: #012169;
        --color-azul-light: #1a3a8a;
        --color-azul-petroleo: #0a192f;
        --color-azul-noite: #020c1b;
        --color-branco-gelo: #f0f4f8;
        --font-sans: 'Inter', system-ui, -apple-system, sans-serif;
        --font-mono: 'JetBrains Mono', 'Fira Code', monospace;
      }
    </style>

  <style>
    /* ===== ANIMATIONS ===== */
    @keyframes fade-in {
      from {
        opacity: 0;
      }

      to {
        opacity: 1;
      }
    }

    @keyframes slide-up {
      from {
        opacity: 0;
        transform: translateY(40px);
      }

      to {
        opacity: 1;
        transform: translateY(0);
      }
    }

    @keyframes glow-pulse {

      0%,
      100% {
        box-shadow: 0 0 20px rgba(0, 255, 136, 0.15);
      }

      50% {
        box-shadow: 0 0 40px rgba(0, 255, 136, 0.35);
      }
    }

    @keyframes float {

      0%,
      100% {
        transform: translateY(0);
      }

      50% {
        transform: translateY(-12px);
      }
    }

    @keyframes gradient-shift {
      0% {
        background-position: 0% 50%;
      }

      50% {
        background-position: 100% 50%;
      }

      100% {
        background-position: 0% 50%;
      }
    }

    .animate-fade-in {
      animation: fade-in 0.7s ease-out both;
    }

    .animate-slide-up {
      animation: slide-up 0.7s ease-out both;
    }

    .animate-glow {
      animation: glow-pulse 3s ease-in-out infinite;
    }

    .animate-float {
      animation: float 8s ease-in-out infinite;
    }

    .delay-100 {
      animation-delay: 0.1s;
    }

    .delay-200 {
      animation-delay: 0.2s;
    }

    .delay-300 {
      animation-delay: 0.3s;
    }

    .delay-400 {
      animation-delay: 0.4s;
    }

    .delay-500 {
      animation-delay: 0.5s;
    }

    /* ===== GRADIENT TEXT ===== */
    .text-gradient-brasil {
      background: linear-gradient(135deg, #00ff88 0%, #FFDF00 50%, #009739 100%);
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
      background-clip: text;
    }

    /* ===== GLASSMORPHISM ===== */
    .glass {
      background: rgba(10, 25, 47, 0.6);
      backdrop-filter: blur(20px);
      -webkit-backdrop-filter: blur(20px);
      border: 1px solid rgba(0, 255, 136, 0.08);
    }

    .glass-card {
      background: rgba(10, 25, 47, 0.8);
      backdrop-filter: blur(16px);
      -webkit-backdrop-filter: blur(16px);
      border: 1px solid rgba(0, 255, 136, 0.1);
      transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
    }

    .glass-card:hover {
      border-color: rgba(0, 255, 136, 0.25);
      transform: translateY(-6px);
      box-shadow: 0 20px 40px rgba(0, 0, 0, 0.3), 0 0 30px rgba(0, 255, 136, 0.08);
    }

    /* ===== HERO BG ===== */
    .hero-bg {
      background:
        radial-gradient(ellipse 800px 600px at 20% 40%, rgba(0, 151, 57, 0.12) 0%, transparent 70%),
        radial-gradient(ellipse 600px 500px at 80% 20%, rgba(0, 33, 105, 0.15) 0%, transparent 70%),
        radial-gradient(ellipse 400px 300px at 50% 90%, rgba(255, 223, 0, 0.06) 0%, transparent 70%);
    }

    /* ===== DOT GRID ===== */
    .dot-pattern {
      background-image: radial-gradient(rgba(0, 255, 136, 0.04) 1px, transparent 1px);
      background-size: 28px 28px;
    }

    /* ===== SYNTAX COLORS (always dark context) ===== */
    .syn-kw {
      color: #00ff88;
      font-weight: 600;
    }

    .syn-tp {
      color: #FFDF00;
      font-weight: 600;
    }

    .syn-str {
      color: #a8e6a3;
    }

    .syn-cmt {
      color: #546e7a;
      font-style: italic;
    }

    .syn-num {
      color: #ffd700;
    }

    .syn-art {
      color: #80cbc4;
    }

    .syn-op {
      color: #82b1ff;
    }

    .syn-fn {
      color: #00c04b;
    }

    .syn-id {
      color: #ccd6f6;
    }

    /* ===== SCROLL INDICATOR ===== */
    #scroll-indicator {
      position: fixed;
      top: 0;
      left: 0;
      height: 3px;
      z-index: 200;
      background: linear-gradient(90deg, #009739, #00ff88, #FFDF00);
      transition: width 0.15s linear;
    }

    /* ===== MISC ===== */
    html {
      scroll-behavior: smooth;
    }

    ::selection {
      background: rgba(0, 255, 136, 0.3);
    }

    /* Code block wrapper */
    .code-window {
      background: #020c1b;
      border: 1px solid rgba(0, 255, 136, 0.1);
      border-radius: 16px;
      overflow: hidden;
    }

    .code-titlebar {
      background: rgba(0, 255, 136, 0.03);
      border-bottom: 1px solid rgba(0, 255, 136, 0.06);
      padding: 10px 16px;
      display: flex;
      align-items: center;
      justify-content: space-between;
    }

    .code-dots {
      display: flex;
      gap: 6px;
    }

    .code-dots span {
      width: 12px;
      height: 12px;
      border-radius: 50%;
    }

    .code-body {
      padding: 20px 24px;
      font-family: 'JetBrains Mono', monospace;
      font-size: 14px;
      line-height: 1.8;
    }

    /* Navbar link active underline anim */
    .nav-link {
      position: relative;
    }

    .nav-link::after {
      content: '';
      position: absolute;
      bottom: -2px;
      left: 50%;
      width: 0;
      height: 2px;
      background: #00ff88;
      transition: all 0.3s;
      transform: translateX(-50%);
    }

    .nav-link:hover::after {
      width: 100%;
    }
  </style>
</head>

<body class="bg-azul-petroleo text-branco-gelo font-sans">

  <!-- Scroll indicator -->
  <div id="scroll-indicator" style="width: 0%"></div>

  <!-- ============================================
     NAVBAR
     ============================================ -->
  <nav class="fixed top-0 left-0 right-0 z-50 glass border-b border-verde-neon/5">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between h-16">
        <a href="#" class="flex items-center gap-3 group">
          <div
            class="w-9 h-9 rounded-xl bg-gradient-to-br from-verde to-verde-dark flex items-center justify-center text-white font-mono font-bold text-lg shadow-lg shadow-verde/30 group-hover:shadow-verde/50 transition-shadow">
            V</div>
          <span class="text-xl font-bold tracking-tight">Verbo</span>
          <span
            class="hidden sm:inline text-[10px] font-mono text-verde-neon bg-verde-neon/10 px-2 py-0.5 rounded-full border border-verde-neon/20">v2.0</span>
        </a>

        <div class="hidden md:flex items-center gap-8 text-sm">
          <a href="#porque" class="nav-link opacity-60 hover:opacity-100 hover:text-verde-neon transition-all">Por que
            Verbo?</a>
          <a href="#codigo"
            class="nav-link opacity-60 hover:opacity-100 hover:text-verde-neon transition-all">Código</a>
          <a href="#ecossistema"
            class="nav-link opacity-60 hover:opacity-100 hover:text-verde-neon transition-all">Ecossistema</a>
          <a href="docs.php" class="nav-link opacity-60 hover:opacity-100 hover:text-verde-neon transition-all">📖
            Docs</a>
          <a href="#download"
            class="nav-link opacity-60 hover:opacity-100 hover:text-verde-neon transition-all">Download</a>
          <a href="playground.php"
            class="px-5 py-2 rounded-xl bg-verde text-white font-semibold text-sm hover:bg-verde-light hover:shadow-lg hover:shadow-verde/30 transition-all">⚡
            Playground</a>
        </div>

        <div class="flex items-center gap-2">
          <button onclick="verboTheme.toggle()" class="p-2 rounded-lg hover:bg-white/5 transition-colors"
            title="Alternar tema">
            <svg data-theme-icon="light" class="w-5 h-5 hidden" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
            </svg>
            <svg data-theme-icon="dark" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
            </svg>
          </button>
          <button onclick="document.getElementById('mobile-menu').classList.toggle('hidden')"
            class="md:hidden p-2 rounded-lg hover:bg-white/5">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <div id="mobile-menu" class="hidden md:hidden glass border-t border-white/5 px-4 py-4 space-y-3">
      <a href="#porque" class="block py-2 opacity-60 hover:opacity-100">Por que Verbo?</a>
      <a href="#codigo" class="block py-2 opacity-60 hover:opacity-100">Código</a>
      <a href="#ecossistema" class="block py-2 opacity-60 hover:opacity-100">Ecossistema</a>
      <a href="docs.php" class="block py-2 opacity-60 hover:opacity-100">📖 Docs</a>
      <a href="#download" class="block py-2 opacity-60 hover:opacity-100">Download</a>
      <a href="playground.php" class="block py-2 text-verde-neon font-semibold">⚡ Playground →</a>
    </div>
  </nav>

  <!-- ============================================
     HERO SECTION
     ============================================ -->
  <section class="relative min-h-screen flex items-center justify-center pt-16 hero-bg dot-pattern overflow-hidden">
    <!-- Decorative blobs -->
    <div
      class="absolute top-10 right-[10%] w-[500px] h-[500px] rounded-full bg-verde/[0.04] blur-[120px] animate-float">
    </div>
    <div class="absolute bottom-10 left-[5%] w-[600px] h-[600px] rounded-full bg-azul/[0.08] blur-[120px] animate-float"
      style="animation-delay:4s"></div>
    <div
      class="absolute top-[40%] left-[40%] w-[300px] h-[300px] rounded-full bg-amarelo/[0.03] blur-[100px] animate-float"
      style="animation-delay:2s"></div>

    <div class="relative max-w-5xl mx-auto px-4 sm:px-6 text-center">
      <!-- Badge -->
      <div class="animate-fade-in inline-flex items-center gap-2.5 px-5 py-2.5 rounded-full glass text-sm mb-10">
        <span class="w-2 h-2 rounded-full bg-verde-neon animate-pulse shadow-lg shadow-verde-neon/50"></span>
        <span class="opacity-70">Linguagem de programação em Português Brasileiro</span>
      </div>

      <!-- Title -->
      <h1 class="animate-slide-up text-4xl sm:text-5xl md:text-6xl lg:text-7xl font-black leading-[1.1] mb-8">
        Programe com a<br>
        <span class="text-gradient-brasil">alma de um brasileiro</span>
      </h1>

      <!-- Subtitle -->
      <p class="animate-slide-up delay-100 text-lg sm:text-xl text-white/50 max-w-2xl mx-auto mb-10 leading-relaxed">
        Escreva como se estivesse explicando uma ideia. Execute com a
        <span class="text-verde-neon font-semibold">performance do Go</span>.
        Verbo é a primeira linguagem de alto desempenho baseada na gramática brasileira.
      </p>

      <!-- Code window -->
      <div class="animate-slide-up delay-200 mb-12 max-w-lg mx-auto">
        <div class="code-window animate-glow">
          <div class="code-titlebar">
            <div class="code-dots">
              <span class="bg-[#ff5f57]"></span>
              <span class="bg-[#febc2e]"></span>
              <span class="bg-[#28c840]"></span>
            </div>
            <span class="font-mono text-xs text-white/25">manifesto.vrb</span>
          </div>
          <div class="code-body text-left">
            <div><span class="syn-art">A</span> <span class="syn-id">mensagem</span> <span class="syn-op">é</span> <span
                class="syn-str">"O futuro se escreve em português"</span>.</div>
            <div><span class="syn-fn">Exibir</span> <span class="syn-op">com</span> (<span
                class="syn-id">mensagem</span>).</div>
          </div>
        </div>
      </div>

      <!-- CTAs -->
      <div class="animate-slide-up delay-300 flex flex-col sm:flex-row items-center justify-center gap-4">
        <a href="playground.php"
          class="group px-8 py-4 rounded-2xl bg-gradient-to-r from-verde to-verde-dark text-white font-bold text-lg shadow-xl shadow-verde/25 hover:shadow-verde/40 hover:-translate-y-1 transition-all flex items-center gap-3">
          <svg class="w-5 h-5 group-hover:scale-110 transition-transform" fill="currentColor" viewBox="0 0 24 24">
            <path d="M8 5v14l11-7z" />
          </svg>
          Testar no Playground
        </a>
        <a href="#download"
          class="px-8 py-4 rounded-2xl border-2 border-amarelo/40 text-amarelo font-bold text-lg hover:bg-amarelo hover:text-azul-noite hover:border-amarelo transition-all hover:-translate-y-1">
          ⬇ Baixar v2.0
        </a>
      </div>

      <!-- Stats -->
      <div class="animate-slide-up delay-500 mt-20 grid grid-cols-3 gap-8 max-w-md mx-auto">
        <div class="text-center">
          <div class="text-3xl sm:text-4xl font-black text-verde-neon">90+</div>
          <div class="text-xs opacity-40 mt-1">Keywords em PT</div>
        </div>
        <div class="text-center">
          <div class="text-3xl sm:text-4xl font-black text-amarelo">Go</div>
          <div class="text-xs opacity-40 mt-1">Transpilado</div>
        </div>
        <div class="text-center">
          <div class="text-3xl sm:text-4xl font-black text-verde-neon">&lt;1s</div>
          <div class="text-xs opacity-40 mt-1">Compilação</div>
        </div>
      </div>
    </div>

    <!-- Scroll hint -->
    <div class="absolute bottom-10 left-1/2 -translate-x-1/2 animate-bounce opacity-20">
      <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3" />
      </svg>
    </div>
  </section>

  <!-- ============================================
     POR QUE VERBO?
     ============================================ -->
  <section id="porque" class="py-28 relative overflow-hidden">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-20">
        <span class="text-verde-neon font-mono text-xs tracking-[0.2em] uppercase opacity-80">Diferenciais</span>
        <h2 class="text-3xl sm:text-4xl lg:text-5xl font-black mt-4">
          Por que <span class="text-gradient-brasil">Verbo</span>?
        </h2>
        <p class="mt-5 text-white/40 max-w-lg mx-auto text-lg">Uma linguagem projetada para quem pensa em português</p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 max-w-5xl mx-auto">
        <!-- Card 1 -->
        <div class="glass-card rounded-2xl p-7 group">
          <div class="flex items-start gap-4 mb-5">
            <div
              class="w-12 h-12 rounded-xl bg-verde/10 border border-verde/20 flex items-center justify-center shrink-0 group-hover:bg-verde/20 group-hover:border-verde/30 transition-all">
              <span class="text-2xl">🇧🇷</span>
            </div>
            <div>
              <h3 class="text-lg font-bold mb-1">Sintaxe em Português</h3>
              <p class="text-sm text-white/50 leading-relaxed">Variáveis com artigos, funções com verbos. Código que
                parece prosa técnica brasileira.</p>
            </div>
          </div>
          <div class="code-window">
            <div class="code-body text-[13px] !py-3 !px-4">
              <span class="syn-art">A</span> <span class="syn-id">nota</span> <span class="syn-op">é</span> <span
                class="syn-num">10</span>.
            </div>
          </div>
        </div>

        <!-- Card 2 -->
        <div class="glass-card rounded-2xl p-7 group">
          <div class="flex items-start gap-4 mb-5">
            <div
              class="w-12 h-12 rounded-xl bg-amarelo/10 border border-amarelo/20 flex items-center justify-center shrink-0 group-hover:bg-amarelo/20 group-hover:border-amarelo/30 transition-all">
              <span class="text-2xl">⚡</span>
            </div>
            <div>
              <h3 class="text-lg font-bold mb-1">Performance de Go</h3>
              <p class="text-sm text-white/50 leading-relaxed">Transpilado para Go puro. Binários nativos, goroutines,
                sem overhead de interpretador.</p>
            </div>
          </div>
          <div class="code-window">
            <div class="code-body text-[13px] !py-3 !px-4">
              <span class="syn-str">.vrb</span> → <span class="syn-fn">Lexer</span> → <span class="syn-fn">AST</span> →
              <span class="syn-kw">binário nativo</span>
            </div>
          </div>
        </div>

        <!-- Card 3 -->
        <div class="glass-card rounded-2xl p-7 group">
          <div class="flex items-start gap-4 mb-5">
            <div
              class="w-12 h-12 rounded-xl bg-verde-neon/10 border border-verde-neon/20 flex items-center justify-center shrink-0 group-hover:bg-verde-neon/20 group-hover:border-verde-neon/30 transition-all">
              <span class="text-2xl">🔒</span>
            </div>
            <div>
              <h3 class="text-lg font-bold mb-1">Tipagem Forte</h3>
              <p class="text-sm text-white/50 leading-relaxed">Tipos inferidos pela gramática. Imutabilidade semântica:
                "é" vs "está".</p>
            </div>
          </div>
          <div class="code-window">
            <div class="code-body text-[13px] !py-3 !px-4">
              <span class="syn-art">O</span> <span class="syn-id">x</span> <span class="syn-op">é</span> <span
                class="syn-num">42</span>. <span class="syn-cmt">// imutável ✓</span>
            </div>
          </div>
        </div>

        <!-- Card 4 -->
        <div class="glass-card rounded-2xl p-7 group">
          <div class="flex items-start gap-4 mb-5">
            <div
              class="w-12 h-12 rounded-xl bg-azul-light/10 border border-azul-light/20 flex items-center justify-center shrink-0 group-hover:bg-azul-light/20 group-hover:border-azul-light/30 transition-all">
              <span class="text-2xl">🔀</span>
            </div>
            <div>
              <h3 class="text-lg font-bold mb-1">Concorrência Nativa</h3>
              <p class="text-sm text-white/50 leading-relaxed">Goroutines em português. Canais, WaitGroups, tudo com
                sintaxe natural.</p>
            </div>
          </div>
          <div class="code-window">
            <div class="code-body text-[13px] !py-3 !px-4">
              <span class="syn-kw">Simultaneamente</span>:<br>
              &nbsp;&nbsp;<span class="syn-fn">Exibir</span> <span class="syn-str">"Paralelo"</span>.
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>

  <!-- ============================================
     CÓDIGO VIVO
     ============================================ -->
  <section id="codigo" class="py-28 relative dot-pattern">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-20">
        <span class="text-verde-neon font-mono text-xs tracking-[0.2em] uppercase opacity-80">Exemplos</span>
        <h2 class="text-3xl sm:text-4xl lg:text-5xl font-black mt-4">
          Código <span class="text-gradient-brasil">Vivo</span>
        </h2>
        <p class="mt-5 text-white/40 max-w-lg mx-auto text-lg">Veja a linguagem em ação — cada exemplo pode ser testado
          no Playground</p>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 max-w-5xl mx-auto">
        <!-- Example 1 -->
        <div class="code-window glass-card !border-verde-neon/8">
          <div class="code-titlebar">
            <div class="flex items-center gap-3">
              <div class="code-dots"><span class="bg-[#ff5f57]"></span><span class="bg-[#febc2e]"></span><span
                  class="bg-[#28c840]"></span></div>
              <span class="font-mono text-xs text-white/30">ola_mundo.vrb</span>
            </div>
            <a href="playground.php?exemplo=ola_mundo"
              class="text-xs text-verde-neon hover:text-verde-light transition-colors font-semibold">▶ Testar</a>
          </div>
          <div class="code-body">
            <div class="syn-cmt">// Meu primeiro programa em Verbo</div>
            <div><span class="syn-art">A</span> <span class="syn-id">saudacao</span> <span class="syn-op">é</span> <span
                class="syn-str">"Olá, Mundo!"</span>.</div>
            <div><span class="syn-fn">Exibir</span> <span class="syn-op">com</span> (<span
                class="syn-id">saudacao</span>).</div>
            <div class="mt-3 pt-3 border-t border-white/5 syn-cmt">→ Olá, Mundo!</div>
          </div>
        </div>

        <!-- Example 2 -->
        <div class="code-window glass-card !border-verde-neon/8">
          <div class="code-titlebar">
            <div class="flex items-center gap-3">
              <div class="code-dots"><span class="bg-[#ff5f57]"></span><span class="bg-[#febc2e]"></span><span
                  class="bg-[#28c840]"></span></div>
              <span class="font-mono text-xs text-white/30">funcoes.vrb</span>
            </div>
            <a href="playground.php?exemplo=funcoes"
              class="text-xs text-verde-neon hover:text-verde-light transition-colors font-semibold">▶ Testar</a>
          </div>
          <div class="code-body">
            <div><span class="syn-kw">Para</span> <span class="syn-id">Dobrar</span> <span class="syn-op">usando</span>
              (<span class="syn-id">x</span>: <span class="syn-tp">Inteiro</span>):</div>
            <div>&nbsp;&nbsp;<span class="syn-kw">Retorne</span> <span class="syn-id">x</span> * <span
                class="syn-num">2</span>.</div>
            <div><span class="syn-op">.</span></div>
            <div class="mt-2"><span class="syn-art">O</span> <span class="syn-id">resultado</span> <span
                class="syn-op">é</span> <span class="syn-id">Dobrar</span> <span class="syn-op">com</span> (<span
                class="syn-num">21</span>).</div>
            <div><span class="syn-fn">Exibir</span> <span class="syn-op">com</span> (<span
                class="syn-id">resultado</span>).</div>
            <div class="mt-3 pt-3 border-t border-white/5 syn-cmt">→ 42</div>
          </div>
        </div>

        <!-- Example 3 -->
        <div class="code-window glass-card !border-verde-neon/8">
          <div class="code-titlebar">
            <div class="flex items-center gap-3">
              <div class="code-dots"><span class="bg-[#ff5f57]"></span><span class="bg-[#febc2e]"></span><span
                  class="bg-[#28c840]"></span></div>
              <span class="font-mono text-xs text-white/30">concorrencia.vrb</span>
            </div>
            <a href="playground.php?exemplo=concorrencia"
              class="text-xs text-verde-neon hover:text-verde-light transition-colors font-semibold">▶ Testar</a>
          </div>
          <div class="code-body">
            <div><span class="syn-kw">Simultaneamente</span>:</div>
            <div>&nbsp;&nbsp;<span class="syn-fn">Exibir</span> <span class="syn-op">com</span> (<span
                class="syn-str">"Brasil no topo"</span>).</div>
            <div>&nbsp;&nbsp;<span class="syn-fn">Exibir</span> <span class="syn-op">com</span> (<span
                class="syn-str">"Performance máxima"</span>).</div>
            <div class="mt-3 pt-3 border-t border-white/5 syn-cmt">→ Goroutines em ação!</div>
          </div>
        </div>

        <!-- Example 4 -->
        <div class="code-window glass-card !border-verde-neon/8">
          <div class="code-titlebar">
            <div class="flex items-center gap-3">
              <div class="code-dots"><span class="bg-[#ff5f57]"></span><span class="bg-[#febc2e]"></span><span
                  class="bg-[#28c840]"></span></div>
              <span class="font-mono text-xs text-white/30">entidade.vrb</span>
            </div>
            <a href="playground.php?exemplo=entidade"
              class="text-xs text-verde-neon hover:text-verde-light transition-colors font-semibold">▶ Testar</a>
          </div>
          <div class="code-body">
            <div><span class="syn-art">A</span> <span class="syn-kw">Entidade</span> <span class="syn-tp">Pessoa</span>
              <span class="syn-op">contendo</span>
            </div>
            <div>&nbsp;&nbsp;(<span class="syn-id">Nome</span>: <span class="syn-tp">Texto</span>, <span
                class="syn-id">Idade</span>: <span class="syn-tp">Inteiro</span>).</div>
            <div><span class="syn-art">Um</span> <span class="syn-id">dev</span> <span class="syn-op">é</span>
              <span class="syn-tp">Pessoa</span> <span class="syn-op">com</span> (<span class="syn-str">"Juan"</span>,
              <span class="syn-num">25</span>).
            </div>
            <div class="mt-3 pt-3 border-t border-white/5 syn-cmt">→ Structs tipadas!</div>
          </div>
        </div>
      </div>
    </div>
  </section>

  <!-- ============================================
     ECOSSISTEMA
     ============================================ -->
  <section id="ecossistema" class="py-28 relative">
    <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-20">
        <span class="text-amarelo font-mono text-xs tracking-[0.2em] uppercase opacity-80">Plataforma</span>
        <h2 class="text-3xl sm:text-4xl lg:text-5xl font-black mt-4">Ecossistema <span
            class="text-gradient-brasil">Crom</span></h2>
      </div>

      <div class="glass-card rounded-3xl p-8 sm:p-12 max-w-4xl mx-auto">
        <div class="grid grid-cols-1 sm:grid-cols-3 gap-6 mb-10">
          <div
            class="text-center p-6 rounded-2xl bg-verde/[0.05] border border-verde/10 hover:border-verde/20 transition-colors">
            <div class="text-3xl mb-3">🔨</div>
            <h4 class="font-bold text-verde-neon mb-1">Compilador</h4>
            <p class="text-xs text-white/40">Lexer → Parser → AST → Go</p>
            <div class="mt-3 font-mono text-[10px] text-white/20">pkg/</div>
          </div>
          <div
            class="text-center p-6 rounded-2xl bg-amarelo/[0.05] border border-amarelo/10 hover:border-amarelo/20 transition-colors">
            <div class="text-3xl mb-3">📚</div>
            <h4 class="font-bold text-amarelo mb-1">BibVerbo</h4>
            <p class="text-xs text-white/40">Matemática, Texto, Arquivo, Html</p>
            <div class="mt-3 font-mono text-[10px] text-white/20">pkg/stdlib/</div>
          </div>
          <div
            class="text-center p-6 rounded-2xl bg-azul-light/[0.05] border border-azul-light/10 hover:border-azul-light/20 transition-colors">
            <div class="text-3xl mb-3">🛠️</div>
            <h4 class="font-bold text-[#64b5f6] mb-1">Ferramentas</h4>
            <p class="text-xs text-white/40">CLI, Playground, Docs</p>
            <div class="mt-3 font-mono text-[10px] text-white/20">cmd/ + www/</div>
          </div>
        </div>

        <!-- Pipeline -->
        <div class="p-6 rounded-2xl bg-white/[0.02] border border-white/5">
          <h4 class="font-mono text-[10px] text-verde-neon mb-4 tracking-[0.15em] uppercase">Pipeline de Compilação</h4>
          <div class="flex items-center justify-center gap-2 sm:gap-3 flex-wrap font-mono text-sm">
            <span class="px-3 py-2 rounded-xl bg-verde/10 text-verde-neon border border-verde/20">.vrb</span>
            <span class="text-white/15">→</span>
            <span class="px-3 py-2 rounded-xl bg-amarelo/10 text-amarelo border border-amarelo/10">Lexer</span>
            <span class="text-white/15">→</span>
            <span class="px-3 py-2 rounded-xl bg-amarelo/10 text-amarelo border border-amarelo/10">Parser</span>
            <span class="text-white/15">→</span>
            <span class="px-3 py-2 rounded-xl bg-amarelo/10 text-amarelo border border-amarelo/10">AST</span>
            <span class="text-white/15">→</span>
            <span class="px-3 py-2 rounded-xl bg-verde/10 text-verde-neon border border-verde/20">.go</span>
            <span class="text-white/15">→</span>
            <span class="px-3 py-2 rounded-xl bg-white/5 text-[#64b5f6] border border-white/10">Binário</span>
          </div>
        </div>
      </div>
    </div>
  </section>

  <!-- ============================================
     DOCUMENTAÇÃO (highlight)
     ============================================ -->
  <section id="docs" class="py-28 relative dot-pattern">
    <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-16">
        <span class="text-verde-neon font-mono text-xs tracking-[0.2em] uppercase opacity-80">Aprenda</span>
        <h2 class="text-3xl sm:text-4xl lg:text-5xl font-black mt-4">
          Documentação <span class="text-gradient-brasil">Interativa</span>
        </h2>
        <p class="mt-5 text-white/40 max-w-2xl mx-auto text-lg">Aprenda Verbo no seu ritmo. Com mini playground
          integrado para testar código enquanto lê, e o modo <strong class="text-amarelo">Ajudante</strong> que explica
          cada conceito para quem nunca programou.</p>
      </div>

      <div class="glass-card rounded-3xl p-8 sm:p-12 max-w-4xl mx-auto">
        <div class="grid grid-cols-1 sm:grid-cols-3 gap-6 mb-8">
          <div class="text-center p-5 rounded-2xl bg-verde/[0.05] border border-verde/10">
            <div class="text-3xl mb-3">📖</div>
            <h4 class="font-bold text-verde-neon text-sm mb-1">Especificação Completa</h4>
            <p class="text-xs text-white/40">Tipos, variáveis, funções, loops, operadores — tudo documentado</p>
          </div>
          <div class="text-center p-5 rounded-2xl bg-amarelo/[0.05] border border-amarelo/10">
            <div class="text-3xl mb-3">🧪</div>
            <h4 class="font-bold text-amarelo text-sm mb-1">Mini Playground</h4>
            <p class="text-xs text-white/40">Edite e execute código diretamente dentro da documentação</p>
          </div>
          <div class="text-center p-5 rounded-2xl bg-azul-light/[0.05] border border-azul-light/10">
            <div class="text-3xl mb-3">🤝</div>
            <h4 class="font-bold text-[#64b5f6] text-sm mb-1">Modo Ajudante</h4>
            <p class="text-xs text-white/40">Explicações detalhadas para quem nunca programou antes</p>
          </div>
        </div>

        <div class="text-center">
          <a href="docs.php"
            class="inline-flex items-center gap-3 px-8 py-4 rounded-2xl bg-gradient-to-r from-azul to-azul-light text-white font-bold text-lg shadow-xl shadow-azul/25 hover:shadow-azul/40 hover:-translate-y-1 transition-all group">
            📖 Abrir Documentação
            <svg class="w-5 h-5 group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor"
              viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7l5 5m0 0l-5 5m5-5H6" />
            </svg>
          </a>
        </div>
      </div>
    </div>
  </section>

  <!-- ============================================
     DOWNLOAD
     ============================================ -->
  <section id="download" class="py-28 relative dot-pattern">
    <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
      <span class="text-verde-neon font-mono text-xs tracking-[0.2em] uppercase opacity-80">Instalar</span>
      <h2 class="text-3xl sm:text-4xl lg:text-5xl font-black mt-4 mb-5">Baixe o <span
          class="text-gradient-brasil">Verbo</span></h2>
      <p class="text-white/40 max-w-lg mx-auto mb-14 text-lg">Comece a programar em português agora.</p>

      <!-- Install command -->
      <div class="code-window max-w-2xl mx-auto mb-10">
        <div class="code-titlebar">
          <span class="font-mono text-xs text-verde-neon/60">Instalação rápida (Linux/macOS)</span>
          <button
            onclick="navigator.clipboard.writeText('curl -fsSL https://raw.githubusercontent.com/juanxto/crom-verbo/main/install.sh | bash')"
            class="text-xs text-white/20 hover:text-verde-neon transition-colors cursor-pointer">📋 Copiar</button>
        </div>
        <div class="code-body text-left">
          <span class="syn-kw">$</span> curl -fsSL https://raw.githubusercontent.com/juanxto/crom-verbo/main/install.sh
          | bash
        </div>
      </div>

      <!-- Download cards -->
      <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 max-w-2xl mx-auto mb-12">
        <a href="downloads/verbo-linux-amd64" class="glass-card rounded-2xl p-6 text-center group" id="dl-linux">
          <div class="text-4xl mb-3">🐧</div>
          <h4 class="font-bold mb-1">Linux</h4>
          <span class="text-xs text-white/30">amd64 / arm64</span>
        </a>
        <a href="downloads/verbo-darwin-arm64" class="glass-card rounded-2xl p-6 text-center group" id="dl-macos">
          <div class="text-4xl mb-3">🍎</div>
          <h4 class="font-bold mb-1">macOS</h4>
          <span class="text-xs text-white/30">Apple Silicon & Intel</span>
        </a>
        <a href="downloads/verbo-windows-amd64.exe" class="glass-card rounded-2xl p-6 text-center group"
          id="dl-windows">
          <div class="text-4xl mb-3">🪟</div>
          <h4 class="font-bold mb-1">Windows</h4>
          <span class="text-xs text-white/30">amd64</span>
        </a>
      </div>

      <!-- Quick start -->
      <div class="code-window max-w-2xl mx-auto text-left">
        <div class="code-titlebar">
          <span class="font-mono text-xs text-amarelo/60">Começar em 3 passos</span>
        </div>
        <div class="code-body space-y-2">
          <div><span class="syn-kw">1.</span> <span class="syn-cmt">echo '</span><span class="syn-fn">Exibir</span>
            <span class="syn-op">com</span> (<span class="syn-str">"Olá!"</span>).<span class="syn-cmt">' >
              ola.vrb</span>
          </div>
          <div><span class="syn-kw">2.</span> <span class="syn-id">verbo executar ola.vrb</span></div>
          <div><span class="syn-kw">3.</span> <span class="syn-kw">→ Olá!</span></div>
        </div>
      </div>
    </div>
  </section>

  <!-- ============================================
     FOOTER
     ============================================ -->
  <footer class="border-t border-white/5 py-14">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-10">
        <div>
          <div class="flex items-center gap-3 mb-4">
            <div
              class="w-8 h-8 rounded-lg bg-gradient-to-br from-verde to-verde-dark flex items-center justify-center text-white font-mono font-bold">
              V</div>
            <span class="text-lg font-bold">Verbo</span>
          </div>
          <p class="text-sm text-white/30 leading-relaxed">A primeira linguagem de programação de alto desempenho
            baseada na gramática do Português Brasileiro.</p>
        </div>
        <div>
          <h4 class="font-bold text-sm mb-4 text-white/50">Recursos</h4>
          <ul class="space-y-3 text-sm text-white/30">
            <li><a href="playground.php" class="hover:text-verde-neon transition-colors">⚡ Playground</a></li>
            <li><a href="docs.php" class="hover:text-verde-neon transition-colors">📖 Documentação</a></li>
            <li><a href="https://github.com/juanxto/crom-verbo" target="_blank"
                class="hover:text-verde-neon transition-colors">📦 GitHub ↗</a></li>
          </ul>
        </div>
        <div>
          <h4 class="font-bold text-sm mb-4 text-white/50">Projeto</h4>
          <p class="text-sm text-white/30">Criado por <span class="text-verde-neon font-semibold">Juan</span></p>
          <p class="text-sm text-white/30 mt-1">Ecossistema Crom 🇧🇷</p>
          <p class="text-sm text-white/15 mt-3 font-mono">MIT License</p>
        </div>
      </div>
      <div class="mt-12 pt-6 border-t border-white/5 text-center">
        <p class="text-xs text-white/15">Feito com 💚💛 no Brasil — © 2025 Projeto Crom</p>
      </div>
    </div>
  </footer>

  <script src="js/theme.js"></script>
  <script>
    window.addEventListener('scroll', () => {
      const p = (window.scrollY / (document.body.scrollHeight - window.innerHeight)) * 100;
      document.getElementById('scroll-indicator').style.width = p + '%';
    });
    // Auto-detect OS
    const ua = navigator.userAgent.toLowerCase();
    if (ua.includes('linux')) document.getElementById('dl-linux')?.classList.add('animate-glow');
    else if (ua.includes('mac')) document.getElementById('dl-macos')?.classList.add('animate-glow');
    else if (ua.includes('win')) document.getElementById('dl-windows')?.classList.add('animate-glow');
  </script>
</body>

</html>