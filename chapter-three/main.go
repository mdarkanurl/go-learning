package main
import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func add(num1, num2 int) int {
	return num1 + num2
}

func incrementSends(x int, y int) int {
	return x + y
}

func profile() (string, int) {
	return "Arkan", 18
}

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	if yearsUntilAdult = 18 - age; yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}

	if yearsUntilDrinking = 21 - age; yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}

	if yearsUntilCarRental = 25 - age; yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}

	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
}

func main() {
	// Functions
	fmt.Println(concat("Arkan,", " he is from Bangladesh!"))
	fmt.Println(add(7, 11))

	// memory
	sendsSoFar := 430
	const sendsToAdd = 25
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("You've sent", sendsSoFar, "messages")

	name, _ := profile()
	fmt.Println("Welcome to ShobApp24,",  name)

	fmt.Println(yearsUntilEvents(21))
}
