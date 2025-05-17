# LECTURE-13: Pointers in Go

## 📘 Code Overview (English)

This lesson introduces the concept of **pointers** in Go, using `*` and `&`. Pointers allow passing and modifying values by reference, working directly with memory addresses.

---

## 🔹 What You Will Learn

- What `*` (dereferencing) and `&` (address-of) mean in Go
- How to pass variables by reference using pointers
- How to update data from a function via pointer
- How to use pointers with slices and structs
- Differences between pass-by-value and pass-by-reference

---

## 🧠 Function Descriptions

- `printAddress(x)` – prints memory address using `&x`
- `modifyByValue(x)` – changes value locally (no effect)
- `modifyByPointer(&x)` – modifies original value by pointer
- `swapValuesPointer(&a, &b)` – swaps values using pointers
- `resetToZero(ptr)` – sets value to zero by pointer
- `doubleValue(ptr)` – doubles the value by reference
- `compareAddresses(&x, &y)` – compares memory addresses
- `copyByReference(src, dst)` – copies value via pointer
- `sumSliceWithPointer(&slice)` – sums elements of a slice pointer
- `appendToSlicePointer(&slice)` – appends to slice using pointer
- `createPointerReturn(value)` – returns pointer to value
- `getValueFromPointer(ptr)` – reads from pointer
- `incrementAll(slice)` – updates slice elements (pass by reference)
- `swapPointers(&ptr1, &ptr2)` – swaps pointers to pointers
- `updateStructField(&user)` – updates struct field using pointer

---

## 🧩 Integration with Previous Lessons

| Lecture       | Concepts Reinforced              |
|---------------|----------------------------------|
| Lecture 4     | Numeric types and math           |
| Lecture 6     | Loop logic with slices           |
| Lecture 8     | Slices and reference manipulation|
| Lecture 11    | Function parameters and return   |
| Lecture 12    | Pointer-based logic reuse        |

---

# AULA-13: Ponteiros em Go

## 📘 Visão Geral do Código (Português)

Esta aula apresenta o uso de **ponteiros (`*` e `&`)** na linguagem Go. Ponteiros permitem que funções acessem e modifiquem valores originais diretamente na memória.

---

## 🔹 O que Você Vai Aprender

- O que significam `*` (acesso ao valor) e `&` (endereço)
- Como passar variáveis por referência com ponteiros
- Como modificar dados a partir de uma função
- Como usar ponteiros com slices e structs
- Diferença entre passar por valor e por referência

---

## 🧠 Descrição das Funções

- `printAddress(x)` – imprime o endereço de memória de uma variável
- `modifyByValue(x)` – altera valor localmente (sem efeito externo)
- `modifyByPointer(&x)` – altera valor original via ponteiro
- `swapValuesPointer(&a, &b)` – troca dois valores usando ponteiros
- `resetToZero(ptr)` – zera o valor apontado
- `doubleValue(ptr)` – dobra o valor via referência
- `compareAddresses(&x, &y)` – compara endereços de memória
- `copyByReference(src, dst)` – copia valor entre ponteiros
- `sumSliceWithPointer(&slice)` – soma itens de um slice via ponteiro
- `appendToSlicePointer(&slice)` – adiciona ao slice por referência
- `createPointerReturn(value)` – retorna ponteiro para valor criado
- `getValueFromPointer(ptr)` – lê valor via ponteiro
- `incrementAll(slice)` – incrementa todos os valores do slice
- `swapPointers(&ptr1, &ptr2)` – troca ponteiros entre si
- `updateStructField(&user)` – modifica campo de struct via ponteiro

---

## 🧩 Integração com Aulas Anteriores

| Aula           | Conceitos Reforçados            |
|----------------|----------------------------------|
| Aula 4         | Tipos numéricos e operações      |
| Aula 6         | Laços com slices                 |
| Aula 8         | Slices e referências             |
| Aula 11        | Parâmetros e retorno             |
| Aula 12        | Uso avançado de ponteiros        |
