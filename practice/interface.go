package main
import "fmt"

type employee interface {
	EmployeeInfo() string
}

type fullTimeEmployee struct {
	name string
	salary int
}

func (f fullTimeEmployee) EmployeeInfo() string {
	return fmt.Sprintf("Employee name: %s and earns %v USD per year.\n", f.name, f.salary)
}

type contractorEmployee struct {
	name string
	hourlyPay int
	hoursPerYear int
}

func (c contractorEmployee) EmployeeInfo() string {
	return fmt.Sprintf("Employee name: %s and earns %v USD per year.\n", c.name, c.hourlyPay * c.hoursPerYear)
}

func getEmployeeInfo(e employee) string {
	return e.EmployeeInfo()
}

func main() {
	FullTimeEmployee := fullTimeEmployee{
		name: "Mohammad Arkan",
		salary: 60015,
	}

	ContractorEmployee := contractorEmployee{
		name: "Jovan Ahmed",
		hourlyPay: 12,
		hoursPerYear: 993,
	}

	fmt.Println(getEmployeeInfo(FullTimeEmployee))
	fmt.Println(getEmployeeInfo(ContractorEmployee))
}

