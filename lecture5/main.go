package main

import "fmt"

func main() {
	printLectureTitle()
	checkIfAdult(20)
	checkIfAdult(15)
	checkGradeIfElse(6.5)
	checkGradeIfElse(4.5)
	checkGradeIfElseIf(9.2)
	checkGradeIfElseIf(6.7)
	checkGradeIfElseIf(3.9)
	checkDayWithSwitch(1)
	checkDayWithSwitch(4)
	checkDayWithSwitch(7)
	checkGradeWithSwitchTrue(9.5)
	checkGradeWithSwitchTrue(6.2)
	checkGradeWithSwitchTrue(3.3)
}

func printLectureTitle() {
	fmt.Println("Lecture 5 - Decision Making Structures in Go")
}

// if statement
func checkIfAdult(age int) {
	if age >= 18 {
		fmt.Printf("Age %d: Adult\n", age)
	} else {
		fmt.Printf("Age %d: Minor\n", age)
	}
}

// if...else
func checkGradeIfElse(score float64) {
	if score >= 6 {
		fmt.Printf("Score %.1f: Approved\n", score)
	} else {
		fmt.Printf("Score %.1f: Failed\n", score)
	}
}

// if...else if...else
func checkGradeIfElseIf(score float64) {
	if score >= 9 {
		fmt.Printf("Score %.1f: Excellent\n", score)
	} else if score >= 7 {
		fmt.Printf("Score %.1f: Good\n", score)
	} else if score >= 5 {
		fmt.Printf("Score %.1f: Regular\n", score)
	} else {
		fmt.Printf("Score %.1f: Insufficient\n", score)
	}
}

// switch with value
func checkDayWithSwitch(day int) {
	switch day {
	case 1:
		fmt.Println("Day 1: Sunday")
	case 2:
		fmt.Println("Day 2: Monday")
	case 3:
		fmt.Println("Day 3: Tuesday")
	case 4:
		fmt.Println("Day 4: Wednesday")
	case 5:
		fmt.Println("Day 5: Thursday")
	case 6:
		fmt.Println("Day 6: Friday")
	case 7:
		fmt.Println("Day 7: Saturday")
	default:
		fmt.Println("Invalid day")
	}
}

// switch true (like if-else chain)
func checkGradeWithSwitchTrue(score float64) {
	switch {
	case score >= 9:
		fmt.Printf("Score %.1f: Excellent\n", score)
	case score >= 7:
		fmt.Printf("Score %.1f: Good\n", score)
	case score >= 5:
		fmt.Printf("Score %.1f: Regular\n", score)
	default:
		fmt.Printf("Score %.1f: Insufficient\n", score)
	}
}
