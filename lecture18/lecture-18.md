# LECTURE-18: Working with Files in Go

## 📘 Topic Overview

In this lesson, we explore how to interact with the file system in Go using packages such as `os`, `io`, `bufio`, and `filepath`. You’ll learn how to create, read, write, and explore files and directories safely and idiomatically in Go.

---

## 🔧 Function-by-Function Summary

### `ListCurrentDirectory()`
Lists files and folders in the current directory. Uses `os.ReadDir()`.

### `ListAllFilesRecursively(path)`
Traverses a given path recursively using `filepath.WalkDir()`.

### `PrintPlatformPathExamples()`
Prints valid path format examples for macOS, Linux, and Windows.

### `SaveLecturesToFile(filename)`
Writes a hardcoded list of lecture titles into a text file.

### `ReadFileLines(filepath)`
Opens and reads the content of a text file line by line.

### `WriteCustomTextToFile(text, filename)`
Writes a single string into a new file using `os.WriteFile`.

### `ReadFileInfo(path)`
Retrieves metadata about a file: name, size, and last modification time.

### `GroupFilesByExtension(path)`
Traverses a path and groups files by their file extension in a `map[string][]string`.

### `EnsureDirectory(path)`
Creates a directory if it doesn't exist using `os.MkdirAll`.

### `FileManager struct`
Encapsulates all behaviors for file and directory operations.

---

## 🔁 Integration with Previous Lessons

| Lecture       | Concepts Reused                            |
|---------------|---------------------------------------------|
| Lecture 8     | Use of slices for storing file names        |
| Lecture 10    | Use of maps to group file extensions        |
| Lecture 16    | POO: Struct `FileManager` + methods         |
| Lecture 17    | Error handling with all file operations     |

---

# AULA-18: Trabalhando com Arquivos em Go

## 📘 Visão Geral

Nesta aula você aprenderá a interagir com arquivos e diretórios em Go, utilizando os pacotes `os`, `io`, `bufio` e `filepath`. Vamos criar, ler, escrever e navegar em arquivos e pastas de maneira segura e idiomática em Go.

---

## 🔧 Resumo das Funções

### `ListCurrentDirectory()`
Lista os arquivos e pastas do diretório atual com `os.ReadDir()`.

### `ListAllFilesRecursively(path)`
Navega recursivamente pelo diretório usando `filepath.WalkDir()`.

### `PrintPlatformPathExamples()`
Exibe exemplos de caminhos para Linux/macOS e Windows.

### `SaveLecturesToFile(filename)`
Grava no arquivo `go_lectures.txt` o nome das aulas estudadas.

### `ReadFileLines(filepath)`
Lê um arquivo linha a linha com `bufio.Scanner`.

### `WriteCustomTextToFile(text, filename)`
Cria e escreve um texto simples em um arquivo com `os.WriteFile`.

### `ReadFileInfo(path)`
Lê as informações de um arquivo: nome, tamanho, modificação.

### `GroupFilesByExtension(path)`
Agrupa arquivos por extensão usando `map[string][]string`.

### `EnsureDirectory(path)`
Cria diretórios (e subdiretórios) se não existirem com `os.MkdirAll`.

### `Struct FileManager`
Aplica POO encapsulando todas as funções como métodos.

---

## 🔁 Integração com Aulas Anteriores

| Aula           | Conceitos Reforçados                       |
|----------------|---------------------------------------------|
| Aula 8         | Uso de slices para manipulação de listas   |
| Aula 10        | Uso de mapas (`map`)                       |
| Aula 16        | Programação Orientada a Objetos com métodos|
| Aula 17        | Tratamento de erros                        |

---

## ✅ Conclusão

Trabalhar com arquivos e diretórios é essencial em qualquer linguagem. Em Go, isso é feito com simplicidade e segurança usando funções que retornam erros, o que garante maior controle e robustez na manipulação de dados do sistema.

