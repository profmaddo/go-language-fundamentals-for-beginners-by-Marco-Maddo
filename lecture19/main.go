package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simples goroutine com mensagem
func printMessage(msg string) {
	fmt.Println("Message:", msg)
}

// Loop concorrente
func countTo(n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("Counting: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

// Envia números para canal
func sendNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

// Recebe números do canal
func receiveNumbers(ch chan int) {
	for num := range ch {
		fmt.Println("Received:", num)
	}
}

// Soma com retorno por canal
func sum(a, b int, ch chan int) {
	ch <- a + b
}

// Goroutine com delay
func runWithDelay(msg string) {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	fmt.Println("Done:", msg)
}

// Lança múltiplas goroutines
func launchMultipleGoroutines() {
	for i := 1; i <= 3; i++ {
		go runWithDelay(fmt.Sprintf("Task %d", i))
	}
	time.Sleep(1 * time.Second)
}

// Calcula quadrados em paralelo
func parallelSquares(numbers []int) {
	ch := make(chan int)
	for _, n := range numbers {
		go func(n int) {
			ch <- n * n
		}(n)
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Square:", <-ch)
	}
}

// Canal bufferizado
func channelBufferedExample() {
	ch := make(chan string, 2)
	ch <- "Go"
	ch <- "Lang"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// Exemplo de direção de canal
func channelDirectionExample(send chan<- string, recv <-chan string) {
	send <- "data sent"
	msg := <-recv
	fmt.Println("Direction channel:", msg)
}

// Exemplo com close()
func closeChannelExample() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch)
	}()
	for val := range ch {
		fmt.Println("From closed channel:", val)
	}
}

// Multiplexação com select
func multiplexingWithSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(300 * time.Millisecond)
		c1 <- "from c1"
	}()
	go func() {
		time.Sleep(200 * time.Millisecond)
		c2 <- "from c2"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("Selected:", msg1)
	case msg2 := <-c2:
		fmt.Println("Selected:", msg2)
	}
}

// Canal como retorno
func taskWithReturn(name string) <-chan string {
	ch := make(chan string)
	go func() {
		time.Sleep(250 * time.Millisecond)
		ch <- fmt.Sprintf("result from %s", name)
	}()
	return ch
}

// Exemplo com WaitGroup
func useWaitGroupForSync() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(400)) * time.Millisecond)
			fmt.Printf("Finished task %d\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("All goroutines done.")
}

// Controle de paralelismo limitado
func limitedParallelism() {
	sem := make(chan struct{}, 2) // max 2 at a time
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sem <- struct{}{}
			fmt.Printf("Running task %d\n", i)
			time.Sleep(300 * time.Millisecond)
			<-sem
		}(i)
	}
	wg.Wait()
}

func main() {
	fmt.Println("== Simple goroutine ==")
	go printMessage("Hello from goroutine")
	time.Sleep(100 * time.Millisecond)

	fmt.Println("== Count to 3 ==")
	go countTo(3)
	time.Sleep(400 * time.Millisecond)

	fmt.Println("== Channel send/receive ==")
	ch := make(chan int)
	go sendNumbers(ch)
	receiveNumbers(ch)

	fmt.Println("== Sum with channel ==")
	sumCh := make(chan int)
	go sum(4, 6, sumCh)
	fmt.Println("Sum:", <-sumCh)

	fmt.Println("== Launching multiple tasks ==")
	launchMultipleGoroutines()

	fmt.Println("== Parallel squares ==")
	parallelSquares([]int{2, 4, 6})

	fmt.Println("== Buffered channel ==")
	channelBufferedExample()

	fmt.Println("== Close channel example ==")
	closeChannelExample()

	fmt.Println("== Multiplexing with select ==")
	multiplexingWithSelect()

	fmt.Println("== Task with return ==")
	result := taskWithReturn("background")
	fmt.Println(<-result)

	fmt.Println("== WaitGroup example ==")
	useWaitGroupForSync()

	fmt.Println("== Limited parallelism ==")
	limitedParallelism()
}
