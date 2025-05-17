# LECTURE-8: Arrays and Slices in Go

## 📘 Code Overview (English)

This lesson introduces **arrays and slices** in Go — two core tools for storing and managing collections of data. Arrays have a fixed size, while slices are dynamic and more flexible.

Students will reuse knowledge from previous lessons (loops, input, conditionals) to build practical examples.

---

## 🔹 What You Will Learn

- The difference between arrays and slices
- How to iterate over arrays and slices using loops
- How to collect user input and store it in a list
- How to grow and modify slices using `append()`
- How to search and remove elements in slices
- How to inspect slice capacity and understand memory behavior

---

## 🧠 Function-by-Function Explanation

### `printLectureTitle()`
Prints the title of the lecture. A good habit to organize the console output.

---

### `printFixedArray()`
- Declares a fixed array with 5 integers.
- Uses a `for` loop to print each element by index.
- Demonstrates the basic use of arrays with known size.

✅ **Integration with Lecture 6 (loops)**: reinforces iteration using `for` with indexes.

---

### `readAndFillArray()`
- Uses `fmt.Scan` in a loop to fill an array from user input.
- Shows how to collect structured numeric data.

✅ **Integration with Lecture 7 (input)**: uses terminal input with `Scan`.

---

### `calculateSumArray()`
- Sums all elements of an array using a `for range` loop.
- Displays the result.

✅ **Integration with Lecture 6**: `range` syntax for clean looping.
✅ **Math operation from Lecture 4**: simple addition on integers.

---

### `createAndPrintSlice()`
- Declares a slice of strings and prints its elements using `range`.
- Shows how slices behave like flexible lists.

✅ **New concept**: slice declaration (`[]string{}`) vs array (`[n]string`).

---

### `appendToSlice()`
- Demonstrates how to grow a slice using `append()`.
- Elements are added dynamically.

✅ **Practical tip**: slices are preferred over arrays in real-world programs.

---

### `sliceFromArray()`
- Converts a portion of an array into a slice using slicing syntax.
- Highlights the shared memory behavior between arrays and slices.

✅ **Foundational slice usage**: syntax like `arr[1:4]`.

---

### `readWordsToSlice()`
- Uses `bufio.Reader` to read full words from user input.
- Appends each word to a slice.
- Demonstrates collecting dynamic data from input.

✅ **Integration with Lecture 7**: buffered input for strings.
✅ **Slice manipulation with `append()`**.

---

### `searchInSlice()`
- Asks the user for a number and checks if it exists in a slice.
- Uses a loop and boolean flag to determine presence.

✅ **Integration with Lecture 5 (if/else)**: applies conditionals.
✅ **Loop from Lecture 6**, `Scan` from Lecture 7.

---

### `removeItemFromSlice()`
- Removes an item from a slice by index using slicing and `append()`.
- Shows how to rebuild a slice without a specific element.

✅ **Core Go trick**: `slice = append(slice[:i], slice[i+1:]...)`.

---

### `sliceCapacityGrowth()`
- Appends integers to a slice and prints `len()` and `cap()` after each append.
- Demonstrates how Go automatically increases slice capacity.

✅ **New concept**: understanding memory behavior (`len`, `cap`).

---

## 🧩 Recap: Integration with Previous Lessons

| Previous Lecture | Concepts Applied                          |
|------------------|-------------------------------------------|
| Lecture 4        | Numeric operations, variable types        |
| Lecture 5        | Conditional logic (`if/else`)             |
| Lecture 6        | Loop structures (`for`, `range`)          |
| Lecture 7        | Reading input (`Scan`, `bufio.Reader`)    |

---

# AULA-8: Arrays e Slices em Go

## 📘 Visão Geral do Código (Português)

Esta aula apresenta os **arrays e slices**, estruturas fundamentais para armazenar coleções de dados. Arrays possuem tamanho fixo. Slices são listas dinâmicas e mais flexíveis.

Nesta aula, o aluno revisa conceitos anteriores e aprende a manipular listas com entrada, laços, busca e remoção.

---

## 🔹 O que Você Vai Aprender

- Diferença entre arrays (fixos) e slices (dinâmicos)
- Como percorrer arrays e slices com `for` e `range`
- Como coletar dados do usuário e armazenar em lista
- Como adicionar elementos com `append()`
- Como procurar e remover valores em slices
- Como verificar tamanho e capacidade com `len()` e `cap()`

---

## 🧠 Explicação das Funções

### `printLectureTitle()`
Imprime o título da aula, organizando a saída no terminal.

---

### `printFixedArray()`
- Declara um array fixo de 5 inteiros.
- Usa `for` com índice para imprimir os valores.

✅ **Integra com Lecture 6 (laços)**: uso de `for` com índice.

---

### `readAndFillArray()`
- Usa `fmt.Scan` para ler 3 números do usuário e armazenar no array.

✅ **Integra com Lecture 7 (entrada)**: leitura com `Scan`.

---

### `calculateSumArray()`
- Soma os valores do array com `range` e exibe o resultado.

✅ **Integra com Lecture 4 (operações)** e **Lecture 6 (range)**.

---

### `createAndPrintSlice()`
- Cria um slice de strings e imprime com `range`.

✅ **Novo conceito**: slices como listas mutáveis.

---

### `appendToSlice()`
- Mostra como adicionar novos elementos a um slice com `append()`.

✅ **Slice dinâmico na prática**.

---

### `sliceFromArray()`
- Gera um slice a partir de parte de um array (`arr[1:4]`).

✅ **Importante para recortes e reutilização de dados**.

---

### `readWordsToSlice()`
- Usa `bufio` para ler 3 palavras e salvar num slice.

✅ **Integra com Lecture 7**, entrada com `bufio`.

---

### `searchInSlice()`
- Verifica se um número informado está presente no slice.

✅ **Integra com Lecture 5 (condicional)**, 6 (laço), 7 (entrada).

---

### `removeItemFromSlice()`
- Remove o item da posição 2 de um slice de strings.

✅ **Demonstra operação comum de limpeza de dados**.

---

### `sliceCapacityGrowth()`
- Mostra como a capacidade de um slice cresce automaticamente.

✅ **Conceito avançado útil para otimização de performance**.

