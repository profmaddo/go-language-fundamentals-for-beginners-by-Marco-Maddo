# LECTURE-17: Error Handling in Go

## 📘 Introduction: What is `error`?

In Go, errors are not handled with `try/catch`. Instead, functions return an `error` value as a second return when something might go wrong. This is the idiomatic and preferred way to write safe and readable code in Go.

---

## 🧠 Overview of Functions and Error Handling

(Refer to each lecture for context)

### LECTURE 3 – String Handling  
- `toUpper()` – error if empty string  
- `containsWord()` – error if search word is empty  
- `replaceWord()` – error if `old` string is empty  
- `indexOfWord()` – error if word is not found

### LECTURE 4 – Numeric Operations  
- `divide()` – error on division by zero  
- `subtractUint16()` – error if result would be negative  
- `areaRectangle()` – error if dimensions are negative  
- `validateComplex()` – error if `NaN`

### LECTURE 7 – Input Validation  
- `validateAgeInput()` – error for negative age  
- `validateOption()` – error for invalid options

### LECTURE 8 – Arrays and Slices  
- `sumSlice()` – error if nil  
- `searchInSlice()` – error if not found  
- `removeItemByIndex()` – error for invalid index

### LECTURE 10 – Maps  
- `getFromMap()` – error if key not found  
- `updateInMap()` – error if key missing  
- `deleteFromMap()` – error if key not found

### LECTURE 11 – Function Returns  
- `repeatWord()` – error for empty string or invalid repeat count  
- `getFirstAndLast()` – error for empty slice

---

## 🔁 Integration with Previous Lessons

| Lecture        | Concepts Reinforced                               |
|----------------|----------------------------------------------------|
| Lecture 3      | Strings and index handling                        |
| Lecture 4      | Arithmetic validation, safe numeric operations    |
| Lecture 7      | Validating user input                             |
| Lecture 8      | Working with slices and indexing                  |
| Lecture 10     | Map access and key checks                         |
| Lecture 11     | Return values, slice boundaries                   |

---

# AULA-17: Tratamento de Erros em Go

## 📘 Introdução: O que é `error`?

Em Go, não usamos `try/catch` como em outras linguagens. Em vez disso, uma função que pode falhar deve retornar um segundo valor do tipo `error`. O tratamento é feito com `if err != nil`, tornando o código mais claro e seguro.

---

## 🧠 Visão Geral das Funções com Tratamento de Erros

(Consulte as lectures originais para entender o contexto)

### LECTURE 3 – Manipulação de Strings  
- `toUpper()` – erro se string vazia  
- `containsWord()` – erro se palavra vazia  
- `replaceWord()` – erro se `old` estiver vazio  
- `indexOfWord()` – erro se não encontrada

### LECTURE 4 – Operações Numéricas  
- `divide()` – erro ao dividir por zero  
- `subtractUint16()` – erro se resultado for negativo  
- `areaRectangle()` – erro se dimensões inválidas  
- `validateComplex()` – erro se real ou imag forem `NaN`

### LECTURE 7 – Validação de Entrada  
- `validateAgeInput()` – erro se idade < 0  
- `validateOption()` – erro se opção inválida

### LECTURE 8 – Arrays e Slices  
- `sumSlice()` – erro se slice for nil  
- `searchInSlice()` – erro se valor não encontrado  
- `removeItemByIndex()` – erro se índice inválido

### LECTURE 10 – Mapas  
- `getFromMap()` – erro se chave não existir  
- `updateInMap()` – erro se chave não existir  
- `deleteFromMap()` – erro se chave ausente

### LECTURE 11 – Retorno de Funções  
- `repeatWord()` – erro se string vazia ou repetição inválida  
- `getFirstAndLast()` – erro se slice vazio

---

## 🔁 Integração com Aulas Anteriores

| Aula           | Conceitos Reforçados                            |
|----------------|--------------------------------------------------|
| Aula 3         | Manipulação e validação de strings              |
| Aula 4         | Operações aritméticas e proteção contra falhas  |
| Aula 7         | Validação de entrada e limites                  |
| Aula 8         | Manipulação segura de slices                    |
| Aula 10        | Leitura e escrita seguras com mapas             |
| Aula 11        | Boas práticas de retorno de funções             |

---

## ✅ Conclusão

Ao aplicar `error`, o código Go se torna mais robusto, legível e confiável. Ao longo dessa lecture, reutilizamos funções de aulas anteriores e aplicamos o padrão idiomático para validar condições antes de prosseguir. Esse padrão será essencial para todos os projetos futuros em Go.

