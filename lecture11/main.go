package main

import (
	"fmt"
)

func main() {
	printLectureTitle()
	printGreeting("Alice")
	fmt.Println("Sum:", sum(5, 3))
	result, ok := divide(10, 2)
	fmt.Printf("Divide result: %.2f, Success: %v ", result, ok)
	fmt.Println("Area:", calculateArea(4.5, 2.0))
	name, age := getUserInfo()
	fmt.Printf("User: %s (%d years old) ", name, age)
	fmt.Printf("Celsius 30 → Fahrenheit: %.1f ", convertCelsiusToFahrenheit(30))
	fmt.Println("Is 4 even?", isEven(4))
	fmt.Println("Repeated word:", repeatWord("Go", 3))
	first, last := getFirstAndLast([]int{10, 20, 30, 40})
	fmt.Printf("First: %d, Last: %d ", first, last)
	fmt.Println("Contains 'apple'?", contains([]string{"banana", "apple", "grape"}, "apple"))
	total, avg := sumAndAverage([]int{10, 20, 30})
	fmt.Printf("Sum: %d, Average: %.2f ", total, avg)
	a, b := swapValues(5, 9)
	fmt.Printf("Swapped: %d, %d ", a, b)
	fmt.Println("Full name:", makeFullName("John", "Doe"))
	fmt.Printf("Total Price: $%.2f ", calculatePrice(9.99, 3))
	showScopeExample()
}

func printLectureTitle() {
	fmt.Println("Lecture 11 - Functions with Parameters and Return Values")
}

// printGreeting prints a greeting message
func printGreeting(name string) {
	fmt.Printf("Hello, %s! ", name)
}

// sum returns the sum of two integers
func sum(a int, b int) int {
	return a + b
}

// divide returns the division result and a success flag
func divide(a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// calculateArea returns the area of a rectangle
func calculateArea(length, width float64) float64 {
	return length * width
}

// getUserInfo returns a sample name and age
func getUserInfo() (string, int) {
	return "Marco", 35
}

// convertCelsiusToFahrenheit converts temperature
func convertCelsiusToFahrenheit(c float64) float64 {
	return c*1.8 + 32
}

// isEven returns true if a number is even
func isEven(n int) bool {
	return n%2 == 0
}

// repeatWord repeats a word n times
func repeatWord(word string, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result += word
	}
	return result
}

// getFirstAndLast returns the first and last elements from a slice
func getFirstAndLast(elements []int) (int, int) {
	if len(elements) == 0 {
		return 0, 0
	}
	return elements[0], elements[len(elements)-1]
}

// contains returns true if target exists in the slice
func contains(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

// sumAndAverage returns the total sum and average of a slice
func sumAndAverage(numbers []int) (int, float64) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	avg := float64(sum) / float64(len(numbers))
	return sum, avg
}

// swapValues swaps two integers
func swapValues(a, b int) (int, int) {
	return b, a
}

// makeFullName concatenates first and last name
func makeFullName(first, last string) string {
	return first + " " + last
}

// calculatePrice returns the total price
func calculatePrice(price float64, quantity int) float64 {
	return price * float64(quantity)
}

// showScopeExample demonstrates variable scope
func showScopeExample() {
	x := 10 // local variable
	fmt.Println("Inside function, x =", x)
}
