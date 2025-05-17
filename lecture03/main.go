package main

import (
	"fmt"
	"strings"
)

const baseText = "GO lang Lecture 3 about Strings"

func main() {
	printLectureTitle()
	printOriginal()
	printUpperCase()
	printLowerCase()
	printLength()
	printContainsWord("Lecture")
	printStartsWith("GO")
	printEndsWith("Strings")
	printWordCount()
	printReplaceWord("Lecture", "Session")
	printSplitWords()
	printJoinWords([]string{"Learning", "Go", "is", "fun"})
	printIndexOfWord("about")
	printTrimSpaces("   Hello Go!   ")
	printCompareStrings("GoLang", "golang")
}

func printLectureTitle() {
	fmt.Println("Lecture 3 - Working with Strings in Go")
}

func printOriginal() {
	fmt.Println("Original:", baseText)
}

func printUpperCase() {
	fmt.Println("Upper case:", strings.ToUpper(baseText))
}

func printLowerCase() {
	fmt.Println("Lower case:", strings.ToLower(baseText))
}

func printLength() {
	fmt.Println("Length:", len(baseText))
}

func printContainsWord(word string) {
	fmt.Printf("Contains %s %v", word, strings.Contains(baseText, word))
}

func printStartsWith(prefix string) {
	fmt.Printf("Starts with %s %v", prefix, strings.HasPrefix(baseText, prefix))
}

func printEndsWith(suffix string) {
	fmt.Printf("Ends with %s %v", suffix, strings.HasSuffix(baseText, suffix))
}

func printWordCount() {
	words := strings.Fields(baseText)
	fmt.Println("Word count:", len(words))
}

func printReplaceWord(old, new string) {
	replaced := strings.ReplaceAll(baseText, old, new)
	fmt.Println("After replacement:", replaced)
}

func printSplitWords() {
	words := strings.Fields(baseText)
	fmt.Println("Split words:", words)
}

func printJoinWords(words []string) {
	joined := strings.Join(words, " ")
	fmt.Println("Joined words:", joined)
}

func printIndexOfWord(word string) {
	index := strings.Index(baseText, word)
	fmt.Printf("Index of %s: %d", word, index)
}

func printTrimSpaces(text string) {
	trimmed := strings.TrimSpace(text)
	fmt.Println("Trimmed text:", trimmed)
}

func printCompareStrings(a, b string) {
	result := strings.Compare(a, b)
	equalInsensitive := strings.EqualFold(a, b)
	fmt.Printf("Compare (case-sensitive): %d", result)
	fmt.Printf("Equal (case-insensitive): %v", equalInsensitive)
}
