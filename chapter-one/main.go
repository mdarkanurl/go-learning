package main
import "fmt"

func main() {
	// Print "ShobApp server is running"
	fmt.Println("ShobApp server is running...")

	// fix a bug
	massagesFromAli := []string {
		"What the hell are you doing???",
		"Did the complete the task",
		"I'm gonna tell to your mom something",
		"Are you a programmer",
	}

	numMessages := float64(len(massagesFromAli))
	costPerMessage := .02

	totalCost := costPerMessage * numMessages
	fmt.Printf("Ali spent %.2f on text messages today\n", totalCost)
}
