package main
import "fmt"

// shared method type
type Speaker interface {
	say() string
}

// Teacher's type and method
type teacher struct {
	name     string
	dialogue string
}

func (t teacher) say() string {
	return fmt.Sprintf("Teacher said: '%s'.", t.dialogue)
}

// Student's type and method
type student struct {
	name     string
	dialogue string
	roll     int
}

func (t student) say() string {
	return fmt.Sprintf("Student said: '%s'.", t.dialogue)
}

// Function that connects two type
func makeThemSpeak(s Speaker) string {
	result := s.say()
	fmt.Println(result)
	return result
}

// shared method type
type employee interface {
	getName() string
	getSalary() int
}

type fullTimeEmployee struct {
	name   string
	salary int
}

func (f fullTimeEmployee) getName() string {
	return f.name
}

func (f fullTimeEmployee) getSalary() int {
	return f.salary
}

type contractorEmployee struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractorEmployee) getName() string {
	return c.name
}

func (c contractorEmployee) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

func printEmployeeInfo(e employee) {
	fmt.Printf("%s earns %d USD per year.\n", e.getName(), e.getSalary())
}

func main() {
	Teacher := teacher{
		name:     "Ali",
		dialogue: "Where are you from?",
	}

	Student := student{
		name:     "Arkan",
		dialogue: "i'm from bangladesh",
		roll:     16,
	}

	FullTimeEmployee := fullTimeEmployee{
		name:   "Mohammad",
		salary: 60000,
	}

	ContractorEmployee := contractorEmployee{
		name:         "Kafi",
		hourlyPay:    15,
		hoursPerYear: 601,
	}

	makeThemSpeak(Teacher)
	makeThemSpeak(Student)

	printEmployeeInfo(FullTimeEmployee)
	printEmployeeInfo(ContractorEmployee)
}
