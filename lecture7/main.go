package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	printLectureTitle()
	readIntegerWithScan()
	readFloatWithScan()
	readMultipleValuesScan()
	readLineWithBufio()
	readCommandLineArgs()
	readChoiceAndSwitch()
}

func printLectureTitle() {
	fmt.Println("Lecture 7 - Reading Input in Go")
}

// Exemplo 1: Lê número inteiro com fmt.Scan
func readIntegerWithScan() {
	var number int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&number)
	fmt.Println("You entered:", number)
}

// Exemplo 2: Lê número decimal (float) com fmt.Scan
func readFloatWithScan() {
	var decimal float64
	fmt.Print("Enter a decimal number: ")
	fmt.Scan(&decimal)
	fmt.Printf("You entered: %.2f\n", decimal)
}

// Exemplo 3: Lê múltiplos valores com fmt.Scan
func readMultipleValuesScan() {
	var name string
	var age int
	fmt.Print("Enter your name and age (e.g. Alice 30): ")
	fmt.Scan(&name, &age)
	fmt.Printf("Hello %s, you are %d years old.\n", name, age)
}

// Exemplo 4: Lê linha completa com bufio
func readLineWithBufio() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a full sentence: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("You wrote:", strings.TrimSpace(text))
}

// Exemplo 5: Lê argumentos da linha de comando
func readCommandLineArgs() {
	fmt.Println("Command-line arguments (os.Args):")
	for i, arg := range os.Args {
		fmt.Printf("Arg %d: %s\n", i, arg)
	}
}

// Exemplo 6: Lê opção do usuário e usa switch
func readChoiceAndSwitch() {
	var choice int
	fmt.Print("Choose a number (1 to 3): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("You chose ONE.")
	case 2:
		fmt.Println("You chose TWO.")
	case 3:
		fmt.Println("You chose THREE.")
	default:
		fmt.Println("Invalid choice.")
	}
}
