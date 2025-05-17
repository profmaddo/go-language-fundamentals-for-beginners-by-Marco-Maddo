# LECTURE-2: Floating Point Numbers in Go

## 📘 Code Explanation (English)

This Go program demonstrates how to work with floating point numbers using two types: `float32` and `float64`.

### 🔹 Code Structure

- `package main`: This is the entry point package for executable programs in Go.
- `import "fmt"`: Imports the standard formatting package to allow printing to the terminal.
- `func main()`: The starting point of the application.
    - It calls `lecturesubject()` to print a subject title.
    - It then calls `variablefloat()` to demonstrate the use of floating point variables.

### 🔹 Functions

- `lecturesubject()`: Prints `"Numbers Float"` to indicate the topic of the code.
- `variablefloat()`: 
    - Declares a `float32` variable called `numberpi32` and assigns the value `3.141516`.
    - Declares a `float64` variable called `numberpi64` with a more precise value.
    - Prints both values using `fmt.Println`.

### 🔹 float32 vs float64

| Type     | Precision | Typical Use                          |
|----------|-----------|--------------------------------------|
| float32  | ~7 digits | Saves memory, faster on some CPUs    |
| float64  | ~15 digits| Default type for high-precision math |

- `float32` supports up to approximately **1.18 × 10^-38 to 3.4 × 10^38**
- `float64` supports up to approximately **2.23 × 10^-308 to 1.8 × 10^308**

### 💡 When to use which?

- Use `float32` when memory space is limited or precision is less critical (e.g., graphics, embedded systems).
- Use `float64` when accuracy is important (e.g., scientific calculations, finance).

---

# AULA-2: Números Ponto Flutuante em Go

## 📘 Explicação do Código (Português)

Este programa em Go demonstra como trabalhar com números de ponto flutuante usando dois tipos: `float32` e `float64`.

### 🔹 Estrutura do Código

- `package main`: Pacote de entrada obrigatório para programas executáveis em Go.
- `import "fmt"`: Importa o pacote padrão para saída formatada.
- `func main()`: Ponto inicial da aplicação.
    - Chama `lecturesubject()` para imprimir o tema.
    - Em seguida chama `variablefloat()` para demonstrar o uso de variáveis float.

### 🔹 Funções

- `lecturesubject()`: Imprime `"Numbers Float"` como título do exemplo.
- `variablefloat()`: 
    - Declara a variável `numberpi32` do tipo `float32` com o valor `3.141516`.
    - Declara a variável `numberpi64` do tipo `float64` com mais precisão.
    - Imprime os dois valores com `fmt.Println`.

### 🔹 float32 vs float64

| Tipo     | Precisão     | Quando Usar                           |
|----------|--------------|----------------------------------------|
| float32  | ~7 dígitos   | Quando a memória é limitada ou alta precisão não é necessária |
| float64  | ~15 dígitos  | Para cálculos científicos, finanças, maior precisão |

- `float32`: suporta valores entre **1.18 × 10^-38 até 3.4 × 10^38**
- `float64`: suporta valores entre **2.23 × 10^-308 até 1.8 × 10^308**

### 💡 Quando usar cada um?

- Use `float32` quando estiver otimizado para memória (ex: gráficos, dispositivos embarcados).
- Use `float64` quando a precisão é essencial (ex: finanças, cálculos científicos).
