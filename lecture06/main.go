package main

import "fmt"

func main() {
	printLectureTitle()
	forCounterLoop()
	forConditionalLoop()
	forRangeSlice()
	forRangeString()
	nestedLoopExample()
}

// Título da aula
func printLectureTitle() {
	fmt.Println("Lecture 6 - Loops and Repetition in Go")
}

// Loop com contador clássico
func forCounterLoop() {
	fmt.Println("For loop with counter (1 to 5):")
	for i := 1; i <= 5; i++ {
		fmt.Println("i =", i)
	}
}

// Loop com condição
func forConditionalLoop() {
	fmt.Println("For loop with condition (until x > 5):")
	x := 1
	for x <= 5 {
		fmt.Println("x =", x)
		x++
	}
}

// Loop com range sobre slice
func forRangeSlice() {
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("For loop with range over slice:")
	for index, value := range numbers {
		fmt.Printf("Index %d: Value %d", index, value)
	}
}

// Loop com range sobre string
func forRangeString() {
	text := "GoLang"
	fmt.Println("For loop with range over string:")
	for index, char := range text {
		fmt.Printf("Index %d: Character %c", index, char)
	}
}

// Loop aninhado
func nestedLoopExample() {
	fmt.Println("Nested loops - multiplication table 1 to 3:")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d x %d = %d", i, j, i*j)
		}
	}
}
