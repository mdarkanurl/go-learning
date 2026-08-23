package main
import "fmt"

type messageToSend struct {
	message string
	phoneNumber int
}

type videoToPublish struct {
	title string
	description string
	commands videoCommands
}

type videoCommands struct {
	sender string
	content string
}

type sender struct {
	rateLimt int
	user
}

type user struct {
	username string
	id int
}

func rateLimt(sender sender) {
	fmt.Println("users' rateLimt:", sender.rateLimt)
	fmt.Println("users' username:", sender.username)
	fmt.Println("users' id:", sender.id)
}

type authenticationInfo struct {
	username string
	password string
}

func (auth authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf(
		"Authorization: Basic %s:%s",
		auth.username, auth.password,
	)
}

func testAuth(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
}

func OS() {
	// Anonymous struct
	myOS := struct {
		name string
		version int
	} {
		name: "Omarchy",
		version: 3,
	}

	fmt.Printf("My OS name %v and version is %v\n", myOS.name, myOS. version)
}

func test(m messageToSend) {
	fmt.Printf("sending message: '%s' to: %v\n", m.message, m.phoneNumber)
}

func canUploadVideos(video videoToPublish) bool {
	if video.title == "" {
		return false
	}

	if video.commands.sender == "" {
		return false
	}

	return true
}

func main() {
	test(messageToSend{
		message: "If you try, you'll be successfull",
		phoneNumber: 102030405060,
	})

	video := canUploadVideos(videoToPublish{
		title: "how to learn go",
		description: "this video helps you learn Go",
		commands: videoCommands{
			sender:  "mdarkanurl",
			content: "not bad!",
		},
	})

	if video == true {
		fmt.Println("Video is upload able.....")
	} else {
		fmt.Println("Video is not upload able.....")
	}

	OS();
	rateLimt(sender{
		rateLimt: 31,
		user: user {
			username: "thearkan",
			id: 1020,
		},
	})

	testAuth(authenticationInfo{
		username: "ali6050",
		password: "password10",
	})
}
