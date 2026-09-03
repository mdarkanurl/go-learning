package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, value := range nums {
		total += value
	}
	return total
}

func main() {

	// Declare an array of 5 integers
	var numbers [5]int
	fmt.Println("Zero-valued array:", numbers)

	// Declare and initialize with values
	words := [3]string{"hello", "world", "go"}
	fmt.Println("String array:", words)

	// Let the compiler count the elements
	fruits := [...]string{"apple", "banana", "cherry", "date"}
	fmt.Println("Fruits:", fruits)

	// Accessing and Modifying Elements
	// Use index (0-based) to access elements
	fmt.Println("First word:", words[0])
	fmt.Println("Last fruit:", fruits[3])

	// Modify an element
	words[1] = "WORLD"
	fmt.Println("Modified words:", words)

	// Array Length
	fmt.Println("Length of numbers:", len(numbers))
	fmt.Println("Length of fruits:", len(fruits))

	// Iterating Over Arrays
	fmt.Println("\nIterating with for loop:")
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("  fruits[%d] = %s\n", i, fruits[i])
	}

	fmt.Println("\nIterating with range:")
	for index, value := range fruits {
		fmt.Printf("  Index: %d, Value: %s\n", index, value)
	}

	// Ignore index with blank identifier
	fmt.Println("\nValues only:")
	for _, fruit := range fruits {
		fmt.Printf("  %s\n", fruit)
	}

	// Array Comparison
	// Two arrays are equal if they have the same type AND same elements
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 4}
	fmt.Println("\narr1 == arr2:", arr1 == arr2)
	fmt.Println("arr1 == arr3:", arr1 == arr3)

	// Multidimensional Arrays
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("\n2D Array (matrix):", matrix)
	fmt.Println("Element at [1][2]:", matrix[1][2])

	fmt.Println("Matrix contents:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("  matrix[%d][%d] = %d\n", i, j, matrix[i][j])
		}
	}

	colors := [3]string{"red", "green", "blue"}
	printArray(colors)

	// Partial Initialization
	partial := [5]int{10, 20}
	fmt.Println("\nPartial init:", partial)

	// Indexed Initialization
	indexed := [5]int{2: 100, 4: 400}
	fmt.Println("Indexed init:", indexed)

	ages := [5]int{18, 20, 17, 23, 77}
	mySlice := ages[1:4]
	fmt.Println(mySlice)

	// create slice using make function
	countrySlice := make([]string, 2, 5)
	fmt.Println(countrySlice)

	// Modify or insert data to existing slots
	countrySlice[0] = "Bangladesh"
	countrySlice[1] = "UAS"
	fmt.Println(countrySlice)

	// Append new slots
	countrySlice = append(countrySlice, "India")
	fmt.Println(countrySlice)

	// copy a slice
	cSlice := make([]string, len(countrySlice))
	copy(cSlice, countrySlice)
	fmt.Println("Copyed Slice:", cSlice)

	// variadic feture in go
	numbersForSum := []int{10, 11, 2}
	fmt.Println(sum(numbersForSum...))

	fmt.Println("Memory location of first item in numbersForSum:", &numbersForSum[0])
	fmt.Println("Memory location of second item in numbersForSum:", &numbersForSum[1])
}

// printArray demonstrates passing arrays to functions
func printArray(arr [3]string) {
	fmt.Println("\nInside printArray:")
	for _, val := range arr {
		fmt.Printf("  %s\n", val)
	}
}
