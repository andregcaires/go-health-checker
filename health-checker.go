package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const numberOfTries = 5
const delay = 3

func main() {

	initializeProgram()
}

func initializeProgram() {

	for {
		showOptionsMenu()

		input := getUsetInput()

		startSelectedOperation(input)
	}
}

func showOptionsMenu() {

	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("9 - Exit program")
}

func getUsetInput() int {

	var input int
	fmt.Scan(&input)
	return input
}

func startSelectedOperation(input int) {

	switch input {
	case 1:
		startMonitoring()
	case 2:
		showLogs()
	case 9:
		fmt.Println("Exiting program...")
		os.Exit(0)
	default:
		fmt.Println("Command not found")
		os.Exit(-1)
	}
}

func startMonitoring() {

	fmt.Println("Starting monitoring...")

	websites := createWebsitesSlice()

	for i := 0; i < numberOfTries; i++ {
		for _, website := range websites {

			callWebsite(website)
		}
		if i != (numberOfTries - 1) {
			time.Sleep(delay * time.Second)
		}
	}

}

func callWebsite(website string) {

	response, _ := http.Get(website)

	if response.StatusCode == 200 {
		fmt.Println(website, "is OK")
	} else {
		fmt.Println(website, "is having issues:", response.StatusCode)
	}
}

func showLogs() {

	fmt.Println("Showing logs...")
}

func createWebsitesSlice() []string {

	return []string{
		"https://www.google.com.br",
		"https://www.alura.com.br",
		"https://www.caelum.com.br",
	}
}
