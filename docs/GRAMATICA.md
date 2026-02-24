# 📐 Gramática Formal da Linguagem Verbo

**Notação**: EBNF (Extended Backus-Naur Form)

---

## Produções Principais

```ebnf
programa        = { declaracao } ;

declaracao      = decl_variavel
                | decl_funcao
                | decl_exibir
                | decl_se
                | decl_repita
                | decl_enquanto
                | decl_retorne
                | decl_atribuicao
                | decl_expressao ;
```

## Variáveis e Constantes

```ebnf
decl_variavel   = artigo IDENTIFICADOR verbo_atrib expressao PONTO ;

artigo          = ARTIGO_DEFINIDO       (* O | A | Os | As *)
                | ARTIGO_INDEFINIDO ;   (* Um | Uma *)

verbo_atrib     = "é" | "está" ;
```

## Funções

```ebnf
decl_funcao     = "Para" IDENTIFICADOR [ "usando" parametros ] DOIS_PONTOS
                  bloco ;

parametros      = "(" param { VIRGULA param } ")" ;

param           = IDENTIFICADOR [ DOIS_PONTOS TIPO ] ;

chamada_funcao  = IDENTIFICADOR "com" argumentos ;

argumentos      = "(" expressao { VIRGULA expressao } ")"
                | expressao ;
```

## Controle de Fluxo

```ebnf
decl_se         = "Se" expr_condicional [ VIRGULA ] [ "então" ] DOIS_PONTOS
                  bloco
                  [ "Senão" [ DOIS_PONTOS ] bloco ] ;

expr_condicional = [ artigo ] expressao [ "for" ] op_comparacao expressao ;

op_comparacao   = "menor" [ "que" ]
                | "maior" [ "que" ]
                | "igual" ;
```

## Loops

```ebnf
decl_repita     = "Repita" expressao "vezes" DOIS_PONTOS bloco
                | "Repita" "para" "cada" IDENTIFICADOR "em" expressao DOIS_PONTOS bloco ;

decl_enquanto   = "Enquanto" expr_condicional DOIS_PONTOS bloco ;
```

## E/S e Retorno

```ebnf
decl_exibir     = "Exibir" [ "com" ] ( "(" expressao ")" | expressao ) PONTO ;

decl_retorne    = "Retorne" [ expressao | "Nulo" ] PONTO ;
```

## Expressões

```ebnf
expressao       = expr_aditiva ;

expr_aditiva    = expr_multiplicativa { ( "+" | "-" | "e" | "menos" ) expr_multiplicativa } ;

expr_multiplicativa = expr_primaria { ( "*" | "/" ) expr_primaria } ;

expr_primaria   = NUMERO
                | TEXTO
                | "Verdadeiro" | "Falso"
                | "Nulo"
                | IDENTIFICADOR [ "com" argumentos ]
                | "(" expressao ")"
                | "não" expr_primaria
                | artigo IDENTIFICADOR ;
```

## Outros

```ebnf
decl_atribuicao = IDENTIFICADOR "está" expressao PONTO ;

decl_expressao  = IDENTIFICADOR [ "com" argumentos ] PONTO ;

bloco           = { declaracao } ;
```

## Tokens Terminais

```ebnf
IDENTIFICADOR   = letra { letra | digito | "_" } ;
NUMERO          = digito { digito } [ "." digito { digito } ] ;
TEXTO           = '"' { caractere } '"' ;
TIPO            = "Texto" | "Inteiro" | "Decimal" | "Logico" | "Lógico" | "Lista" ;

PONTO           = "." ;
DOIS_PONTOS     = ":" ;
VIRGULA         = "," ;

letra           = ? qualquer caractere Unicode da categoria "Letter" ? ;
digito          = "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9" ;
caractere       = ? qualquer caractere exceto '"' não-escapado ? ;

COMENTARIO      = "//" { caractere_sem_newline } ;
```

---

## Precedência de Operadores

| Prioridade | Operadores        | Associatividade |
|------------|-------------------|-----------------|
| 1 (maior)  | `não`, `-` (unário) | Direita       |
| 2          | `*`, `/`          | Esquerda        |
| 3          | `+`, `-`, `e`     | Esquerda        |
| 4 (menor)  | `menor que`, `maior que`, `igual` | — |
