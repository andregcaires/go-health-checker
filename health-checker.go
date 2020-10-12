package main

import (
	"fmt"
	"os"
)

func main() {

	showOptionsMenu()

	input := getUsetInput()

	startSelectedOperation(input)
}

func showOptionsMenu() {

	fmt.Println("1 - Start monitoring")
	fmt.Println("1 - Show logs")
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
}

func showLogs() {

	fmt.Println("Showing logs...")
}
