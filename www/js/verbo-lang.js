/**
 * Verbo Language — Monaco Editor Monarch Tokenizer
 * Mapeado diretamente de pkg/lexer/token.go
 *
 * Define syntax highlighting para arquivos .vrb
 */
function registerVerboLanguage(monaco) {
    // Register the language
    monaco.languages.register({
        id: 'verbo',
        extensions: ['.vrb'],
        aliases: ['Verbo', 'verbo'],
        mimetypes: ['text/x-verbo'],
    });

    // Monarch tokenizer
    monaco.languages.setMonarchTokensProvider('verbo', {
        defaultToken: '',
        ignoreCase: false,

        // Articles
        articles: [
            'O', 'A', 'Os', 'As', 'o', 'a', 'os', 'as',
            'Um', 'Uma', 'um', 'uma',
        ],

        // Control flow
        keywords: [
            'Para', 'para', 'Se', 'se', 'Senão', 'senão',
            'Enquanto', 'enquanto', 'Repita', 'repita',
            'Retorne', 'retorne', 'então', 'Então',
            'cada', 'em', 'vezes',
        ],

        // Concurrency & Error handling
        concurrency: [
            'Simultaneamente', 'simultaneamente',
            'Aguarde', 'aguarde',
            'Canal', 'canal',
            'Enviar', 'enviar',
            'Receber', 'receber',
            'Tente', 'tente',
            'Capture', 'capture',
            'Sinalize', 'sinalize',
        ],

        // Types
        typeKeywords: [
            'Texto', 'Inteiro', 'Decimal', 'Logico', 'Lógico', 'Lista',
        ],

        // Builtins
        builtins: [
            'Exibir', 'exibir', 'Incluir', 'incluir',
            'Entidade', 'entidade', 'contendo', 'Contendo',
            'novo', 'Novo',
        ],

        // Operators (semantic)
        operators: [
            'é', 'É', 'está', 'Está',
            'com', 'usando', 'de', 'do', 'da', 'dos', 'das',
            'for', 'menor', 'maior', 'igual', 'que', 'não', 'Não',
            'ao', 'aos', 'no', 'na', 'nos', 'nas',
            'pelo', 'pela', 'pelos', 'pelas', 'por',
            'soma', 'subtrai', 'multiplica', 'divide', 'porcentagem',
            'módulo', 'modulo', 'idêntico', 'identico', 'diferente',
        ],

        // Literals
        literals: [
            'Verdadeiro', 'verdadeiro', 'Falso', 'falso',
            'Nulo', 'nulo',
        ],

        // Demonstratives
        demonstratives: [
            'Este', 'Esta', 'Aquele', 'Aquela',
            'este', 'esta', 'aquele', 'aquela',
        ],

        // Symbols
        symbols: /[=><!~?:&|+\-*\/\^%]+/,

        tokenizer: {
            root: [
                // Comments
                [/\/\/.*$/, 'comment'],

                // Strings
                [/"([^"\\]|\\.)*"/, 'string'],

                // Numbers
                [/\d+\.\d+/, 'number.float'],
                [/\d+/, 'number'],

                // Identifiers and keywords
                [/[a-zA-ZÀ-ÿ_][a-zA-ZÀ-ÿ0-9_]*/, {
                    cases: {
                        '@articles': 'variable.language',
                        '@keywords': 'keyword',
                        '@concurrency': 'keyword.control',
                        '@typeKeywords': 'type',
                        '@builtins': 'support.function',
                        '@operators': 'keyword.operator',
                        '@literals': 'constant.language',
                        '@demonstratives': 'variable.language',
                        '@default': 'identifier',
                    },
                }],

                // Operators
                [/[+\-*\/%]/, 'operator'],
                [/[!=]=/, 'operator'],
                [/[=]/, 'operator'],

                // Delimiters
                [/[.]/, 'delimiter.period'],
                [/:/, 'delimiter.colon'],
                [/,/, 'delimiter.comma'],
                [/[()]/, 'delimiter.parenthesis'],
                [/[\[\]]/, 'delimiter.bracket'],
                [/[{}]/, 'delimiter.bracket'],

                // Whitespace
                [/\s+/, 'white'],
            ],
        },
    });

    // Auto-completion
    monaco.languages.registerCompletionItemProvider('verbo', {
        provideCompletionItems: function (model, position) {
            const word = model.getWordUntilPosition(position);
            const range = {
                startLineNumber: position.lineNumber,
                endLineNumber: position.lineNumber,
                startColumn: word.startColumn,
                endColumn: word.endColumn,
            };

            const suggestions = [
                // Control
                { label: 'Para', kind: monaco.languages.CompletionItemKind.Keyword, insertText: 'Para ${1:Nome} usando (${2:param}: ${3:Tipo}) :\n    ${0}\n.', insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet, detail: 'Declarar função', range },
                { label: 'Se', kind: monaco.languages.CompletionItemKind.Keyword, insertText: 'Se ${1:condicao} então :\n    ${0}\n.', insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet, detail: 'Condicional', range },
                { label: 'Enquanto', kind: monaco.languages.CompletionItemKind.Keyword, insertText: 'Enquanto ${1:condicao} :\n    ${0}\n.', insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet, detail: 'Loop condicional', range },
                { label: 'Repita', kind: monaco.languages.CompletionItemKind.Keyword, insertText: 'Repita ${1:N} vezes :\n    ${0}\n.', insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet, detail: 'Loop por contagem', range },
                { label: 'Simultaneamente', kind: monaco.languages.CompletionItemKind.Keyword, insertText: 'Simultaneamente :\n    ${0}\n.', insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet, detail: 'Execução concorrente', range },
                { label: 'Tente', kind: monaco.languages.CompletionItemKind.Keyword, insertText: 'Tente :\n    ${1}\nCapture ${2:erro} :\n    ${0}\n.', insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet, detail: 'Tratamento de erros', range },
                // Declarations
                { label: 'Exibir', kind: monaco.languages.CompletionItemKind.Function, insertText: 'Exibir "${1:mensagem}".', insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet, detail: 'Saída padrão', range },
                { label: 'Retorne', kind: monaco.languages.CompletionItemKind.Keyword, insertText: 'Retorne ${1:valor}.', insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet, detail: 'Retornar valor', range },
                { label: 'Incluir', kind: monaco.languages.CompletionItemKind.Module, insertText: 'Incluir ${1:Pacote}.', insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet, detail: 'Importar pacote', range },
                { label: 'Entidade', kind: monaco.languages.CompletionItemKind.Struct, insertText: 'A Entidade ${1:Nome} contendo (${2:campo}: ${3:Tipo}).', insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet, detail: 'Declarar struct', range },
                // Types
                { label: 'Texto', kind: monaco.languages.CompletionItemKind.TypeParameter, insertText: 'Texto', detail: 'Tipo string', range },
                { label: 'Inteiro', kind: monaco.languages.CompletionItemKind.TypeParameter, insertText: 'Inteiro', detail: 'Tipo int', range },
                { label: 'Decimal', kind: monaco.languages.CompletionItemKind.TypeParameter, insertText: 'Decimal', detail: 'Tipo float64', range },
                { label: 'Lógico', kind: monaco.languages.CompletionItemKind.TypeParameter, insertText: 'Lógico', detail: 'Tipo bool', range },
                { label: 'Lista', kind: monaco.languages.CompletionItemKind.TypeParameter, insertText: 'Lista', detail: 'Tipo slice', range },
            ];

            return { suggestions };
        },
    });
}

