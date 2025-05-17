# LECTURE-6: Loops and Repetition in Go

## 📘 Code Overview (English)

This lesson introduces different types of loops in Go and how to use them to repeat instructions.

### 🔹 What You Will Learn

- How to repeat actions using a `for` loop with a counter
- How to use condition-based loops
- How to use `range` to iterate over slices and strings
- How to nest loops to create multiplication tables

---

## 🔍 Function Breakdown

### `printLectureTitle()`

- Prints the title of the current lecture.
- Helps the user identify the topic of the lesson.

### `forCounterLoop()`

- Demonstrates a classic `for` loop with a counter variable.
- It counts from 1 to 5 and prints each value.
- The loop uses three parts: `i := 1` (start), `i <= 5` (condition), and `i++` (increment).

### `forConditionalLoop()`

- A `for` loop using only a condition.
- Starts with `x = 1` and loops while `x <= 5`.
- Demonstrates a while-like behavior using Go's `for`.

### `forRangeSlice()`

- Uses a `range` loop to iterate over a list (slice) of numbers.
- The loop provides both the index (position) and the value at each step.
- It prints: `Index 0: Value 10`, etc.

### `forRangeString()`

- Iterates over each character in a string.
- `range` returns the position and the character (as a rune).
- It prints characters like: `Index 0: Character G`, `Index 1: Character o`, etc.

### `nestedLoopExample()`

- Demonstrates how to nest one loop inside another.
- Creates a multiplication table from 1 to 3.
- For example, it prints: `1 x 1 = 1`, `1 x 2 = 2`, ..., `3 x 3 = 9`.

---

# AULA-6: Laços e Repetições em Go

## 📘 Visão Geral do Código (Português)

Esta aula apresenta diferentes formas de repetir ações em Go usando laços (`loops`).

### 🔹 O que você vai aprender

- Como repetir ações com um `for` com contador
- Como usar `for` com condição
- Como usar `range` para percorrer listas e strings
- Como criar tabelas de multiplicação com laços aninhados

---

## 🔍 Explicação das Funções

### `printLectureTitle()`

- Imprime o título da aula no terminal.
- Ajuda o aluno a saber qual tema está sendo estudado.

### `forCounterLoop()`

- Demonstra o uso clássico do laço `for` com contador.
- Conta de 1 até 5 e imprime os valores.
- Usa três partes: início (`i := 1`), condição (`i <= 5`) e incremento (`i++`).

### `forConditionalLoop()`

- Usa um `for` apenas com condição, semelhante ao `while`.
- Começa com `x = 1` e executa enquanto `x <= 5`.

### `forRangeSlice()`

- Usa `range` para percorrer uma lista (slice) de números.
- O `range` retorna o índice (posição) e o valor.
- Exemplo: `Index 0: Value 10`

### `forRangeString()`

- Percorre cada letra da string `"GoLang"`.
- O `range` retorna a posição e o caractere.
- Exemplo: `Index 0: Character G`

### `nestedLoopExample()`

- Mostra como criar um laço dentro de outro.
- Imprime a tabuada de multiplicação de 1 a 3.
- Exemplo: `2 x 3 = 6`

