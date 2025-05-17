package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	printLectureTitle()
	printFixedArray()
	readAndFillArray()
	calculateSumArray()
	createAndPrintSlice()
	appendToSlice()
	sliceFromArray()
	readWordsToSlice()
	searchInSlice()
	removeItemFromSlice()
	sliceCapacityGrowth()
}

func printLectureTitle() {
	fmt.Println("Lecture 8 - Arrays and Slices in Go")
}

// 1. Array fixo e impressão com loop
func printFixedArray() {
	var numbers = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Fixed array values:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Index %d: %d\n", i, numbers[i])
	}
}

// 2. Entrada de dados para preencher array
func readAndFillArray() {
	var arr [3]int
	fmt.Println("Enter 3 integers:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Value %d: ", i+1)
		fmt.Scan(&arr[i])
	}
	fmt.Println("You entered:", arr)
}

// 3. Soma dos elementos do array
func calculateSumArray() {
	numbers := [5]int{2, 4, 6, 8, 10}
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	fmt.Println("Sum of array elements:", sum)
}

// 4. Criar e imprimir slice
func createAndPrintSlice() {
	slice := []string{"Go", "Java", "Python"}
	fmt.Println("Languages slice:")
	for index, value := range slice {
		fmt.Printf("Index %d: %s\n", index, value)
	}
}

// 5. Adicionar elementos com append
func appendToSlice() {
	var fruits []string
	fruits = append(fruits, "Apple", "Banana")
	fruits = append(fruits, "Cherry")
	fmt.Println("Fruits slice after append:", fruits)
}

// 6. Criar slice a partir de array
func sliceFromArray() {
	arr := [5]int{1, 2, 3, 4, 5}
	sub := arr[1:4] // indexes 1 to 3
	fmt.Println("Slice from array (arr[1:4]):", sub)
}

// 7. Ler palavras do usuário para slice
func readWordsToSlice() {
	reader := bufio.NewReader(os.Stdin)
	var words []string
	fmt.Print("Enter 3 words separated by ENTER:\n")
	for i := 0; i < 3; i++ {
		fmt.Printf("Word %d: ", i+1)
		text, _ := reader.ReadString('\n')
		words = append(words, strings.TrimSpace(text))
	}
	fmt.Println("You entered:", words)
}

// 8. Procurar valor no slice
func searchInSlice() {
	slice := []int{5, 10, 15, 20}
	var value int
	fmt.Print("Enter a number to search: ")
	fmt.Scan(&value)
	found := false
	for _, v := range slice {
		if v == value {
			found = true
			break
		}
	}
	if found {
		fmt.Println("Value found in slice.")
	} else {
		fmt.Println("Value not found.")
	}
}

// 9. Remover item do slice por índice
func removeItemFromSlice() {
	slice := []string{"red", "green", "blue", "yellow"}
	fmt.Println("Original slice:", slice)
	index := 2 // remove "blue"
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println("After removing index 2:", slice)
}

// 10. Mostrar crescimento de capacidade do slice
func sliceCapacityGrowth() {
	var numbers []int
	fmt.Println("Slice capacity growth:")
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("Len: %d, Cap: %d, Values: %v\n", len(numbers), cap(numbers), numbers)
	}
}
