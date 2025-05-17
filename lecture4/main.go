package main

import (
	"fmt"
	"math"
)

func main() {
	printLectureTitle()
	printInt8Example()
	printUint8Example()
	printInt16Example()
	printUint16Example()
	printInt32Example()
	printInt64Example()
	printFloat32Example()
	printFloat64Example()
	printComplex64Example()
	printComplex128Example()
}

func printLectureTitle() {
	fmt.Println("Lecture 4 - Numeric Data Types in Go")
}

// Integers with sign
func printInt8Example() {
	var num int8 = 127
	fmt.Println("int8 example (max value):", num)
	// num = num + 1 // would overflow!
}

func printUint8Example() {
	var num uint8 = 255
	fmt.Println("uint8 example (max value):", num)
}

func printInt16Example() {
	var a int16 = 20000
	var b int16 = -15000
	fmt.Println("int16 example: a =", a, ", b =", b, ", a + b =", a+b)
}

func printUint16Example() {
	var a uint16 = 40000
	var b uint16 = 10000
	fmt.Println("uint16 example: a =", a, ", b =", b, ", a - b =", a-b)
}

func printInt32Example() {
	var num int32 = 1000000000
	fmt.Println("int32 example:", num, "* 2 =", num*2)
}

func printInt64Example() {
	var num int64 = 9223372036854775807
	fmt.Println("int64 example (max value):", num)
}

// Floating point numbers
func printFloat32Example() {
	var pi float32 = 3.1415926535
	fmt.Println("float32 example:", pi, "* 2 =", pi*2)
}

func printFloat64Example() {
	var pi float64 = 3.141592653589793238
	fmt.Println("float64 example:", pi, "* 2 =", pi*2)
}

// Complex numbers
func printComplex64Example() {
	var z complex64 = complex(2, 3) // 2 + 3i
	fmt.Println("complex64 example:", z, ", real part:", real(z), ", imaginary part:", imag(z))
}

func printComplex128Example() {
	var z complex128 = complex(math.Pi, math.E)
	fmt.Println("complex128 example:", z, ", real part:", real(z), ", imaginary part:", imag(z))
}
