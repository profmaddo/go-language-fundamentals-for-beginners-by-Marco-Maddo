# LECTURE-12: Advanced Functions in Go

## 📘 Code Overview (English)

This lesson introduces advanced function features in Go, including recursion, variadic functions, closures, pointers, anonymous functions, and functions as parameters. These are powerful tools for creating flexible and reusable code.

---

## 🔹 What You Will Learn

- How to define recursive functions
- How to use variadic parameters
- How to modify values using pointers
- How to pass functions as parameters and return them
- How to work with closures and anonymous functions
- How to apply higher-order logic like map, filter, and decorators

---

## 🧠 Function Descriptions

- `factorial(n)` – classic recursive function to calculate factorial.
- `sumVariadic(nums...)` – variadic function summing arbitrary values.
- `incrementByPointer(x *int)` – modifies the value by pointer.
- `swapPointers(a, b *int)` – swaps values using references.
- `applyFunction(a, b, op)` – accepts a binary function and applies it.
- `getMultiplier(factor)` – returns a closure (function that returns function).
- `printAnonymous()` – declares and calls an anonymous function.
- `sortDescending(slice)` – sorts a slice using an inline function.
- `mapFunction(slice, func)` – applies a function to all elements (like map).
- `filter(slice, func)` – filters slice elements using boolean function.
- `closureCounter()` – generates a counter function using closure.
- `isPalindrome(word)` – checks if a word is a palindrome (recursive).
- `calculateStats(numbers...)` – returns min and max values.
- `printValuesByReference(&slice)` – prints elements using pointer to slice.
- `runWithLog(func)` – decorator-style wrapper with log messages.

---

## 🧩 Integration with Previous Lessons

| Lecture       | Concepts Reinforced                 |
|---------------|--------------------------------------|
| Lecture 4     | Numeric operations and conditions    |
| Lecture 6     | Loops and counters                   |
| Lecture 8     | Slices as parameters                 |
| Lecture 9     | Logical evaluation                   |
| Lecture 11    | Parameters and return values         |

---

# AULA-12: Funções Avançadas em Go

## 📘 Visão Geral do Código (Português)

Esta aula apresenta funções avançadas da linguagem Go, incluindo recursão, parâmetros variádicos, closures, ponteiros, funções anônimas e funções como argumentos. São recursos poderosos para criar código mais flexível e reutilizável.

---

## 🔹 O que Você Vai Aprender

- Como criar funções recursivas
- Como usar parâmetros variádicos
- Como modificar valores usando ponteiros
- Como passar e retornar funções
- Como usar closures e funções anônimas
- Como aplicar lógica de ordem superior (map, filter, decorators)

---

## 🧠 Descrição das Funções

- `factorial(n)` – função recursiva para fatorial.
- `sumVariadic(nums...)` – soma valores com quantidade variável.
- `incrementByPointer(x *int)` – altera valor por referência.
- `swapPointers(a, b *int)` – troca valores com ponteiros.
- `applyFunction(a, b, op)` – aplica uma função como parâmetro.
- `getMultiplier(fator)` – retorna uma função multiplicadora (closure).
- `printAnonymous()` – declara e executa função anônima.
- `sortDescending(slice)` – ordena em ordem decrescente usando função inline.
- `mapFunction(slice, func)` – aplica uma função em cada item do slice.
- `filter(slice, func)` – filtra itens de acordo com condição booleana.
- `closureCounter()` – gera função contadora com estado interno.
- `isPalindrome(palavra)` – verifica se é palíndromo (recursiva).
- `calculateStats(numbers...)` – retorna o mínimo e máximo.
- `printValuesByReference(&slice)` – imprime slice por ponteiro.
- `runWithLog(func)` – executa função com logs antes e depois (estilo decorator).

---

## 🧩 Integração com Aulas Anteriores

| Aula           | Conceitos Reforçados               |
|----------------|------------------------------------|
| Aula 4         | Operações numéricas e condições    |
| Aula 6         | Laços e contadores                 |
| Aula 8         | Slices como parâmetros             |
| Aula 9         | Avaliação lógica                   |
| Aula 11        | Parâmetros e retornos              |