// Define custom themes
function registerVerboThemes(monaco) {
    // Brasil Dark theme
    monaco.editor.defineTheme('verbo-dark', {
        base: 'vs-dark',
        inherit: true,
        rules: [
            { token: 'comment', foreground: '94a3b8', fontStyle: 'italic' },
            { token: 'string', foreground: 'a8e6a3' },
            { token: 'number', foreground: 'ffd700' },
            { token: 'number.float', foreground: 'ffd700' },
            { token: 'keyword', foreground: '00ff88', fontStyle: 'bold' },
            { token: 'keyword.control', foreground: '64b5f6', fontStyle: 'bold' },
            { token: 'keyword.operator', foreground: '82b1ff' },
            { token: 'type', foreground: 'FFDF00', fontStyle: 'bold' },
            { token: 'support.function', foreground: '00c04b' },
            { token: 'variable.language', foreground: '80cbc4' },
            { token: 'constant.language', foreground: 'ff8a65' },
            { token: 'identifier', foreground: 'e0e0e0' },
            { token: 'delimiter.period', foreground: '94a3b8' },
            { token: 'delimiter.colon', foreground: '94a3b8' },
            { token: 'delimiter.comma', foreground: '94a3b8' },
            { token: 'delimiter.parenthesis', foreground: 'FFDF00' },
            { token: 'delimiter.bracket', foreground: '64b5f6' },
            { token: 'operator', foreground: '82b1ff' },
        ],
        colors: {
            'editor.background': '#0a192f',
            'editor.foreground': '#e0e0e0',
            'editor.lineHighlightBackground': '#112240',
            'editor.selectionBackground': '#1a3a5c',
            'editorCursor.foreground': '#00ff88',
            'editorLineNumber.foreground': '#334155',
            'editorLineNumber.activeForeground': '#00ff88',
            'editor.selectionHighlightBackground': '#14533340',
            'editorBracketMatch.background': '#00ff8820',
            'editorBracketMatch.border': '#00ff88',
        },
    });

    // Brasil Light theme
    monaco.editor.defineTheme('verbo-light', {
        base: 'vs',
        inherit: true,
        rules: [
            { token: 'comment', foreground: '94a3b8', fontStyle: 'italic' },
            { token: 'string', foreground: '2e7d32' },
            { token: 'number', foreground: 'e65100' },
            { token: 'number.float', foreground: 'e65100' },
            { token: 'keyword', foreground: '009739', fontStyle: 'bold' },
            { token: 'keyword.control', foreground: '012169', fontStyle: 'bold' },
            { token: 'keyword.operator', foreground: '1565c0' },
            { token: 'type', foreground: 'b8860b', fontStyle: 'bold' },
            { token: 'support.function', foreground: '006b28' },
            { token: 'variable.language', foreground: '00695c' },
            { token: 'constant.language', foreground: 'd84315' },
            { token: 'identifier', foreground: '212121' },
            { token: 'delimiter.period', foreground: '757575' },
            { token: 'delimiter.colon', foreground: '757575' },
            { token: 'delimiter.comma', foreground: '757575' },
            { token: 'delimiter.parenthesis', foreground: 'b8860b' },
            { token: 'delimiter.bracket', foreground: '012169' },
            { token: 'operator', foreground: '1565c0' },
        ],
        colors: {
            'editor.background': '#ffffff',
            'editor.foreground': '#212121',
            'editor.lineHighlightBackground': '#f5f5f5',
            'editor.selectionBackground': '#c8e6c966',
            'editorCursor.foreground': '#009739',
            'editorLineNumber.foreground': '#bdbdbd',
            'editorLineNumber.activeForeground': '#009739',
            'editorBracketMatch.background': '#00973920',
            'editorBracketMatch.border': '#009739',
        },
    });
}

// Export both registration functions
if (typeof module !== 'undefined' && module.exports) {
    module.exports = { registerVerboLanguage, registerVerboThemes };
}
