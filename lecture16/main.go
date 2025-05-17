package main

import (
	"fmt"
	"math"
)

// --- ABSTRACTION ---
type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) Person {
	return Person{name: name, age: age}
}

func (p Person) Introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old. ", p.name, p.age)
}

// --- ENCAPSULATION ---
func (p *Person) SetAge(age int) {
	if age > 0 {
		p.age = age
	}
}

func (p Person) GetAge() int {
	return p.age
}

// --- INHERITANCE via COMPOSITION ---
type Employee struct {
	Person
	position string
}

func (e Employee) Introduce() {
	fmt.Printf("I'm %s and I work as a %s. ", e.name, e.position)
}

type Manager struct {
	Employee
	department string
}

func (m Manager) Introduce() {
	fmt.Printf("I'm %s, manager of %s department. ", m.name, m.department)
}

// --- POLYMORPHISM via INTERFACE ---
type Printer interface {
	Introduce()
}

func printIntroduction(p Printer) {
	p.Introduce()
}

// --- ENCAPSULATION EXAMPLE ---
type AuthService struct {
	users map[string]string // username:password
}

func NewAuthService() AuthService {
	return AuthService{users: map[string]string{"admin": "1234"}}
}

func (a AuthService) Login(user, pass string) bool {
	if pw, ok := a.users[user]; ok {
		return pw == pass
	}
	return false
}

// --- POLYMORPHISM with Shape interface ---
type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func describeShape(s Shape) {
	fmt.Printf("Shape area: %.2f ", s.Area())
}

// --- ANIMAL INTERFACE EXAMPLE ---
type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (Dog) Speak() string {
	return "Woof!"
}

func (Cat) Speak() string {
	return "Meow!"
}

func animalTalk(a Animal) {
	fmt.Println("Animal says:", a.Speak())
}

// --- COMPOSITION EXAMPLE ---
type Engine struct {
	powered bool
}

func (e *Engine) Start() {
	e.powered = true
	fmt.Println("Engine started.")
}

func (e *Engine) Stop() {
	e.powered = false
	fmt.Println("Engine stopped.")
}

type Car struct {
	brand  string
	Engine // Embedded
}

func (c Car) Drive() {
	if c.powered {
		fmt.Printf("%s is driving! ", c.brand)
	} else {
		fmt.Printf("%s won't move without starting the engine. ", c.brand)
	}
}

// --- MAIN ---
func main() {
	// Abstraction
	person := NewPerson("Lucas", 21)
	person.Introduce()

	// Encapsulation
	person.SetAge(30)
	fmt.Println("Updated age:", person.GetAge())

	// Composition (Inheritance) & Polymorphism
	emp := Employee{Person{"Ana", 28}, "Engineer"}
	mgr := Manager{Employee{Person{"Carlos", 40}, "Senior Engineer"}, "IT"}

	printIntroduction(person) // Abstraction
	printIntroduction(emp)    // Polymorphism
	printIntroduction(mgr)    // Polymorphism

	// Encapsulation in service
	auth := NewAuthService()
	fmt.Println("Login admin/1234:", auth.Login("admin", "1234"))
	fmt.Println("Login admin/wrong:", auth.Login("admin", "wrong"))

	// Polymorphism with interfaces
	r := Rectangle{5, 3}
	c := Circle{2}
	describeShape(r)
	describeShape(c)

	// Animal interface
	dog := Dog{}
	cat := Cat{}
	animalTalk(dog)
	animalTalk(cat)

	// Composition in action
	car := Car{"Fusca", Engine{}}
	car.Drive()
	car.Start()
	car.Drive()
	car.Stop()
}
