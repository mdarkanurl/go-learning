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

type email struct {
	isSubscribed bool
	body string
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return 0.05 * float64(len(e.body))
	}

	return 0.01 * float64(len(e.body))
}

func (e email) print() {
	fmt.Println(e.body)
}

type expense interface {
	cost() float64
}

type pinter interface {
	print()
}

func print(p pinter) {
	p.print()
}

func test(e expense, p pinter) {
	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
	p.print()
}

// Using type assertion,
// determine what type of value was passed and return an appropriate message.

func describe(value interface{}) {
	if v, ok := value.(string); ok {
		fmt.Println("This is a string and here's the value:", v)
		return
	}

	if v, ok := value.(int); ok {
		fmt.Println("This is a integer and here's the value:", v)
		return
	}

	if v, ok := value.(float64); ok {
		fmt.Println("This is a float64 and here's the value:", v)
		return
	}

	if v, ok := value.(bool); ok {
		fmt.Println("This is a bool and here's the value:", v)
		return
	}

	fmt.Println("Unknown type")
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

	subscribedEmail := email{
		isSubscribed: true,
		body: "Hey! This is from subscribed user",
	}

	test(subscribedEmail, subscribedEmail)

	notSubscribedEmail := email{
		isSubscribed: false,
		body: "Hey! This is from non-subscribed user",
	}

	test(notSubscribedEmail, notSubscribedEmail)

	describe("Go")
	describe(42)
	describe(3.14)
	describe(true)
	describe([]int{1, 2, 3})
}
