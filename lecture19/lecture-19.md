# LECTURE-19: Goroutines and Basic Concurrency in Go

## 📘 Topic Overview

Concurrency is one of Go’s key strengths. In this lesson, you will learn to create lightweight threads of execution using `go` and to communicate safely between them using `channels`. You’ll also learn how to synchronize concurrent tasks using `sync.WaitGroup` and control parallelism with buffered channels and semaphores.

---

## 🔧 Function-by-Function Summary

### `printMessage(msg)`
Launches a goroutine to print a message concurrently.

### `countTo(n)`
Runs a loop inside a goroutine using `time.Sleep()` for delay.

### `sendNumbers(ch)` and `receiveNumbers(ch)`
Demonstrates how to send and receive values using an unbuffered channel.

### `sum(a, b, ch)`
Performs a sum and sends the result through a channel.

### `runWithDelay(msg)`
Simulates delayed tasks.

### `launchMultipleGoroutines()`
Starts multiple concurrent tasks using goroutines.

### `parallelSquares(numbers)`
Calculates squares of numbers in parallel and collects results.

### `channelBufferedExample()`
Demonstrates a buffered channel to allow asynchronous writes.

### `channelDirectionExample(send, recv)`
Shows unidirectional channel usage for safer function contracts.

### `closeChannelExample()`
Safely receives from a closed channel using `range`.

### `multiplexingWithSelect()`
Uses `select` to wait on multiple channel operations.

### `taskWithReturn(name)`
Returns a receive-only channel from a background goroutine.

### `useWaitGroupForSync()`
Uses `sync.WaitGroup` to wait for goroutines to finish.

### `limitedParallelism()`
Limits the number of goroutines executing concurrently using a buffered channel as semaphore.

---

## 🔁 Integration with Previous Lessons

| Lecture       | Concepts Reused                            |
|---------------|---------------------------------------------|
| Lecture 6     | Loops with `range` and delay with `Sleep`   |
| Lecture 11    | Functions with parameters and return values |
| Lecture 12    | Anonymous functions and goroutines          |
| Lecture 17    | Error-safe structure using channels         |
| Lecture 18    | Parallel file writing could benefit from goroutines |

---

# AULA-19: Goroutines e Concorrência Básica em Go

## 📘 Visão Geral

Concorrência é um dos pilares da linguagem Go. Nesta aula você aprenderá a usar goroutines (`go`) para executar tarefas em paralelo e canais (`chan`) para permitir a comunicação entre essas tarefas de forma segura e eficiente. Também aprenderá como sincronizar execuções com `sync.WaitGroup` e limitar paralelismo com semáforos via canais.

---

## 🔧 Resumo das Funções

### `printMessage(msg)`
Imprime mensagem usando goroutine.

### `countTo(n)`
Executa laço com `Sleep` em concorrência.

### `sendNumbers(ch)` / `receiveNumbers(ch)`
Demonstra envio e recepção em canal não-bufferizado.

### `sum(a, b, ch)`
Executa soma e envia resultado pelo canal.

### `runWithDelay(msg)`
Simula tarefa assíncrona com atraso aleatório.

### `launchMultipleGoroutines()`
Executa várias tarefas concorrentes.

### `parallelSquares(numbers)`
Calcula quadrados em paralelo e coleta resultados.

### `channelBufferedExample()`
Mostra uso de canal bufferizado para escrita assíncrona.

### `channelDirectionExample(send, recv)`
Exemplo de canais com direção restrita.

### `closeChannelExample()`
Recepção segura de canal fechado com `range`.

### `multiplexingWithSelect()`
Uso de `select` para múltiplos canais.

### `taskWithReturn(name)`
Função que retorna canal com resultado assíncrono.

### `useWaitGroupForSync()`
Sincroniza goroutines com `sync.WaitGroup`.

### `limitedParallelism()`
Limita execução concorrente usando canal como semáforo.

---

## 🔁 Integração com Aulas Anteriores

| Aula           | Conceitos Reforçados                       |
|----------------|---------------------------------------------|
| Aula 6         | Laços e atrasos com `for` e `Sleep`        |
| Aula 11        | Funções com parâmetros e retornos          |
| Aula 12        | Funções anônimas e goroutines              |
| Aula 17        | Segurança e controle de concorrência       |
| Aula 18        | Uso potencial de concorrência para I/O     |

---

## ✅ Conclusão

Com goroutines e canais, você pode construir aplicações eficientes, paralelas e com excelente controle de execução. Go simplifica concorrência com uma sintaxe poderosa e intuitiva, que é ideal para sistemas modernos e de alto desempenho.

