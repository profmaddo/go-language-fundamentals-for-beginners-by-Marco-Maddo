# LECTURE-20: Using SQLite in Go with POO Concepts

## 📘 Topic Overview

This lecture introduces how to integrate the Go programming language with SQLite, a lightweight relational database. We focus on saving structured data based on concepts already studied in OOP (Object-Oriented Programming). In this version, we manage a single table — `Person` — and implement full CRUD (Create, Read, Update, Delete) functionality through a terminal menu.

---

## 🔧 Function-by-Function Explanation

This program begins by connecting to a SQLite database named `go_database.db`. The `connectDB()` function opens the connection and handles possible errors. If successful, the program creates the `Person` table using `createPersonTable()`, which executes a `CREATE TABLE IF NOT EXISTS` SQL command to define columns `id`, `name`, and `age`.

The `insertPerson()` function prepares and executes an `INSERT INTO` SQL statement to add new records to the table. Similarly, `listPersons()` uses `SELECT` to fetch all rows and display them with a simple terminal output.

To update an existing record, `updatePersonByID()` executes an `UPDATE` query based on the person's `id`, and `deletePersonByID()` runs a `DELETE FROM` query to remove a record by ID. For search functionality, `findPersonByID()` uses `QueryRow()` to fetch a single record matching the given ID and reports if no record is found.

All these operations are wrapped inside the `menuPerson()` function, which provides an interactive text menu for the user. The `bufio.NewReader()` is used to read user input, which is cleaned with `strings.TrimSpace()` to avoid newline issues. The `strconv.Atoi()` function is used to convert user-provided strings into integers where necessary.

Each CRUD operation displays a success or failure message immediately after execution. The design ensures that the program always returns to the menu until the user chooses to exit by typing `0`.

---

## 🔁 Integration with Previous Lessons

| Lecture       | Concepts Reinforced                      |
|---------------|-------------------------------------------|
| Lecture 11    | Functions with parameters and returns     |
| Lecture 16    | OOP concepts: working with structs        |
| Lecture 17    | Error handling pattern with `error`       |
| Lecture 18    | Terminal input/output and menus           |

---

# AULA-20: Usando SQLite em Go com Conceitos de POO

## 📘 Visão Geral

Nesta aula, integramos o banco de dados relacional SQLite com programas escritos em Go, aplicando os conceitos de POO já vistos. Neste exemplo, manipulamos uma única entidade — `Person` — e implementamos as funcionalidades de CRUD (Create, Read, Update, Delete) acessíveis por um menu interativo no terminal.

---

## 🔧 Explicação das Funções

O programa inicia conectando-se ao banco de dados `go_database.db` com a função `connectDB()`, que retorna um ponteiro para `sql.DB`. A função `createPersonTable()` cria a tabela `Person` com os campos `id`, `name` e `age` utilizando SQL. O campo `id` é do tipo autoincremento e chave primária.

Para inserir dados, a função `insertPerson()` prepara uma instrução SQL com parâmetros e executa a inclusão. Já a função `listPersons()` busca todos os registros com `SELECT` e imprime cada linha no terminal.

A função `updatePersonByID()` atualiza um registro baseado no `id` e os novos valores informados. Para deletar, usamos `deletePersonByID()`, que executa um `DELETE FROM` seguro. A função `findPersonByID()` busca um único registro por ID e trata a situação quando nenhum dado é encontrado.

Todas essas funções estão centralizadas na função `menuPerson()`, que exibe um menu textual e interpreta os comandos do usuário. A entrada de dados é feita com `bufio.NewReader()` e limpa com `strings.TrimSpace()` para evitar problemas com quebras de linha. A conversão de strings para inteiros é feita com `strconv.Atoi()`.

Cada operação mostra imediatamente uma mensagem indicando sucesso ou falha. O menu é exibido continuamente até o usuário digitar `0` para sair, garantindo uma boa experiência interativa no terminal.

---

## 🔁 Integração com Aulas Anteriores

| Aula           | Conceitos Reforçados                        |
|----------------|----------------------------------------------|
| Aula 11        | Funções com parâmetros e retorno             |
| Aula 16        | Programação orientada a objetos (POO)        |
| Aula 17        | Tratamento de erros com `error`              |
| Aula 18        | Leitura e escrita de entrada no terminal     |

---

## ✅ Conclusão

Essa aula representa um marco importante: conectar a teoria da programação com armazenamento persistente. Aprender a usar SQLite com Go é essencial para qualquer aplicação real que precisa salvar dados entre execuções. O uso de boas práticas como tratamento de erros, validação de entrada, e organização do código em funções modulares ajuda a construir uma base sólida para aplicações mais complexas.

