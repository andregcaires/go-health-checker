package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const numberOfTries = 3
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

	websites := createWebsitesSliceFromFile()

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

	response, err := http.Get(website)

	if err != nil {
		fmt.Println("Error at HTTP request:", err)
	} else {

		if response.StatusCode == 200 {
			fmt.Println(website, "is OK")
		} else {
			fmt.Println(website, "is having issues:", response.StatusCode)
		}
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

func createWebsitesSliceFromFile() []string {

	var websites []string

	file, err := os.Open("websites.txt")

	if err != nil {
		fmt.Println("Error at opening file:", err)

	} else {

		reader := bufio.NewReader(file)

		for {

			line, err := reader.ReadString('\n')

			line = strings.TrimSpace(line)

			websites = append(websites, line)

			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println("Error at reading file:", err)
			}
		}
	}

	file.Close()

	return websites
}
