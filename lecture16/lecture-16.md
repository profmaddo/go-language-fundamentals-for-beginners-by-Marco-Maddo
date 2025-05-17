# LECTURE-16: Introduction to Object-Oriented Programming (OOP) in Go

## 📘 Understanding the 4 Pillars of OOP

Object-Oriented Programming (OOP) is a way to organize code using real-world concepts like "objects", "attributes", and "behaviors". In Go, we don’t use classes, but we can apply OOP principles using `struct`, `methods`, `interfaces`, and `composition`.

---

### 🧱 The Four Pillars of OOP in Go

#### 1. Abstraction
Simplifies complex reality by modeling classes/structs based on relevant attributes and behaviors
```go
type Person struct {
    name string
    age  int
}
func (p Person) Introduce() {
    fmt.Printf("Hi, I'm %s and I'm %d years old.", p.name, p.age)
}
```

#### 2. Encapsulation
Protects internal data by controlling access via methods.
```go
func (p *Person) SetAge(age int) {
    if age > 0 {
        p.age = age
    }
}
```

#### 3. Inheritance (via Composition)
Allows a struct to reuse another’s properties.
```go
type Employee struct {
    Person
    position string
}
```

#### 4. Polymorphism
Allows different structs to implement the same interface.
```go
type Printer interface {
    Introduce()
}
```

---

## 🧠 Method-by-Method Explanation

Each function in this lecture demonstrates a key OOP principle with a real-world example.

### `Person.Introduce()` – Abstraction  
Represents the basic behavior of a person introducing themselves.

### `Person.SetAge()` / `Person.GetAge()` – Encapsulation  
Controls how `age` is accessed and updated.

### `Employee.Introduce()` – Polymorphism  
Overrides default introduction with job title.

### `Manager.Introduce()` – Polymorphism + Composition  
Adds department information using embedded structs.

### `printIntroduction(p Printer)` – Polymorphism  
Accepts any type implementing `Introduce()`.

### `AuthService.Login()` – Encapsulation  
Protects access logic with private fields (`map` of users).

### `Rectangle` & `Circle` implement `Shape` – Polymorphism  
Different area calculations with the same interface.

### `describeShape(s Shape)` – Polymorphism  
Handles any shape type using interface.

### `Dog.Speak()` / `Cat.Speak()` – Polymorphism  
Two animals with different implementations of `Speak()`.

### `animalTalk(a Animal)` – Polymorphism  
Talks to any animal generically.

### `Engine.Start()` / `Stop()` – Encapsulation  
Internal control of engine state.

### `Car` with `Engine` – Inheritance via Composition  
Uses embedded engine to build behavior of `Car`.

### `Car.Drive()` – Abstraction + Composition  
Behavior depends on internal engine status.

---

# AULA-16: Introdução à Programação Orientada a Objetos (POO) em Go

## 📘 Entendendo os 4 Pilares da POO

Programação Orientada a Objetos (POO) organiza o código com base em objetos, atributos e comportamentos. Em Go, não usamos classes como em Java, mas aplicamos os conceitos com `struct`, `métodos`, `interfaces` e composição.

---

### 🧱 Os Quatro Pilares da POO em Go

#### 1. Abstração
Modela o mundo real com atributos e comportamentos relevantes.
```go
type Person struct {
    name string
    age  int
}
func (p Person) Introduce() {
    fmt.Printf("Oi, eu sou %s e tenho %d anos.", p.name, p.age)
}
```

#### 2. Encapsulamento
Controla o acesso aos dados com métodos.
```go
func (p *Person) SetAge(age int) {
    if age > 0 {
        p.age = age
    }
}
```

#### 3. Herança (via Composição)
Permite reutilizar comportamento através de struct dentro de struct.
```go
type Employee struct {
    Person
    position string
}
```

#### 4. Polimorfismo
Permite diferentes structs implementarem a mesma interface.
```go
type Printer interface {
    Introduce()
}
```

---

## 🧠 Explicação Função por Função

Cada função ilustra um pilar da POO com um exemplo prático.

### `Person.Introduce()` – Abstração  
Modela comportamento real: se apresentar.

### `Person.SetAge()` / `GetAge()` – Encapsulamento  
Controle de leitura e escrita com regras.

### `Employee.Introduce()` – Polimorfismo  
Apresentação com cargo, substituindo o comportamento base.

### `Manager.Introduce()` – Polimorfismo + Composição  
Apresentação com departamento, reutilizando `Employee`.

### `printIntroduction(p Printer)` – Polimorfismo  
Aceita qualquer tipo que implemente `Introduce()`.

### `AuthService.Login()` – Encapsulamento  
Controle de acesso a dados internos.

### `Rectangle` e `Circle` implementam `Shape` – Polimorfismo  
Cálculos diferentes com interface comum.

### `describeShape(s Shape)` – Polimorfismo  
Aceita qualquer forma e chama `Area()`.

### `Dog.Speak()` / `Cat.Speak()` – Polimorfismo  
Implementações diferentes da mesma interface.

### `animalTalk(a Animal)` – Polimorfismo  
Aceita qualquer animal e chama `Speak()`.

### `Engine.Start()` / `Stop()` – Encapsulamento  
Controle interno do estado do motor.

### `Car` com `Engine` – Herança via Composição  
Usa motor embutido para compor funcionalidade.

### `Car.Drive()` – Abstração + Composição  
Dirigir depende do motor, simula comportamento real.

---

## ✅ Conclusão

Com Go, podemos aplicar a POO de forma simples e eficiente, focando em **organização, clareza e reuso**. Essa lecture é essencial para estruturar sistemas maiores com boas práticas de design.

