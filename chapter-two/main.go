package main
import "fmt"

func main() {
	// Types in Go
	var username string = "mdarkanurl"
	var password string = "arkan1020"

	fmt.Println("Auth: basic", username + ":" + password)

	var mynum int8 = 127 // int8 can't hold more than 127 number value
	fmt.Println(mynum)

	var smsSendingLimit int
	var constPerSMS float64
	var hasPermission bool
	var name string

	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimit,
		constPerSMS,
		hasPermission,
		name,
	)

	age := 18 // This is the way that helps you infer types
	fmt.Println(age)

	penniesPerText := 2.0
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)

	// Different way of declaring variable in go
	averageOpenRate, displayMessage := 0.16, "is the average open rate of your mail"
	fmt.Println(averageOpenRate, displayMessage)

	// constants in go
	const myName = "Arkan"; // we can't change the value of this variable
	fmt.Println(myName);

	const secondsInMinute = 60
	const minuteInHour = 60
	const secondsInHour = secondsInMinute * minuteInHour

	fmt.Printf("number of seconds in an hour: %v\n", secondsInHour)

	// convert number
	myAge := 18.85

	myAgeInt := int(myAge)
	fmt.Println(myAgeInt)

	// Conditionals
	messagesLen := 20
	maxMessageLen := 20

	if messagesLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message can't be sent")
	}
}
