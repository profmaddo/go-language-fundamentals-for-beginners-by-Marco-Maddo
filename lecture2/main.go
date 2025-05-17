package main

import "fmt"

func main() {
	lecturesubject()
	variablefloat()
}

func lecturesubject() {
	fmt.Println("Numbers Float")
}

func variablefloat() {
	var numberpi32 float32 = 3.141516
    var numberpi64 float64 = 3.1415161718192021222324

	fmt.Println("Float 32 bits: ", numberpi32)
	fmt.Println("Float 64 bits: ", numberpi64)

}
