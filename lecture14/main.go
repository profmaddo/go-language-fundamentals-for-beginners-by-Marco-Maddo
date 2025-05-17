package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

type Address struct {
	City  string
	State string
}

type Developer struct {
	Name   string
	Skills []string
}

type Member struct {
	Name string
	Role string
}

type Project struct {
	Title string
	Tags  map[string]string
}

func main() {
	createUser()
	updateUserAge()
	compareUsers()
	printUserDetails(User{"Ana", 22})
	u := returnUser()
	fmt.Printf("Returned user: %+v ", u)
	createProduct()
	calculateTotalPrice(Product{"Book", 19.90, 3})
	createAddressEmbedded()
	createAnonymousStruct()
	pointerToStruct()
	structWithSliceField()
	addSkillToDeveloper()
	employeeFullName()
	createTeam()
	projectWithMapField()
	iterateTeamMembers()
	compareStructValues()
	sortUsersByAge()
	deepCopyStruct()
	updateStructInFunction(&User{Name: "Mario", Age: 28})
}

func createUser() {
	user := User{"John", 30}
	fmt.Printf("Created user: %+v ", user)
}

func updateUserAge() {
	user := User{"Emma", 25}
	user.Age += 1
	fmt.Printf("Updated user age: %+v ", user)
}

func compareUsers() {
	a := User{"Liam", 20}
	b := User{"Liam", 20}
	fmt.Println("Users equal?", a == b)
}

func printUserDetails(u User) {
	fmt.Printf("User details: %+v ", u)
}

func returnUser() User {
	return User{"Julia", 18}
}

func createProduct() {
	p := Product{"Laptop", 3500.00, 2}
	fmt.Printf("Product: %+v ", p)
}

func calculateTotalPrice(p Product) {
	total := p.Price * float64(p.Quantity)
	fmt.Printf("Total price: %.2f ", total)
}

func createAddressEmbedded() {
	type Person struct {
		Name    string
		Address Address
	}
	p := Person{"Carlos", Address{"Curitiba", "PR"}}
	fmt.Printf("Person with address: %+v ", p)
}

func createAnonymousStruct() {
	item := struct {
		ID    int
		Title string
	}{1, "Notebook"}
	fmt.Printf("Anonymous struct: %+v ", item)
}

func pointerToStruct() {
	u := &User{"Laura", 27}
	u.Age++
	fmt.Printf("Pointer to struct: %+v ", *u)
}

func structWithSliceField() {
	dev := Developer{"Rob", []string{"Go", "Docker"}}
	fmt.Printf("Developer: %+v ", dev)
}

func addSkillToDeveloper() {
	dev := Developer{"Alice", []string{"HTML"}}
	dev.Skills = append(dev.Skills, "CSS")
	fmt.Printf("Updated skills: %+v ", dev)
}

func employeeFullName() {
	type Employee struct {
		FirstName string
		LastName  string
	}
	fullName := func(e Employee) string {
		return e.FirstName + " " + e.LastName
	}
	e := Employee{"Marcos", "Silva"}
	fmt.Println("Full name:", fullName(e))
}

func createTeam() {
	team := []Member{
		{"Ana", "Dev"},
		{"Paulo", "Tester"},
	}
	fmt.Println("Team:", team)
}

func projectWithMapField() {
	p := Project{"Go Web App", map[string]string{"status": "active", "owner": "teamA"}}
	fmt.Printf("Project: %+v ", p)
}

func iterateTeamMembers() {
	team := []Member{
		{"Jonas", "DevOps"},
		{"Sofia", "Designer"},
	}
	for _, m := range team {
		fmt.Printf("Member: %s - %s ", m.Name, m.Role)
	}
}

func compareStructValues() {
	a := Product{"Pen", 2.5, 1}
	b := Product{"Pen", 2.5, 1}
	fmt.Println("Products equal?", a == b)
}

func sortUsersByAge() {
	users := []User{
		{"Lana", 22},
		{"Bea", 29},
		{"Mark", 21},
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})
	fmt.Println("Sorted by age:", users)
}

func deepCopyStruct() {
	original := User{"Nina", 30}
	copy := original
	copy.Name = "Nina Clone"
	fmt.Printf("Original: %+v ", original)
	fmt.Printf("Copy: %+v ", copy)
}

func updateStructInFunction(u *User) {
	u.Age += 10
	fmt.Printf("Updated via function: %+v ", u)
}
