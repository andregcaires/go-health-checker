package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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
		newLog(website, false)
	} else {

		if response.StatusCode == 200 {
			fmt.Println(website, "is OK")
			newLog(website, true)
		} else {
			fmt.Println(website, "is having issues:", response.StatusCode)
			newLog(website, false)
		}
	}
}

func showLogs() {

	fmt.Println("Showing logs...")

	// ReadFile doesn't need to close file
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("An error has ocurred while reading log file:", err)
	} else {

		fmt.Println(string(file))
	}
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

func newLog(website string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	if err != nil {
		fmt.Println("Error at opening log file:", err)
	} else {

		file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + website + " - online: " + strconv.FormatBool(status) + "\n")
	}

	file.Close()
}
