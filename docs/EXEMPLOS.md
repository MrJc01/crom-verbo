# 📚 Galeria de Exemplos — Linguagem Verbo

Cada exemplo abaixo demonstra um conceito da linguagem com explicação passo a passo.

---

## 1. Olá, Mundo! (`ola_mundo.vrb`)

O programa mais básico em Verbo:

```
A mensagem é "Olá, Mundo!".
Exibir com (mensagem).
```

**Conceitos**: Constante (`A`), atribuição estática (`é`), saída padrão (`Exibir`).

**Saída esperada**:
```
Olá, Mundo!
```

---

## 2. Fibonacci (`fibonacci.vrb`)

Calcula os primeiros 10 números da sequência:

```
Um a está 0.
Um b está 1.
Um temp está 0.

Exibir com ("Sequência de Fibonacci:").
Exibir com (a).
Exibir com (b).

Repita 8 vezes:
    temp está a + b.
    a está b.
    b está temp.
    Exibir com (b).
```

**Conceitos**: Variáveis mutáveis (`Um/está`), loop (`Repita N vezes`), reatribuição, expressão aritmética.

**Saída esperada**:
```
Sequência de Fibonacci:
0
1
1
2
3
5
8
13
21
34
```

---

## 3. Calculadora (`calculadora.vrb`)

Funções aritméticas com parâmetros tipados:

```
Para Somar usando (a: Inteiro, b: Inteiro):
    Retorne a + b.

Para Subtrair usando (a: Inteiro, b: Inteiro):
    Retorne a - b.

Para Multiplicar usando (a: Inteiro, b: Inteiro):
    Retorne a * b.

A titulo é "Calculadora Verbo".
Exibir com (titulo).

Um resultado_soma é Somar com (10, 5).
Exibir com (resultado_soma).
```

**Conceitos**: Funções (`Para/usando`), tipos (`Inteiro`), retorno (`Retorne`), chamada (`com`).

---

## 4. Contagem Regressiva (`contador.vrb`)

Loop com decremento:

```
A titulo é "Contagem Regressiva".
Exibir com (titulo).

Um contador está 10.

Repita 10 vezes:
    Exibir com (contador).
    contador está contador - 1.

Exibir com ("Lançamento!").
```

**Conceitos**: Variável mutável, loop fixo, reatribuição com expressão aritmética.

---

## 5. Saudação (`saudacao.vrb`)

Funções simples e múltiplas chamadas:

```
A saudacao é "Bem-vindo ao sistema Verbo!".

Para Apresentar usando (nome: Texto):
    Exibir com (nome).

Exibir com (saudacao).
Apresentar com ("Brasil").
Apresentar com ("Portugal").
Apresentar com ("Moçambique").
```

**Conceitos**: Constante com texto, função com parâmetro `Texto`, múltiplas chamadas.

---

## Tabela Rápida de Conceitos

| Conceito          | Sintaxe Verbo              | Equivalente Go/Python     |
|-------------------|----------------------------|---------------------------|
| Constante         | `A x é 10.`               | `const x = 10` / `x = 10`|
| Variável          | `Um y está 0.`             | `y := 0` / `y = 0`       |
| Função            | `Para Func usando (p: T):` | `func Func(p T)` / `def` |
| Chamada           | `Func com (arg).`          | `Func(arg)` / `Func(arg)`|
| Print             | `Exibir com (x).`          | `fmt.Println(x)` / `print`|
| Loop fixo         | `Repita 5 vezes:`          | `for i := 0; i < 5` / `for i in range(5)`|
| Condicional       | `Se x for menor que y:`    | `if x < y` / `if x < y:` |
| Retorno           | `Retorne valor.`           | `return valor` / `return` |
| Reatribuição      | `x está x + 1.`            | `x = x + 1` / `x += 1`   |
