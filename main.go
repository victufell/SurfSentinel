package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

const AVAILABLE_STATUS = 200

func main() {
	for {
		handleMenu()
		userOptionSelected := handleGetUserOption()

		switch userOptionSelected {
		case 1:
			webSiteUrl := handleGetUserSite()
			handleTestWebSite(webSiteUrl)

		case 2:
			handleListLogs()

		case 0:
			fmt.Println("Surf Sentinel finished")
			os.Exit(0)

		default:
			fmt.Println("The Select Option is invalid")
			os.Exit(-1)
		}
	}

}

func handleMenu() {
	fmt.Println("*****************************")
	fmt.Println("******* Surf Sentinel *******")
	fmt.Println("*****************************")

	fmt.Println("1) Start monitoring")
	fmt.Println("2) View logs")
	fmt.Println("0) Finish")
}

func handleGetUserOption() int {
	var userOption int
	fmt.Scan(&userOption)

	return userOption
}

func handleGetUserSite() string {
	var webSiteUrl string
	fmt.Println("Enter the URL of the website you want to monitor")
	fmt.Scan(&webSiteUrl)

	return webSiteUrl
}

func handleTestWebSite(webSiteUrl string) {
	response, err := http.Get(webSiteUrl)

	if err != nil {
		fmt.Println("There was an error trying to access the website: ", err)
	}

	fmt.Println("Starting Monitoring the Web Site:", webSiteUrl)
	fmt.Println("...")
	isAvailable := response.StatusCode == AVAILABLE_STATUS

	if isAvailable {
		fmt.Println("The Web Site is Online - Available: Yes")
	} else {
		fmt.Printf("The Web Site is Offline - Available: Not")
	}

	handleRegisterLog(webSiteUrl, isAvailable)

	fmt.Printf("...")
}

func handleRegisterLog(webSiteUrl string, isAvailable bool) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("We had trouble recording the log", err)
	}

	logMessage := time.Now().Format("02/01/2006 15:04:05") + " WebSite: " + webSiteUrl + " • available: " + strconv.FormatBool(isAvailable) + "\n"
	file.WriteString(logMessage)

	file.Close()
}

func handleListLogs() {
	logs, err := os.ReadFile("logs.txt")

	if err != nil {
		fmt.Println("we had problems reading the file")
	}

	fmt.Println(string(logs))
}
