package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	printLectureTitle()

	x := 42
	printAddress(x)
	modifyByValue(x)
	fmt.Println("After modifyByValue:", x)
	modifyByPointer(&x)
	fmt.Println("After modifyByPointer:", x)

	a, b := 3, 7
	swapValuesPointer(&a, &b)
	fmt.Println("Swapped a and b:", a, b)

	resetToZero(&x)
	fmt.Println("After resetToZero:", x)

	x = 5
	doubleValue(&x)
	fmt.Println("After doubleValue:", x)

	y := 10
	compareAddresses(&x, &y)

	copyByReference(&x, &y)
	fmt.Println("y after copy from x:", y)

	numbers := []int{1, 2, 3}
	sumSliceWithPointer(&numbers)

	words := []string{"Go"}
	appendToSlicePointer(&words)
	fmt.Println("Words after append:", words)

	ptr := createPointerReturn(99)
	fmt.Println("Value from created pointer:", *ptr)

	getValueFromPointer(ptr)

	list := []int{10, 20, 30}
	incrementAll(list)
	fmt.Println("After incrementAll:", list)

	ptr1 := new(int)
	ptr2 := new(int)
	*ptr1 = 100
	*ptr2 = 200
	swapPointers(&ptr1, &ptr2)
	fmt.Println("ptr1:", *ptr1, "ptr2:", *ptr2)

	user := User{"Alice", 25}
	updateStructField(&user)
	fmt.Printf("Updated user: %+v ", user)
}

func printLectureTitle() {
	fmt.Println("Lecture 13 - Pointers in Go")
}

// printAddress prints the memory address of a variable
func printAddress(x int) {
	fmt.Printf("Address of x: %p ", &x)
}

// modifyByValue demonstrates that passing by value does not modify the original
func modifyByValue(x int) {
	x = 0
}

// modifyByPointer changes the original variable through pointer
func modifyByPointer(x *int) {
	*x = 99
}

// swapValuesPointer swaps two integers using pointers
func swapValuesPointer(a, b *int) {
	*a, *b = *b, *a
}

// resetToZero sets a value to zero via pointer
func resetToZero(ptr *int) {
	*ptr = 0
}

// doubleValue multiplies the value by 2 via pointer
func doubleValue(ptr *int) {
	*ptr *= 2
}

// compareAddresses checks if two pointers point to same address
func compareAddresses(a, b *int) {
	fmt.Println("a == b?", a == b)
}

// copyByReference copies value from src to dst via pointer
func copyByReference(src, dst *int) {
	*dst = *src
}

// sumSliceWithPointer sums elements in slice using pointer
func sumSliceWithPointer(slice *[]int) {
	sum := 0
	for _, v := range *slice {
		sum += v
	}
	fmt.Println("Sum of slice:", sum)
}

// appendToSlicePointer appends values to a slice using pointer
func appendToSlicePointer(slice *[]string) {
	*slice = append(*slice, "Lang")
}

// createPointerReturn returns a pointer to a new int
func createPointerReturn(value int) *int {
	return &value
}

// getValueFromPointer reads value from pointer
func getValueFromPointer(ptr *int) {
	fmt.Println("Value from pointer:", *ptr)
}

// incrementAll increases each value in slice by 1
func incrementAll(values []int) {
	for i := range values {
		values[i]++
	}
}

// swapPointers swaps two pointers to int
func swapPointers(ptr1, ptr2 **int) {
	*ptr1, *ptr2 = *ptr2, *ptr1
}

// updateStructField modifies a field in a struct using pointer
func updateStructField(u *User) {
	u.Age += 1
}
