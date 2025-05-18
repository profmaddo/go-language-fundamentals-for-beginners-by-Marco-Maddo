package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// LECTURE 3 – STRING FUNCTIONS WITH ERROR
func toUpper(s string) (string, error) {
	if s == "" {
		return "", errors.New("cannot convert empty string to uppercase")
	}
	return strings.ToUpper(s), nil
}

func containsWord(base, word string) (bool, error) {
	if word == "" {
		return false, errors.New("word to search cannot be empty")
	}
	return strings.Contains(base, word), nil
}

func replaceWord(base, old, new string) (string, error) {
	if old == "" {
		return "", errors.New("old word cannot be empty")
	}
	return strings.ReplaceAll(base, old, new), nil
}

func indexOfWord(base, word string) (int, error) {
	index := strings.Index(base, word)
	if index == -1 {
		return -1, fmt.Errorf("word '%s' not found", word)
	}
	return index, nil
}

// LECTURE 4 – NUMERIC EXAMPLES WITH ERROR
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func subtractUint16(a, b uint16) (uint16, error) {
	if b > a {
		return 0, errors.New("cannot subtract to negative in uint16")
	}
	return a - b, nil
}

func areaRectangle(w, h float64) (float64, error) {
	if w < 0 || h < 0 {
		return 0, errors.New("dimensions must be non-negative")
	}
	return w * h, nil
}

func validateComplex(real, imag float64) (complex128, error) {
	if math.IsNaN(real) || math.IsNaN(imag) {
		return 0, errors.New("real or imaginary part is NaN")
	}
	return complex(real, imag), nil
}

// LECTURE 7 – INPUT VALIDATION
func validateAgeInput(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}
	return nil
}

func validateOption(option int) error {
	if option < 1 || option > 3 {
		return errors.New("invalid option: must be 1, 2 or 3")
	}
	return nil
}

// LECTURE 8 – SLICE AND ARRAY
func sumSlice(slice []int) (int, error) {
	if slice == nil {
		return 0, errors.New("slice is nil")
	}
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum, nil
}

func searchInSlice(slice []int, target int) (bool, error) {
	for _, v := range slice {
		if v == target {
			return true, nil
		}
	}
	return false, fmt.Errorf("value %d not found in slice", target)
}

func removeItemByIndex(slice []string, index int) ([]string, error) {
	if index < 0 || index >= len(slice) {
		return nil, errors.New("index out of bounds")
	}
	return append(slice[:index], slice[index+1:]...), nil
}

// LECTURE 10 – MAP
func getFromMap(m map[string]string, key string) (string, error) {
	val, ok := m[key]
	if !ok {
		return "", fmt.Errorf("key '%s' not found", key)
	}
	return val, nil
}

func updateInMap(m map[string]int, key string, value int) error {
	if _, ok := m[key]; !ok {
		return fmt.Errorf("cannot update: key '%s' not found", key)
	}
	m[key] = value
	return nil
}

func deleteFromMap(m map[string]string, key string) error {
	if _, ok := m[key]; !ok {
		return fmt.Errorf("cannot delete: key '%s' not found", key)
	}
	delete(m, key)
	return nil
}

// LECTURE 11 – FUNCTION RETURNS
func repeatWord(word string, times int) (string, error) {
	if word == "" || times <= 0 {
		return "", errors.New("invalid word or repetition count")
	}
	return strings.Repeat(word, times), nil
}

func getFirstAndLast(slice []int) (int, int, error) {
	if len(slice) == 0 {
		return 0, 0, errors.New("slice is empty")
	}
	return slice[0], slice[len(slice)-1], nil
}

func main() {
	// String example with error
	upper, err := toUpper("go")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Upper:", upper)
	}

	// Division example
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Division error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Map read example
	myMap := map[string]string{"lang": "Go"}
	val, err := getFromMap(myMap, "version")
	if err != nil {
		fmt.Println("Map error:", err)
	} else {
		fmt.Println("Value from map:", val)
	}

	// Slice sum example
	numbers := []int{1, 2, 3}
	sum, err := sumSlice(numbers)
	if err != nil {
		fmt.Println("Sum error:", err)
	} else {
		fmt.Println("Sum:", sum)
	}
}
