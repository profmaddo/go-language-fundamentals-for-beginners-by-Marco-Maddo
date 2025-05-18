package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "go_database.db"

type Person struct {
	ID   int
	Name string
	Age  int
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createPersonTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS Person (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL
	);
	`
	_, err := db.Exec(query)
	return err
}

func insertPerson(db *sql.DB, name string, age int) error {
	stmt, err := db.Prepare("INSERT INTO Person(name, age) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, age)
	return err
}

func listPersons(db *sql.DB) error {
	rows, err := db.Query("SELECT id, name, age FROM Person")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.ID, &p.Name, &p.Age)
		if err != nil {
			return err
		}
		fmt.Printf("ID: %d | Name: %s | Age: %d ", p.ID, p.Name, p.Age)
	}
	return nil
}

func updatePersonByID(db *sql.DB, id int, name string, age int) error {
	stmt, err := db.Prepare("UPDATE Person SET name = ?, age = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, age, id)
	return err
}

func deletePersonByID(db *sql.DB, id int) error {
	stmt, err := db.Prepare("DELETE FROM Person WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}

func findPersonByID(db *sql.DB, id int) error {
	row := db.QueryRow("SELECT id, name, age FROM Person WHERE id = ?", id)
	var p Person
	err := row.Scan(&p.ID, &p.Name, &p.Age)
	if err == sql.ErrNoRows {
		fmt.Println("No person found with that ID.")
		return nil
	} else if err != nil {
		return err
	}
	fmt.Printf("ID: %d | Name: %s | Age: %d ", p.ID, p.Name, p.Age)
	return nil
}

func menuPerson(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("--- PERSON MENU ---")
		fmt.Println("1 - Insert Person")
		fmt.Println("2 - List Persons")
		fmt.Println("3 - Update Person by ID")
		fmt.Println("4 - Delete Person by ID")
		fmt.Println("5 - Find Person by ID")
		fmt.Println("0 - Back to Main Menu")
		fmt.Print("Choose option: ")

		optStr, _ := reader.ReadString('\n')
		optStr = strings.TrimSpace(optStr)
		option, _ := strconv.Atoi(optStr)

		switch option {
		case 1:
			fmt.Print("Enter name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Print("Enter age: ")
			ageStr, _ := reader.ReadString('\n')
			ageStr = strings.TrimSpace(ageStr)
			age, _ := strconv.Atoi(ageStr)
			err := insertPerson(db, name, age)
			if err != nil {
				fmt.Println("Insert failed:", err)
			} else {
				fmt.Println("Person inserted.")
			}
		case 2:
			_ = listPersons(db)
		case 3:
			fmt.Print("Enter ID to update: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			fmt.Print("New name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Print("New age: ")
			ageStr, _ := reader.ReadString('\n')
			age, _ := strconv.Atoi(strings.TrimSpace(ageStr))
			err := updatePersonByID(db, id, name, age)
			if err != nil {
				fmt.Println("Update failed:", err)
			} else {
				fmt.Println("Person updated.")
			}
		case 4:
			fmt.Print("Enter ID to delete: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			err := deletePersonByID(db, id)
			if err != nil {
				fmt.Println("Delete failed:", err)
			} else {
				fmt.Println("Person deleted.")
			}
		case 5:
			fmt.Print("Enter ID to search: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			_ = findPersonByID(db, id)
		case 0:
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Println("Failed to connect to DB:", err)
		return
	}
	defer db.Close()

	err = createPersonTable(db)
	if err != nil {
		fmt.Println("Failed to create table:", err)
		return
	}

	menuPerson(db)
}
