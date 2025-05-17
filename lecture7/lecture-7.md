# LECTURE-7: Reading Input in Go

## 📘 Code Overview (English)

This lesson teaches how to get input from users using the terminal in Go. You’ll learn how to read integers, decimals, entire lines, and command-line arguments, and how to use that input in decisions.

---

## 🔹 What You Will Learn

- How to use `fmt.Scan` to read integers and floats
- How to read multiple values in one line
- How to read full lines using `bufio.NewReader`
- How to access command-line arguments with `os.Args`
- How to combine input and decision-making using `switch`

---

## 🧠 Function-by-Function Explanation

### `printLectureTitle()`
Prints the lecture title to help the user know the topic of this lesson.

### `readIntegerWithScan()`
Prompts the user to enter an integer. It uses `fmt.Scan()` to read the value and print it back.

### `readFloatWithScan()`
Asks the user for a decimal number (e.g. 3.14). Uses `fmt.Scan()` with a float64 variable and prints it formatted.

### `readMultipleValuesScan()`
Reads a string and an integer in a single line (e.g. "Alice 30"). Demonstrates how to capture multiple inputs at once.

### `readLineWithBufio()`
Uses `bufio.NewReader` to read a full sentence, including spaces. Useful when `Scan` is too limited.

### `readCommandLineArgs()`
Uses `os.Args` to read arguments passed during program execution (e.g. `go run main.go hello 123`).

### `readChoiceAndSwitch()`
Reads a number (1 to 3) and uses a `switch` statement to print a response based on user input.

---

## 📦 Explanation of Imported Packages

- `fmt`: Used to print and scan input/output in the terminal (`Println`, `Scan`, `Printf`).
- `bufio`: Provides buffered input methods. Here we use it to read full lines from the keyboard.
- `os`: Gives access to operating system features. In this example, we use `os.Stdin` for input and `os.Args` for command-line arguments.
- `strings`: Used to clean up or modify strings, e.g. trimming spaces using `TrimSpace`.

---

# AULA-7: Entrada de Dados em Go

## 📘 Visão Geral do Código (Português)

Esta aula ensina como capturar dados do usuário via terminal em Go. Você aprenderá a ler números inteiros, decimais, linhas completas e argumentos da linha de comando, além de como usar essas entradas em decisões.

---

## 🔹 O que Você Vai Aprender

- Como usar `fmt.Scan` para ler inteiros e floats
- Como ler múltiplos valores em uma única linha
- Como ler frases completas com `bufio.NewReader`
- Como acessar argumentos com `os.Args`
- Como usar entrada do usuário com `switch`

---

## 🧠 Explicação das Funções

### `printLectureTitle()`
Imprime o título da aula para organizar o conteúdo.

### `readIntegerWithScan()`
Pede ao usuário para digitar um número inteiro. Usa `fmt.Scan()` e imprime o valor.

### `readFloatWithScan()`
Solicita um número decimal (como 3.14). Usa `fmt.Scan()` com uma variável `float64` e imprime formatado.

### `readMultipleValuesScan()`
Lê dois valores em uma linha (por exemplo: "Alice 30"). Demonstra captura de múltiplos dados.

### `readLineWithBufio()`
Utiliza `bufio.NewReader` para ler uma linha inteira, inclusive com espaços.

### `readCommandLineArgs()`
Usa `os.Args` para capturar os argumentos passados ao rodar o programa (ex: `go run main.go teste`).

### `readChoiceAndSwitch()`
Lê um número entre 1 e 3 e usa um `switch` para exibir uma resposta baseada na escolha do usuário.

---

## 📦 Explicação dos Pacotes Importados

- `fmt`: Responsável por imprimir e ler dados do terminal com `Println`, `Scan` e `Printf`.
- `bufio`: Permite leitura mais controlada, como ler uma linha inteira do teclado.
- `os`: Permite acesso ao sistema operacional. Aqui usamos `os.Stdin` para entrada e `os.Args` para argumentos.
- `strings`: Permite manipulação de textos, como remover espaços com `TrimSpace`.
