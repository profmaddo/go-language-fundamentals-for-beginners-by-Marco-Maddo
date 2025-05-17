# LECTURE-10: Maps in Go

## 📘 Code Overview (English)

This lesson introduces **maps** in Go – a key-value data structure used to store and retrieve data efficiently. Students will build on previous knowledge of input, loops, strings, and conditionals to work with maps dynamically.

---

## 🔹 What You Will Learn

- How to create, access, update, and delete map elements
- How to loop through a map using `range`
- How to read input and populate a map
- How to count word frequencies using a map
- How to store complex values like structs in maps
- How to extract and sort map keys
- How to compare maps and use boolean flags

---

## 🧠 Function-by-Function Explanation

Each function demonstrates a real use of Go maps with good practices.

- `createAndPrintMap()` – create a map and print key-value pairs
- `insertKeyValueToMap()` – insert values manually
- `readMapFromUserInput()` – use `bufio.Reader` to fill map
- `getValueFromMap()` – safely get map values
- `updateMapValue()` – update map values
- `deleteKeyFromMap()` – remove entries from map
- `checkKeyExists()` – check if key exists
- `loopThroughMap()` – loop using `range`
- `countWordsInSentence()` – map of word frequencies
- `mapWithStructValues()` – use a struct as map value
- `mapKeysToSlice()` – extract keys into slice
- `groupItemsByCategory()` – use slice as map value
- `sortMapKeys()` – sort map keys for display
- `compareTwoMaps()` – check equality manually
- `mapWithBoolFlags()` – use map for presence/visited flags

---

## 🧩 Integration with Previous Lessons

| Lecture       | Concept Used                  |
|---------------|-------------------------------|
| Lecture 3     | Text and string trimming       |
| Lecture 5     | Conditionals (`if`, `exists`)  |
| Lecture 6     | Loops and `range`              |
| Lecture 7     | Input with `bufio.NewReader`   |
| Lecture 8     | Slices as map values           |
| Lecture 9     | Arithmetic and assignment      |

---

# AULA-10: Mapas em Go (Dicionários)

## 📘 Visão Geral do Código (Português)

Esta aula apresenta o tipo `map` em Go – uma estrutura de chave-valor ideal para armazenar dados de forma rápida e organizada.  
A aula reforça habilidades anteriores ao aplicar entrada de dados, laços, strings e condicionais.

---

## 🔹 O que Você Vai Aprender

- Como criar, acessar, atualizar e remover elementos do mapa
- Como percorrer mapas com `range`
- Como preencher mapas com entrada do usuário
- Como contar frequência de palavras
- Como usar structs como valores de mapas
- Como extrair e ordenar chaves
- Como comparar dois mapas
- Como usar `map[string]bool` para simular flags

---

## 🧠 Explicação das Funções

Cada função demonstra um uso real dos mapas com boas práticas em Go.

- `createAndPrintMap()` – cria e imprime chave-valor
- `insertKeyValueToMap()` – insere pares manualmente
- `readMapFromUserInput()` – usa `bufio.Reader` para preencher o mapa
- `getValueFromMap()` – acessa valores com segurança
- `updateMapValue()` – atualiza valores
- `deleteKeyFromMap()` – remove entradas
- `checkKeyExists()` – verifica existência de chave
- `loopThroughMap()` – percorre o mapa com `range`
- `countWordsInSentence()` – conta palavras com mapa
- `mapWithStructValues()` – usa struct como valor
- `mapKeysToSlice()` – extrai chaves para slice
- `groupItemsByCategory()` – agrupa categorias
- `sortMapKeys()` – ordena chaves
- `compareTwoMaps()` – compara mapas manualmente
- `mapWithBoolFlags()` – flags booleanas em mapa

---

## 🧩 Integração com Aulas Anteriores

| Aula           | Conceitos Aplicados            |
|----------------|--------------------------------|
| Aula 3         | Strings e limpeza de texto     |
| Aula 5         | Condições com `if`             |
| Aula 6         | Laços com `range`              |
| Aula 7         | Entrada com `bufio.NewReader`  |
| Aula 8         | Slice como valor no mapa       |
| Aula 9         | Operadores e lógica condicional|

