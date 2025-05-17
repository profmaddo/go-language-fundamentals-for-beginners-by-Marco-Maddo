# LECTURE-9: Operators in Go

## 📘 Code Overview (English)

This lesson focuses on arithmetic, comparison, logical, assignment, and precedence operators in Go.  
Operators allow you to create dynamic behavior by comparing values, performing math, and combining conditions.

By this point, students have learned how to read input, use control structures, and work with data types — this lecture connects them through practical operator logic.

---

## 🔹 What You Will Learn

- How to use `+`, `-`, `*`, `/`, `%` for calculations
- How to use comparison operators `==`, `!=`, `>`, `<`, `>=`, `<=`
- How to evaluate logic with `&&`, `||`, and `!`
- How to combine multiple conditions
- How operator precedence affects results
- How to use compound assignment (`+=`, `-=`, etc.)
- How to display logic with `fmt.Printf`

---

## 🧠 Function-by-Function Explanation

### `printLectureTitle()`
Prints the lesson title. Helps organize output.

---

### `basicArithmeticOperations()`
Performs basic math between two numbers.

✅ **Related to Lecture 4** (numeric types)  
✅ **Builds on Lecture 6** (`Printf`)  
✅ **Uses `fmt.Printf()` formatting for clarity**

---

### `incrementDecrementOperators()`
Demonstrates the `++` and `--` operators to increase or decrease an integer value.

✅ **Used later in loops (Lecture 6)**

---

### `comparisonOperators()`
Compares two integers using equality and inequality operators.

✅ **Prepares logic for `if/else` (Lecture 5)**

---

### `logicalOperators()`
Shows how logical expressions (`&&`, `||`, `!`) work using booleans.

✅ **Core concept for conditional structures (Lecture 5)**

---

### `arithmeticWithInput(a, b int)`
Accepts two integers and prints arithmetic results.

✅ **Reuses `Printf()` (Lecture 3, 7)**  
✅ **Could easily be adapted for `Scan` input**

---

### `comparisonWithInput(a, b int)`
Compares two numbers and prints results with `Printf()` using `%v`.

✅ **Reinforces booleans as values**

---

### `logicalEvaluationWithFlags(isAdult, hasLicense)`
Combines two boolean flags to control access (drive or not).

✅ **Builds real-world logic from previous lectures**

---

### `boolExpressionsInPrintf(n1, n2 int)`
Prints two integers and whether they are equal, using inline expressions.

✅ **Combines Lecture 5 (conditionals) with `Printf()` format (Lecture 3)**

---

### `operatorPrecedenceExample()`
Shows how Go follows operator precedence in math expressions.

✅ **Important for accurate calculations**

---

### `assignmentOperators()`
Applies combined operations like `+=`, `-=`, `*=`, `/=`, `%=`.

✅ **Saves typing and makes expressions more concise**

---

### `combinedLogicalComparison(age, hasTicket)`
Uses `if` with a combination of comparison and logical conditions.

✅ **Connects multiple concepts into one practical decision**

---

## 🧩 Recap: Integration with Previous Lessons

| Previous Lecture | Concepts Applied                        |
|------------------|-----------------------------------------|
| Lecture 3        | Formatted output with `Printf`          |
| Lecture 4        | Arithmetic with numeric types           |
| Lecture 5        | Conditionals (`if`, `switch`)           |
| Lecture 6        | Loop structures and increment/decrement |
| Lecture 7        | Input (used here via parameters)        |

---

# AULA-9: Operadores em Go

## 📘 Visão Geral do Código (Português)

Esta aula apresenta os principais operadores em Go: aritméticos, de comparação, lógicos, de atribuição e de precedência.  
Os operadores permitem construir lógica, realizar cálculos e tomar decisões baseadas em expressões.

O aluno já aprendeu entrada de dados, controle de fluxo e tipos numéricos — agora ele aplica isso de forma combinada.

---

## 🔹 O que Você Vai Aprender

- Como usar `+`, `-`, `*`, `/`, `%` em cálculos
- Como comparar valores com `==`, `!=`, `>`, `<`, `>=`, `<=`
- Como usar expressões com `&&`, `||`, `!`
- Como combinar condições com `if`
- Como a ordem dos operadores afeta o resultado
- Como usar operadores compostos `+=`, `-=`, etc.
- Como exibir valores booleanos no `Printf`

---

## 🧠 Explicação das Funções

### `printLectureTitle()`
Imprime o título da aula no terminal.

---

### `basicArithmeticOperations()`
Faz operações básicas com dois números e imprime os resultados com `Printf`.

✅ **Relaciona com Lecture 4 e Lecture 3**

---

### `incrementDecrementOperators()`
Mostra o uso de `++` e `--`, úteis em laços e contadores.

✅ **Usado em loops (Lecture 6)**

---

### `comparisonOperators()`
Compara dois inteiros com vários operadores e mostra os resultados.

✅ **Base para decisões com `if` (Lecture 5)**

---

### `logicalOperators()`
Avalia expressões booleanas simples com `&&`, `||`, `!`.

✅ **Preparação para lógica combinada (Lecture 5)**

---

### `arithmeticWithInput(a, b int)`
Mostra os resultados de operações aritméticas entre dois números recebidos como parâmetros.

✅ **Simula uso de `Scan()` (Lecture 7)**

---

### `comparisonWithInput(a, b int)`
Compara dois valores recebidos e mostra as comparações formatadas.

✅ **Exibe booleanos no terminal (`%v`)**

---

### `logicalEvaluationWithFlags(isAdult, hasLicense)`
Verifica se duas condições booleanas são verdadeiras para permitir ou não uma ação.

✅ **Exemplo real de lógica combinada**

---

### `boolExpressionsInPrintf(n1, n2 int)`
Usa `Printf` para exibir dois números e se são iguais.

✅ **Integra `Printf`, booleanos e lógica**

---

### `operatorPrecedenceExample()`
Demonstra como a ordem dos operadores muda o resultado da expressão.

✅ **Conceito importante para cálculos corretos**

---

### `assignmentOperators()`
Aplica operações compostas a uma variável (`+=`, `-=`, etc.).

✅ **Ajuda a simplificar expressões matemáticas**

---

### `combinedLogicalComparison(age, hasTicket)`
Avalia se a pessoa pode entrar, com base em idade e ticket.

✅ **Integra comparação e lógica com `if`**

