package main

import (
	"fmt"
)

func main() {
	printLectureTitle()
	basicArithmeticOperations()
	incrementDecrementOperators()
	comparisonOperators()
	logicalOperators()
	arithmeticWithInput(8, 3)
	comparisonWithInput(10, 5)
	logicalEvaluationWithFlags(true, false)
	boolExpressionsInPrintf(7, 7)
	operatorPrecedenceExample()
	assignmentOperators()
	combinedLogicalComparison(20, true)
}

func printLectureTitle() {
	fmt.Println("Lecture 9 - Operators in Go")
}

// 1. Basic arithmetic operations
func basicArithmeticOperations() {
	a, b := 10, 4
	fmt.Println("Arithmetic operations:")
	fmt.Printf("%d + %d = %d", a, b, a+b)
	fmt.Printf("%d - %d = %d", a, b, a-b)
	fmt.Printf("%d * %d = %d", a, b, a*b)
	fmt.Printf("%d / %d = %d", a, b, a/b)
	fmt.Printf("%d %% %d = %d", a, b, a%b)
}

// 2. Increment and decrement
func incrementDecrementOperators() {
	x := 5
	fmt.Println("Before increment:", x)
	x++
	fmt.Println("After increment:", x)
	x--
	fmt.Println("After decrement:", x)
}

// 3. Comparison operators
func comparisonOperators() {
	a, b := 5, 10
	fmt.Println("Comparison results:")
	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a > b:", a > b)
	fmt.Println("a < b:", a < b)
	fmt.Println("a >= b:", a >= b)
	fmt.Println("a <= b:", a <= b)
}

// 4. Logical operators
func logicalOperators() {
	fmt.Println("Logical operations:")
	fmt.Println("true && false:", true && false)
	fmt.Println("true || false:", true || false)
	fmt.Println("!true:", !true)
}

// 5. Arithmetic operations with input
func arithmeticWithInput(a, b int) {
	fmt.Printf("Arithmetic with input: %d and %d ", a, b)
	fmt.Printf("%d + %d = %d", a, b, a+b)
	fmt.Printf("%d * %d = %d", a, b, a*b)
}

// 6. Comparisons with input	
func comparisonWithInput(a, b int) {
	fmt.Printf(" Comparing %d and %d ", a, b)
	fmt.Printf("Equal? %v | Greater? %v | Less? %v", a == b, a > b, a < b)
}

// 7. Logical evaluation with flags
func logicalEvaluationWithFlags(isAdult, hasLicense bool) {
	if isAdult && hasLicense {
		fmt.Println(" You can drive.")
	} else {
		fmt.Println(" You cannot drive.")
	}
}

// 8. Boolean expressions in Printf
func boolExpressionsInPrintf(n1, n2 int) {
	equal := n1 == n2
	fmt.Printf("Numbers: %d and %d - Equal? %v", n1, n2, equal)
}

// 9. Operator precedence
func operatorPrecedenceExample() {
	a, b, c := 2, 3, 4
	result1 := a + b*c
	result2 := (a + b) * c
	fmt.Println(" Operator precedence:")
	fmt.Printf("a + b * c = %d", result1)
	fmt.Printf("(a + b) * c = %d", result2)
}

// 10. Assignment operators
func assignmentOperators() {
	x := 10
	fmt.Println(" Original x:", x)
	x += 5
	fmt.Println("x after += 5:", x)
	x *= 2
	fmt.Println("x after *= 2:", x)
	x -= 3
	fmt.Println("x after -= 3:", x)
	x /= 4
	fmt.Println("x after /= 4:", x)
	x %= 3
	fmt.Println("x after %= 3:", x)
}

// 11. Combined comparison and logic
func combinedLogicalComparison(age int, hasTicket bool) {
	if age >= 18 && hasTicket {
		fmt.Println("Access granted.")
	} else {
		fmt.Println("Access denied.")
	}
}
