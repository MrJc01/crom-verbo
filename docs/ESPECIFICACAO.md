# 📖 Especificação da Linguagem Verbo

**Versão**: 0.1.0 (MVP)  
**Status**: Em Desenvolvimento  
**Autor**: Juan / Projeto Crom

---

## 1. Visão Geral

**Verbo** é uma linguagem de programação transpilada que utiliza a gramática da norma culta do Português Brasileiro como sintaxe. Ela é transpilada para Go, garantindo performance nativa.

### Princípios de Design
1. **Legibilidade máxima** — Código Verbo deve parecer prosa técnica em português
2. **Gramática como semântica** — Artigos, verbos e preposições têm significado lógico
3. **Tipagem forte inferida** — O sistema de tipos é inferido pela gramática
4. **Sem ambiguidade** — Ordem SVO (Sujeito-Verbo-Objeto) estrita

---

## 2. Tipos de Dados

| Tipo      | Descrição           | Exemplo            |
|-----------|---------------------|--------------------|
| `Texto`   | Cadeia de caracteres| `"Olá"`            |
| `Inteiro` | Número inteiro      | `42`               |
| `Decimal` | Número de ponto flutuante | `3.14`       |
| `Lógico`  | Verdadeiro/Falso    | `Verdadeiro`       |
| `Nulo`    | Ausência de valor   | `Nulo`             |

---

## 3. Variáveis e Constantes

### 3.1 Constantes (Artigo Definido)
```
O limite é 100.
A mensagem é "Olá, Mundo!".
```
Artigos `O` / `A` declaram valores **imutáveis**.

### 3.2 Variáveis (Artigo Indefinido)
```
Um contador está 0.
Uma taxa está 0.15.
```
Artigos `Um` / `Uma` declaram valores **mutáveis**.

### 3.3 Semântica de Estado
- **`é`** → Atribuição estática (natureza/definição)
- **`está`** → Atribuição de estado (temporário/mutável)

---

## 4. Funções

### 4.1 Declaração
```
Para Calcular usando (valor: Inteiro):
    Retorne valor + 10.
```

### 4.2 Chamada de Função
```
Um resultado é Calcular com (5).
Exibir com (resultado).
```

### 4.3 Saída Padrão
```
Exibir com ("Mensagem aqui").
Exibir com (variavel).
```

---

## 5. Controle de Fluxo

### 5.1 Condicional (Se/Senão)
```
Se a idade for menor que 18, então:
    Exibir com ("Menor de idade").
Senão:
    Exibir com ("Maior de idade").
```

### 5.2 Operadores de Comparação
| Verbo          | Significado |
|----------------|-------------|
| `menor que`    | `<`         |
| `maior que`    | `>`         |
| `igual`        | `==`        |

---

## 6. Loops

### 6.1 Repetição por Contagem
```
Repita 10 vezes:
    Exibir com ("Iteração").
```

### 6.2 Repetição Condicional
```
Enquanto o contador for menor que 100:
    contador está contador + 1.
```

### 6.3 Iteração sobre Coleção (futuro)
```
Repita para cada item em lista:
    Exibir com (item).
```

---

## 7. Operadores

### 7.1 Aritméticos
| Operador | Símbolo | Exemplo      |
|----------|---------|-------------|
| Adição   | `+`     | `a + b`     |
| Subtração| `-`     | `a - b`     |
| Multiplicação | `*` | `a * b`   |
| Divisão  | `/`     | `a / b`     |

### 7.2 Concatenação
A palavra `e` funciona como operador de concatenação/adição:
```
Exibir com ("Olá" e " Mundo").
```

---

## 8. Comentários

```
// Isto é um comentário de linha
A versao é 1.  // Comentário inline
```

---

## 9. Fim de Instrução

Toda instrução termina com ponto final (`.`):
```
A mensagem é "Olá".
Um numero está 42.
Exibir com (mensagem).
```

---

## 10. Palavras Reservadas

| Palavra       | Função                    |
|---------------|---------------------------|
| `O/A/Os/As`   | Artigo definido (constante)|
| `Um/Uma`      | Artigo indefinido (variável)|
| `é`           | Atribuição estática        |
| `está`        | Atribuição de estado       |
| `Para`        | Declaração de função       |
| `usando`      | Parâmetros de função       |
| `com`         | Argumentos de chamada      |
| `Se`          | Condicional                |
| `Senão`       | Alternativa condicional    |
| `então`       | Início do bloco condicional|
| `for`         | Subjuntivo em comparações  |
| `Repita`      | Loop                       |
| `vezes`       | Contagem de repetições     |
| `Enquanto`    | Loop condicional           |
| `Retorne`     | Retorno de função          |
| `Exibir`      | Saída padrão               |
| `Verdadeiro`  | Literal lógico             |
| `Falso`       | Literal lógico             |
| `Nulo`        | Ausência de valor          |
| `não`         | Negação                    |
| `menor`       | Comparação                 |
| `maior`       | Comparação                 |
| `igual`       | Igualdade                  |
| `que`         | Complemento de comparação  |
| `de/do/da`    | Acesso a propriedades      |
