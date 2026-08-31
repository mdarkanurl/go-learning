package main
import "fmt"

func sendMessageToPeople(sendToTeachers, sendToStudents string) (int, error) {
	constForTeacher, err := sendMessage(sendToTeachers)
	if err != nil {
		return 0, err
	}

	constForStudent, err := sendMessage(sendToTeachers)
	if err != nil {
		return 0, err
	}

	return constForTeacher + constForStudent, nil
}

func sendMessage(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

type dbError struct {
	statusCode int
	message string
}

func (de dbError) Error() string {
	return fmt.Sprintf("db error message %s (statusCode %d)", de.message, de.statusCode)
}

func queryDatabase() error {
	return dbError{statusCode: 500, message: "Too many connection"}
}


func main() {
	fmt.Println(sendMessageToPeople("Hello this for teacher", "Hello this for student"))
	err := queryDatabase()
	if err != nil {
		fmt.Println("Error from DB connection")
	}
}

