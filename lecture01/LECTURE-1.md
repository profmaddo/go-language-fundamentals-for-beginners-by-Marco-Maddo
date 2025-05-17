# LECTURE-1: Go Language Fundamentals (Golang)

## 📦 Code Structure

This Go example demonstrates the basic foundations of the language in a simple and educational way.  
The code is structured as follows:

- `package main`: Declares the main package, required to compile an executable.
- `import "fmt"`: Imports the `fmt` package, which provides functions for formatted I/O, like `Println`.
- `func main()`: The program's entry point, where execution begins.

## 🧠 Created Functions

The program defines two functions outside `main()` and calls them within `main()` for better organization:

### 1. `helloworld()`

- Uses `fmt.Println` to print a message:  
  `Hello, World! I'm GO... Let's go?`

### 2. `variableinteger()`

- Declares a variable named `year` of type `int` with value `2025`.
- Uses `fmt.Println` to print the variable:
  - Just the value: `2025`
  - With text: `Year: 2025`
- Uses `fmt.Sprintln(...)`, which formats a string but does **not** print it.  
  It should use `fmt.Printf("The Year is %d\n", year)` instead.

## 💡 Key Concepts Taught

- Organizing code into functions
- Declaring typed variables
- Printing with `fmt.Println`
- Using the standard `fmt` package
- Using `main()` as the entry point

## 🚨 Important Tip.

`fmt.Sprintln(...)` returns a formatted string but does not print it.  
To properly print with format like `%d`, use `fmt.Printf(...)` or store the result of `Sprintln` in a variable and print that.

**Correct example:**

```go
fmt.Printf("The Year is %d\n", year)
```

---

# AULA-1: Fundamentos da Linguagem Go (Golang)

## 📦 Estrutura do Código

Este exemplo em Go mostra os fundamentos básicos da linguagem de maneira simples e didática.  
A estrutura do código está organizada da seguinte forma:

- `package main`: Indica que este é o pacote principal do programa, obrigatório para gerar um executável.
- `import "fmt"`: Importa o pacote `fmt`, que fornece funções para entrada e saída formatadas, como `Println`.
- `func main()`: Ponto de entrada do programa, onde tudo começa a ser executado.

## 🧠 Funções Criadas

O programa define duas funções fora da `main()` e as chama dentro da `main()`. Isso ajuda a organizar o código:

### 1. `helloworld()`

- Usa `fmt.Println` para imprimir uma mensagem na tela:  
  `Hello, World! I'm GO... Let's go?`

### 2. `variableinteger()`

- Declara uma variável chamada `year` do tipo `int` com valor `2025`.
- Usa `fmt.Println` para imprimir o valor da variável de duas formas:
  - Apenas o valor: `2025`
  - Texto mais o valor: `Year: 2025`
- Usa incorretamente `fmt.Sprintln(...)`, que cria uma string mas não a imprime.  
  Deveria ser `fmt.Printf("The Year is %d\n", year)` para formatar corretamente.

## 💡 Conceitos Ensinados

- Organização em funções
- Declaração de variáveis com tipo explícito
- Impressão de valores com `fmt.Println`
- Uso de pacote padrão (`fmt`)
- Função principal `main()` como ponto de partida

## 🚨 Dica Importante

A linha com `fmt.Sprintln(...)` não imprime nada porque essa função retorna uma string formatada.  
Para imprimir usando formatação com `%d`, use `fmt.Printf(...)` ou atribua o resultado de `Sprintln` a uma variável e depois imprima.

**Exemplo correto:**

```go
fmt.Printf("The Year is %d\n", year)
```
