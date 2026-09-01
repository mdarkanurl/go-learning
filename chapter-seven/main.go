package main
import "fmt"

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i % 3 == 0 {
			fmt.Println("fizz")
		} else if i % 5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// Write a function that takes an array of ages and returns the number of adults between the ages of 18 and 30.
func checkAge(ages [4]int) int {
	count := 0
	for _, value := range ages {
		if value < 18 {
			continue
		}

		count++

		if value > 30 {
			break
		}
	}

	return count
}

func main() {
	// print 0 through 4
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 1
	for i < 5 {
		fmt.Println(i)
		i++
	}

	count := 1
	for {
		fmt.Println("Count:", count)
		count++

		if count > 5 {
			break
		}
	}

	sum := 0
    num := 1
    for num <= 5 {
        sum += num
        num++
    }
    fmt.Println("Sum:", sum)

	numbers := []int{10, 20, 30, 40}
    for index, value := range numbers {

        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

	str := "Hello"

    for index, char := range str {

        fmt.Printf("Index: %d, Character: %c\n", index, char)
    }

	students := map[int]string{
        101: "Arkan",
        102: "Ali",
        103: "Mohammad",
    }

    for key, value := range students {

        fmt.Printf("ID: %d, Name: %s\n", key, value)
    }

	fizzbuzz()
	ages := [4]int{20, 17, 30, 25}
	adults := checkAge(ages)
    fmt.Println("There are", adults, "adults")
}

