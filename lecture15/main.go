package main

import (
	"fmt"
	"strings"
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
	Street string
	City   string
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

type Team struct {
	Name    string
	Members []Member
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {
	u := User{"Lucas", 17}
	u.Greet()
	fmt.Println("Is adult?", u.IsAdult())
	u.IncreaseAge()
	fmt.Println("After birthday:", u.Age)

	p := Product{"Pen", 2.5, 3}
	fmt.Println("Total price:", p.TotalPrice())
	p.Discount(10)
	fmt.Printf("After discount: %.2f ", p.Price)

	a := Address{"Av. Brasil", "São Paulo"}
	fmt.Println("Full address:", a.FullAddress())

	dev := Developer{"Alice", []string{"Go", "Docker"}}
	dev.ListSkills()
	dev.AddSkill("Kubernetes")
	dev.ListSkills()

	m := Member{"João", "Intern"}
	m.ChangeRole("Developer")
	fmt.Printf("Member after change: %+v ", m)

	pr := Project{"Go App", map[string]string{"env": "prod"}}
	fmt.Println("Tag summary:", pr.TagSummary())
	pr.AddTag("version", "1.0")
	fmt.Println("Updated tags:", pr.Tags)

	other := User{"Lucas", 18}
	fmt.Println("Same user?", u.Compare(other))

	t := Team{"DevTeam", []Member{{"Ana", "QA"}}}
	t.ListMembers()
	t.AddMember(Member{"Leo", "DevOps"})
	t.ListMembers()
	fmt.Println("Team size:", t.MemberCount())

	r := Rectangle{4, 6}
	fmt.Println("Area:", r.Area())
	fmt.Println("Perimeter:", r.Perimeter())

	b := Book{"Go Basics", "Maddo", 120}
	fmt.Println("Book summary:", b.Summary())
	b.UpdateTitle("Go Advanced")
	fmt.Println("Updated title:", b.Title)

	b2 := Book{"Go Advanced", "Maddo", 130}
	fmt.Println("Same title?", b.CompareTitle(b2))
}

// --- USER METHODS ---
func (u User) Greet() {
	fmt.Printf("Hello, %s! ", u.Name)
}

func (u *User) IncreaseAge() {
	u.Age++
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}

func (u User) Compare(other User) bool {
	return u.Name == other.Name && u.Age == other.Age
}

// --- PRODUCT METHODS ---
func (p Product) TotalPrice() float64 {
	return p.Price * float64(p.Quantity)
}

func (p *Product) Discount(percent float64) {
	p.Price -= p.Price * percent / 100
}

// --- ADDRESS METHOD ---
func (a Address) FullAddress() string {
	return a.Street + ", " + a.City
}

// --- DEVELOPER METHODS ---
func (d Developer) ListSkills() {
	fmt.Println("Skills:", strings.Join(d.Skills, ", "))
}

func (d *Developer) AddSkill(skill string) {
	d.Skills = append(d.Skills, skill)
}

// --- MEMBER METHOD ---
func (m *Member) ChangeRole(newRole string) {
	m.Role = newRole
}

// --- PROJECT METHODS ---
func (p Project) TagSummary() string {
	result := ""
	for k, v := range p.Tags {
		result += k + "=" + v + "; "
	}
	return result
}

func (p *Project) AddTag(key, value string) {
	p.Tags[key] = value
}

// --- TEAM METHODS ---
func (t Team) ListMembers() {
	for _, m := range t.Members {
		fmt.Printf("%s (%s) ", m.Name, m.Role)
	}
}

func (t *Team) AddMember(member Member) {
	t.Members = append(t.Members, member)
}

func (t Team) MemberCount() int {
	return len(t.Members)
}

// --- RECTANGLE METHODS ---
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// --- BOOK METHODS ---
func (b Book) Summary() string {
	return fmt.Sprintf("%s by %s, %d pages", b.Title, b.Author, b.Pages)
}

func (b *Book) UpdateTitle(title string) {
	b.Title = title
}

func (b Book) CompareTitle(other Book) bool {
	return b.Title == other.Title
}
