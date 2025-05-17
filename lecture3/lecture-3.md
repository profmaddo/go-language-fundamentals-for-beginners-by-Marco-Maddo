# LECTURE-3: String Manipulation in Go

## 📘 Code Overview (English)

This Go program demonstrates how to work with strings using the `strings` package and prints output using `fmt`.

### 🔹 Code Highlights.

- Uses a constant base string: `"GO lang Lecture 3 about Strings"`
- Demonstrates how to transform, inspect, and manipulate string content

### 🔹 Functions Explained

| Function Name           | Purpose                                                             |
|-------------------------|---------------------------------------------------------------------|
| `printLectureTitle()`   | Prints the lecture header                                           |
| `printOriginal()`       | Displays the original string                                        |
| `printUpperCase()`      | Converts string to uppercase                                        |
| `printLowerCase()`      | Converts string to lowercase                                        |
| `printLength()`         | Shows the total number of characters                                |
| `printContainsWord()`   | Checks if a word is contained in the string                         |
| `printStartsWith()`     | Checks if the string starts with a given prefix                     |
| `printEndsWith()`       | Checks if the string ends with a given suffix                       |
| `printWordCount()`      | Counts the number of words using `strings.Fields()`                 |
| `printReplaceWord()`    | Replaces a word in the string                                       |
| `printSplitWords()`     | Splits string into a slice of words                                 |
| `printJoinWords()`      | Joins a slice of words into one string                              |
| `printIndexOfWord()`    | Finds the index of a word                                           |
| `printTrimSpaces()`     | Removes leading and trailing spaces from another string             |
| `printCompareStrings()` | Compares two strings, both case-sensitive and case-insensitive      |

### 💡 Tips

- `strings.Fields()` is useful to tokenize the string by space.
- Use `fmt.Printf` for formatted output like indexes and comparisons.
- String comparison with `Compare()` returns -1, 0, or 1.
- Use `EqualFold()` for case-insensitive equality check.

---

# AULA-3: Manipulação de Strings em Go

## 📘 Visão Geral do Código (Português)

Este programa em Go mostra como trabalhar com strings utilizando o pacote `strings` e a função `fmt.Println` para saída no terminal.

### 🔹 Destaques do Código

- Utiliza uma constante base de texto: `"GO lang Lecture 3 about Strings"`
- Demonstra como transformar, inspecionar e manipular strings

### 🔹 Funções Explicadas

| Nome da Função            | Finalidade                                                         |
|---------------------------|---------------------------------------------------------------------|
| `printLectureTitle()`     | Imprime o cabeçalho da aula                                        |
| `printOriginal()`         | Mostra a string original                                           |
| `printUpperCase()`        | Converte a string para letras maiúsculas                           |
| `printLowerCase()`        | Converte a string para letras minúsculas                           |
| `printLength()`           | Mostra o número total de caracteres                                |
| `printContainsWord()`     | Verifica se a string contém uma determinada palavra                |
| `printStartsWith()`       | Verifica se a string começa com um prefixo                         |
| `printEndsWith()`         | Verifica se a string termina com um sufixo                         |
| `printWordCount()`        | Conta quantas palavras existem usando `strings.Fields()`           |
| `printReplaceWord()`      | Substitui uma palavra na string                                    |
| `printSplitWords()`       | Divide a string em palavras (slice de strings)                     |
| `printJoinWords()`        | Junta palavras de um slice em uma string única                     |
| `printIndexOfWord()`      | Encontra a posição de uma palavra                                  |
| `printTrimSpaces()`       | Remove espaços do início e fim de outra string                     |
| `printCompareStrings()`   | Compara duas strings (sensível e insensível a maiúsculas/minúsculas)|

### 💡 Dicas

- `strings.Fields()` ajuda a dividir a string por espaços em branco.
- `fmt.Printf` permite saída formatada para índices e comparações.
- A função `Compare()` retorna -1, 0 ou 1 conforme a ordem das strings.
- Use `EqualFold()` para comparar ignorando letras maiúsculas ou minúsculas.
