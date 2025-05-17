package main

import "fmt"

func main() {
	helloworld()
	variableinteger()
}

func helloworld() {
	fmt.Println("Hello, World! I'm GO... Let's go?")
}

func variableinteger() {
	var year int = 2025
	fmt.Println(year)
	fmt.Println("Year: ", year)
	fmt.Sprintln("The Year is %d", year)

}
