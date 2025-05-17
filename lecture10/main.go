package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	printLectureTitle()
	createAndPrintMap()
	insertKeyValueToMap()
	readMapFromUserInput()
	getValueFromMap()
	updateMapValue()
	deleteKeyFromMap()
	checkKeyExists()
	loopThroughMap()
	countWordsInSentence()
	mapWithStructValues()
	mapKeysToSlice()
	groupItemsByCategory()
	sortMapKeys()
	compareTwoMaps()
	mapWithBoolFlags()
}

func printLectureTitle() {
	fmt.Println("Lecture 10 - Maps in Go")
}

// 1. Create and print a simple map
func createAndPrintMap() {
	capitals := map[string]string{
		"France": "Paris",
		"Japan":  "Tokyo",
	}
	fmt.Println("Country capitals:")
	for country, capital := range capitals {
		fmt.Printf("%s → %s ", country, capital)
	}
}

// 2. Insert key-value pairs
func insertKeyValueToMap() {
	ages := make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25
	fmt.Println("Inserted ages:", ages)
}

// 3. Read map from user input
func readMapFromUserInput() {
	reader := bufio.NewReader(os.Stdin)
	persons := make(map[string]string)
	fmt.Println("Enter 2 names and cities:")
	for i := 0; i < 2; i++ {
		fmt.Print("Name: ")
		name, _ := reader.ReadString('\n')
		fmt.Print("City: ")
		city, _ := reader.ReadString('\n')
		persons[strings.TrimSpace(name)] = strings.TrimSpace(city)
	}
	fmt.Println("Persons map:", persons)
}

// 4. Get value from map
func getValueFromMap() {
	grades := map[string]int{"Math": 90, "Science": 85}
	subject := "Math"
	if grade, ok := grades[subject]; ok {
		fmt.Printf("%s grade: %d ", subject, grade)
	} else {
		fmt.Println("Subject not found.")
	}
}

// 5. Update value in map
func updateMapValue() {
	scores := map[string]int{"Player1": 10}
	scores["Player1"] += 5
	fmt.Println("Updated score:", scores)
}

// 6. Delete key from map
func deleteKeyFromMap() {
	colors := map[string]string{"r": "red", "g": "green", "b": "blue"}
	delete(colors, "g")
	fmt.Println("Colors after deletion:", colors)
}

// 7. Check if key exists
func checkKeyExists() {
	users := map[string]bool{"admin": true, "guest": false}
	if _, exists := users["admin"]; exists {
		fmt.Println("Admin exists.")
	}
}

// 8. Loop through map
func loopThroughMap() {
	languages := map[string]string{"Go": "Google", "Python": "Open Source"}
	for lang, origin := range languages {
		fmt.Printf("%s was created by %s ", lang, origin)
	}
}

// 9. Count words in a sentence
func countWordsInSentence() {
	sentence := "go go learn go code"
	words := strings.Fields(sentence)
	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++
	}
	fmt.Println("Word frequencies:", frequency)
}

// 10. Map with struct values
type Product struct {
	Name  string
	Price float64
}

func mapWithStructValues() {
	products := map[string]Product{
		"P001": {"Pen", 1.99},
		"P002": {"Notebook", 3.50},
	}
	for code, p := range products {
		fmt.Printf("%s → %s: $%.2f ", code, p.Name, p.Price)
	}
}

// 11. Extract keys to slice
func mapKeysToSlice() {
	data := map[string]int{"a": 1, "b": 2}
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	fmt.Println("Keys slice:", keys)
}

// 12. Group items by category
func groupItemsByCategory() {
	catalog := map[string][]string{
		"Fruits":     {"Apple", "Banana"},
		"Vegetables": {"Carrot", "Broccoli"},
	}
	for category, items := range catalog {
		fmt.Printf("%s: %v ", category, items)
	}
}

// 13. Sort map keys
func sortMapKeys() {
	data := map[string]int{"z": 3, "a": 1, "m": 2}
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println("Sorted keys:")
	for _, k := range keys {
		fmt.Printf("%s → %d ", k, data[k])
	}
}

// 14. Compare two maps
func compareTwoMaps() {
	map1 := map[string]int{"x": 1, "y": 2}
	map2 := map[string]int{"x": 1, "y": 2}
	equal := true
	for k, v := range map1 {
		if map2[k] != v {
			equal = false
			break
		}
	}
	fmt.Println("Maps are equal?", equal)
}

// 15. Map with bool flags
func mapWithBoolFlags() {
	seen := make(map[string]bool)
	seen["apple"] = true
	seen["banana"] = false
	fmt.Println("Seen map:", seen)
}
