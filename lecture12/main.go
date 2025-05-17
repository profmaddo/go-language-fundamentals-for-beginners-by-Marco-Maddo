package main

import (
	"fmt"
	"sort"
)

func main() {
	printLectureTitle()

	fmt.Println("Factorial of 5:", factorial(5))
	fmt.Println("Sum variadic:", sumVariadic(2, 4, 6, 8))

	x := 10
	incrementByPointer(&x)
	fmt.Println("Incremented by pointer:", x)

	a, b := 3, 7
	swapPointers(&a, &b)
	fmt.Println("Swapped values:", a, b)

	fmt.Println("Apply sum function:", applyFunction(5, 3, func(a, b int) int { return a + b }))

	double := getMultiplier(2)
	fmt.Println("Double 5:", double(5))

	printAnonymous()

	values := []int{5, 2, 8, 1}
	sortDescending(values)
	fmt.Println("Sorted descending:", values)

	mapped := mapFunction([]int{1, 2, 3}, func(n int) int { return n * n })
	fmt.Println("Mapped slice (squared):", mapped)

	filtered := filter([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 })
	fmt.Println("Filtered (evens):", filtered)

	counter := closureCounter()
	fmt.Println("Counter:", counter())
	fmt.Println("Counter:", counter())

	fmt.Println("Is 'radar' palindrome?", isPalindrome("radar"))

	min, max := calculateStats(7, 3, 9, 2, 8)
	fmt.Printf("Min: %d, Max: %d ", min, max)

	names := []string{"Alice", "Bob"}
	printValuesByReference(&names)

	runWithLog(func() {
		fmt.Println("Running important task...")
	})
}

func printLectureTitle() {
	fmt.Println("Lecture 12 - Advanced Functions in Go")
}

// factorial calculates factorial recursively
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// sumVariadic sums any number of integers
func sumVariadic(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

// incrementByPointer increases value via pointer
func incrementByPointer(x *int) {
	*x++
}

// swapPointers exchanges values by reference
func swapPointers(a, b *int) {
	*a, *b = *b, *a
}

// applyFunction executes any binary function
func applyFunction(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// getMultiplier returns a closure that multiplies by factor
func getMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

// printAnonymous defines and runs an anonymous function
func printAnonymous() {
	func() {
		fmt.Println("Hello from anonymous function!")
	}()
}

// sortDescending sorts a slice in descending order
func sortDescending(slice []int) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
}

// mapFunction applies a function to each slice element
func mapFunction(slice []int, f func(int) int) []int {
	var result []int
	for _, v := range slice {
		result = append(result, f(v))
	}
	return result
}

// filter returns elements that pass the test
func filter(slice []int, test func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if test(v) {
			result = append(result, v)
		}
	}
	return result
}

// closureCounter returns a counter function
func closureCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// isPalindrome checks if word reads the same backward
func isPalindrome(word string) bool {
	if len(word) <= 1 {
		return true
	}
	if word[0] != word[len(word)-1] {
		return false
	}
	return isPalindrome(word[1 : len(word)-1])
}

// calculateStats returns min and max from variadic input
func calculateStats(numbers ...int) (int, int) {
	if len(numbers) == 0 {
		return 0, 0
	}
	min, max := numbers[0], numbers[0]
	for _, n := range numbers[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

// printValuesByReference prints slice values using pointer
func printValuesByReference(values *[]string) {
	for i, v := range *values {
		fmt.Printf("Index %d: %s ", i, v)
	}
}

// runWithLog wraps a function with logging
func runWithLog(f func()) {
	fmt.Println("Starting task...")
	f()
	fmt.Println("Task finished.")
}
