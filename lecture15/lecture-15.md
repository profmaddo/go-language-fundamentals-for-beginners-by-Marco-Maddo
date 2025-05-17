# LECTURE-15: Methods with Structs in Go

## 📘 Code Overview (English)

In this lesson, we explore how to associate methods with `struct` types in Go. A method is simply a function with a **receiver** – it allows the struct to perform actions or return computed results. This enables a structured and reusable programming style, similar to object-oriented programming.

---

## 🧠 Method-by-Method Explanations

### `User.Greet()`
Prints a greeting using the struct's `Name`. Demonstrates how methods can access struct fields.

### `User.IncreaseAge()`
Pointer receiver. Modifies the `Age` field directly, reinforcing pointer behavior (Lecture 13).

### `User.IsAdult()`
Returns `true` if `Age >= 18`. Logic based on struct field value.

### `User.Compare(other User)`
Compares two users by name and age. Shows how methods can compare instances.

### `Product.TotalPrice()`
Returns calculated total by multiplying `Price * Quantity`. Applies numeric logic from Lecture 4.

### `Product.Discount(percent)`
Pointer receiver. Reduces `Price` by a percentage. Updates state.

### `Address.FullAddress()`
Concatenates fields into a full address string. Simple field formatting.

### `Developer.ListSkills()`
Prints slice contents. Combines structs and slices (Lecture 8).

### `Developer.AddSkill(skill)`
Pointer receiver. Adds new skill to the `Skills` slice. Shows in-place slice update.

### `Member.ChangeRole(newRole)`
Changes the `Role` field directly. Useful in team management logic.

### `Project.TagSummary()`
Returns a string summarizing `map` key-value pairs. Reuses map logic (Lecture 10).

### `Project.AddTag(key, value)`
Pointer receiver. Adds new tags dynamically. Reinforces reference manipulation.

### `Team.ListMembers()`
Loops over a slice of structs and prints each. Demonstrates nested data handling.

### `Team.AddMember(member)`
Pointer receiver. Adds a `Member` to the team. Shows dynamic slice growth.

### `Team.MemberCount()`
Returns length of `Members`. Combines method logic with slice introspection.

### `Rectangle.Area()`
Calculates and returns area of a rectangle. Uses numeric fields.

### `Rectangle.Perimeter()`
Computes perimeter based on width and height.

### `Book.Summary()`
Formats all fields into a descriptive string. Reuses concepts from Lecture 3.

### `Book.UpdateTitle(title)`
Pointer receiver. Updates the `Title` field in-place.

### `Book.CompareTitle(other)`
Compares title with another book. Useful in equality filtering logic.

---

# AULA-15: Métodos com Structs em Go

## 📘 Visão Geral do Código (Português)

Nesta aula, associamos **métodos a structs**. Um método é uma função com **receiver** – que permite que a struct execute ações, acesse ou modifique seus próprios campos. Esse padrão traz clareza, encapsulamento e reaproveitamento de lógica.

---

## 🧠 Explicação de Cada Método

### `User.Greet()`
Imprime uma saudação com base no campo `Name`.

### `User.IncreaseAge()`
Receiver com ponteiro. Altera diretamente o campo `Age`.

### `User.IsAdult()`
Retorna se a idade do usuário é maior ou igual a 18.

### `User.Compare(other)`
Compara dois usuários por nome e idade.

### `Product.TotalPrice()`
Retorna total com base em `Preço * Quantidade`.

### `Product.Discount(percent)`
Altera o preço aplicando desconto percentual. Usa receiver com ponteiro.

### `Address.FullAddress()`
Concatena rua e cidade em uma string.

### `Developer.ListSkills()`
Mostra os skills do developer. Uso de slice com struct.

### `Developer.AddSkill(skill)`
Receiver com ponteiro. Adiciona skill ao slice.

### `Member.ChangeRole(newRole)`
Altera o campo `Role`.

### `Project.TagSummary()`
Concatena os pares chave/valor de um `map`.

### `Project.AddTag(key, value)`
Adiciona um novo par ao `map`. Receiver com ponteiro.

### `Team.ListMembers()`
Percorre e exibe membros de um slice de structs.

### `Team.AddMember(member)`
Adiciona membro à equipe. Receiver com ponteiro.

### `Team.MemberCount()`
Retorna o número de membros.

### `Rectangle.Area()`
Retorna a área (largura × altura).

### `Rectangle.Perimeter()`
Calcula o perímetro do retângulo.

### `Book.Summary()`
Retorna um resumo do livro com título, autor e páginas.

### `Book.UpdateTitle(title)`
Atualiza o título do livro. Receiver com ponteiro.

### `Book.CompareTitle(other)`
Compara títulos de dois livros.

---

## ✅ Conclusão

Essa lecture reforça o conceito de **comportamento associado a dados**, usando métodos com e sem ponteiros. Ao entender como `struct` e métodos trabalham juntos, você estará pronto para escrever **código mais limpo, reutilizável e alinhado ao estilo Go**.
