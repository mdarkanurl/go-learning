package main

import (
	"fmt"
	"strings"
)

func changeAge(age *int) {
	*age = 19
}

func replaceNames(name string, message *string) {
	*message = strings.ReplaceAll(*message, "NAME", name)
}

type UserInfo struct {
	id           int
	username     string
	reqPerSecond int
}

func (u *UserInfo) resetReatLimit() {
	if u.reqPerSecond > 200 {
		u.reqPerSecond = 0
	}

	u.reqPerSecond++
}

func main() {
	applesPrice := 360.00

	pointerOfApplesPrice := &applesPrice
	fmt.Println("value of applesPrice var", applesPrice)
	*pointerOfApplesPrice = 350.00
	fmt.Println("Memory address of applesPrice var", pointerOfApplesPrice)
	fmt.Println("Changed value of applesPrice var", applesPrice)

	// Change the age value using pointer
	age := 18
	changeAge(&age)
	fmt.Println("Age becomes:", age)

	message1 := "Hello, this is NAME"
	message2 := "I'm NAME from Bangladesh"
	replaceNames("Arkan", &message1)
	replaceNames("Jovan", &message2)
	fmt.Println(message1)
	fmt.Println(message2)

	userInfo := UserInfo{
		id:           6081,
		username:     "thearkan",
		reqPerSecond: 201,
	}

	userInfo.resetReatLimit()
	fmt.Println(userInfo)
}
