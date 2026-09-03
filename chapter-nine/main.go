package main

import "fmt"

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)

	for _, name := range names {
		if name == "" {
			continue
		}

		firstChar := rune(name[0])
		_, ok := counts[firstChar]

		if !ok {
			counts[firstChar] = make(map[string]int)
		}
		counts[firstChar][name]++
	}

	return counts
}

func main() {
	// There are two ways to create a map
	// first way use built-in make func
	peopleAges := make(map[string]int)
	peopleAges["Arkan"] = 18

	// Second way
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	fmt.Println(peopleAges)
	fmt.Println(colors)

	// get length of a map
	fmt.Println(len(colors))

	// read a signal value
	fmt.Println("Value of blue:", colors["blue"])

	// delete a value of a map
	delete(colors, "red")

	// Checking for Existence
	fruitsPrice := map[string]float64{
		"apples":       360.00,
		"Dragon Fruit": 350.00,
		"Guava":        78.00,
		"Grapes":       460.00,
	}

	value, exists := fruitsPrice["Guava"]
	if exists {
		fmt.Println("Found banana:", value, "taka")
	} else {
		fmt.Println("Banana not found")
	}

	// Iterating Over a Map
	for key, value := range colors {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	names := []string{
		"Arkan",
		"Arkan",
		"Kafi",
		"Arkan",
		"Joe",
		"Kafi",
	}

	fmt.Println(getNameCounts(names))
}
