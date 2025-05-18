package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Struct FileManager for POO encapsulation
type FileManager struct{}

// List current directory contents
func (fm FileManager) ListCurrentDirectory() ([]string, error) {
	entries, err := os.ReadDir(".")
	if err != nil {
		return nil, err
	}
	var result []string
	for _, entry := range entries {
		result = append(result, entry.Name())
	}
	return result, nil
}

// List all files recursively
func (fm FileManager) ListAllFilesRecursively(path string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		files = append(files, p)
		return nil
	})
	return files, err
}

// Print path examples per OS
func PrintPlatformPathExamples() {
	fmt.Println("📁 Path Examples:")
	fmt.Println("Linux/Mac: /home/user/folder or ./")
	fmt.Println("Windows  : C:\\Users\\user\\folder or .\\")
}

// Save list of lectures to a file
func (fm FileManager) SaveLecturesToFile(filename string) error {
	lectures := []string{
		"LECTURE 1: Go Language Fundamentals",
		"LECTURE 2: Floating Point Numbers",
		"LECTURE 3: String Manipulation",
		"LECTURE 4: Numeric Data Types",
		"LECTURE 5: Decision Making",
		"LECTURE 6: Loops and Repetition",
		"LECTURE 7: Reading Input",
		"LECTURE 8: Arrays and Slices",
		"LECTURE 9: Operators",
		"LECTURE 10: Maps",
		"LECTURE 11: Functions with Return",
		"LECTURE 12: Advanced Functions",
		"LECTURE 13: Pointers",
		"LECTURE 14: Structs",
		"LECTURE 15: Methods",
		"LECTURE 16: OOP in Go",
		"LECTURE 17: Error Handling",
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lectures {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

// Read and print lines from file
func (fm FileManager) ReadFileLines(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Printf("📄 Contents of %s:\n", filepath)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return scanner.Err()
}

// Write custom text to file
func (fm FileManager) WriteCustomTextToFile(text, filename string) error {
	return os.WriteFile(filename, []byte(text), 0644)
}

// Get file info
func (fm FileManager) ReadFileInfo(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	fmt.Printf("📘 File Info:\nName: %s\nSize: %d bytes\nModified: %s\n",
		info.Name(), info.Size(), info.ModTime().Format(time.RFC822))
	return nil
}

// Group files by extension
func (fm FileManager) GroupFilesByExtension(path string) (map[string][]string, error) {
	extMap := make(map[string][]string)
	err := filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			ext := filepath.Ext(d.Name())
			extMap[ext] = append(extMap[ext], p)
		}
		return nil
	})
	return extMap, err
}

// Create directory if it doesn't exist
func (fm FileManager) EnsureDirectory(path string) error {
	if path == "" {
		return errors.New("path cannot be empty")
	}
	return os.MkdirAll(path, 0755)
}

func main() {
	fm := FileManager{}

	// Print example paths
	PrintPlatformPathExamples()

	// Save lectures to file
	err := fm.SaveLecturesToFile("go_lectures.txt")
	if err != nil {
		fmt.Println("❌ Error saving lectures:", err)
	} else {
		fmt.Println("✅ Lectures saved to go_lectures.txt")
	}

	// Read current directory
	files, err := fm.ListCurrentDirectory()
	if err != nil {
		fmt.Println("❌ Error listing current dir:", err)
	} else {
		fmt.Println("📂 Current directory files:")
		for _, f := range files {
			fmt.Println(" -", f)
		}
	}

	// List all files recursively from user input
	fmt.Print("Enter path to list recursively: ")
	reader := bufio.NewReader(os.Stdin)
	userPath, _ := reader.ReadString('\n')
	userPath = strings.TrimSpace(userPath)

	allFiles, err := fm.ListAllFilesRecursively(userPath)
	if err != nil {
		fmt.Println("❌ Error listing files:", err)
	} else {
		fmt.Printf("📁 %d files found under '%s':\n", len(allFiles), userPath)
		for _, file := range allFiles {
			fmt.Println(" •", file)
		}
	}
}
