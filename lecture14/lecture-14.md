# LECTURE-14: Structs in Go

## 📘 Code Overview (English)

This lesson introduces `struct` types in Go, allowing you to group multiple fields into custom data structures. With structs, you can model real-world entities like users, products, and teams in a clear and organized way.

---

## 🧠 Function-by-Function Descriptions

### `createUser()`
Creates a new `User` struct and prints its contents using `fmt.Printf`. This shows how to initialize and view structured data.

### `updateUserAge()`
Modifies a field (`Age`) of a struct. Reinforces how values inside structs can be changed directly.

### `compareUsers()`
Uses the `==` operator to compare two structs. This is possible because all fields are comparable.

### `printUserDetails(u User)`
Passes a struct as a parameter and prints it. Shows how structs can be used as arguments in functions (Lecture 11).

### `returnUser()`
Returns a `User` struct from a function. Demonstrates struct return values (Lecture 11).

### `createProduct()`
Defines a `Product` struct with multiple field types: `string`, `float64`, `int`.

### `calculateTotalPrice(p Product)`
Calculates total using struct fields. Combines arithmetic from Lecture 4 with structured data.

### `createAddressEmbedded()`
Defines a nested struct. Helps understand how to embed a struct inside another.

### `createAnonymousStruct()`
Shows a temporary, unnamed struct useful for short-lived groupings.

### `pointerToStruct()`
Uses a pointer to a struct and modifies a field. Reinforces pointer logic from Lecture 13.

### `structWithSliceField()`
Introduces structs with slices as fields. Combines structs and collections (Lecture 8).

### `addSkillToDeveloper()`
Appends new values to a slice inside a struct. Reinforces slice manipulation and field updates.

### `employeeFullName()`
Defines a `fullName()` function inside another function. Demonstrates scoping and composition.

### `createTeam()`
Creates a slice of structs (`[]Member`). Introduces collections of structured data.

### `projectWithMapField()`
Stores a `map[string]string` inside a struct. Combines struct and map usage (Lecture 10).

### `iterateTeamMembers()`
Loops through a slice of structs and prints each field. Applies `range` (Lecture 6).

### `compareStructValues()`
Uses `==` to compare two products. Reinforces struct comparability when fields are all comparable.

### `sortUsersByAge()`
Sorts a slice of structs using a custom comparator. Combines structs, slices, and sorting logic.

### `deepCopyStruct()`
Creates a manual copy of a struct to demonstrate value-type behavior in Go.

### `updateStructInFunction(u *User)`
Passes a pointer to a struct and modifies it. Connects struct + pointer knowledge.

---

# AULA-14: Estruturas (Structs) em Go

## 📘 Visão Geral do Código (Português)

Nesta aula você aprenderá como usar `struct` em Go, permitindo agrupar diferentes campos em uma única estrutura. Structs são ideais para representar entidades reais como usuários, produtos e times de forma clara e modular.

---

## 🧠 Explicação Função por Função

### `createUser()`
Cria uma struct `User` e imprime usando `fmt.Printf`. Demonstra inicialização e exibição de dados agrupados.

### `updateUserAge()`
Altera o campo `Age` da struct. Mostra como modificar dados internos.

### `compareUsers()`
Usa `==` para comparar duas structs. Funciona quando todos os campos são comparáveis.

### `printUserDetails(u User)`
Recebe uma struct como argumento e imprime. Reforça passagem por valor estudada na Lecture 11.

### `returnUser()`
Retorna uma struct a partir de uma função. Reforça retorno de tipos personalizados.

### `createProduct()`
Exemplo com campos de tipos mistos (`string`, `float64`, `int`).

### `calculateTotalPrice(p Product)`
Calcula total multiplicando campos. Reforça uso de operadores com structs.

### `createAddressEmbedded()`
Cria uma struct que contém outra como campo. Mostra uso de composição.

### `createAnonymousStruct()`
Cria uma struct anônima sem definição prévia. Útil para casos pontuais.

### `pointerToStruct()`
Usa ponteiro para struct e altera valor. Reforça ponteiros (Lecture 13).

### `structWithSliceField()`
Define struct com campo slice. Une struct + coleção (Lecture 8).

### `addSkillToDeveloper()`
Adiciona elemento a slice dentro da struct. Reforça manipulação de slices.

### `employeeFullName()`
Função interna para gerar nome completo. Ilustra escopo local e retorno composto.

### `createTeam()`
Cria slice de structs `Member`. Introduz coleção de entidades.

### `projectWithMapField()`
Cria struct com `map` interno. Une mapa e struct (Lecture 10).

### `iterateTeamMembers()`
Percorre slice de structs com `range`. Reforça laço com estrutura (Lecture 6).

### `compareStructValues()`
Compara duas structs com `==`. Mostra como structs podem ser comparadas diretamente.

### `sortUsersByAge()`
Ordena slice de structs por idade. Reforça uso de `sort.Slice`.

### `deepCopyStruct()`
Copia struct manualmente. Mostra que structs são valores, não referências.

### `updateStructInFunction(u *User)`
Altera struct por ponteiro. Reforça função com ponteiro como argumento.

