package main

import (
	"bufio"
	pass "code.google.com/p/gopass"
	"fmt"
	"github.com/robfig/config"
	"github.com/wsxiaoys/terminal"
	"log"
	"os"
	"strings"
)

func PrintError() {
	home := fmt.Sprintf("ERROR - Unable to locate api configuration ( %s/.tropo-api.cfg ).\n", UserHomeDir())
	terminal.Stdout.Color("r").Print(home).Nl().Reset()
	terminal.Stdout.Color("b").Print("---------------------------------\n").Nl().Reset()
	fmt.Println("   ; Default api-config\n   ; https://github.com/robfig/config\n   [DEFAULT]\n   host: api.aws.tropo.com\n   route: /rest/v1\n   protocol: https://\n   base-url: %(protocol)s%(host)s%(route)s\n   \n   [hosted]\n   url: %(base-url)s\n   username: <username>\n   password: <password>\n")
	terminal.Stdout.Color("b").Print("---------------------------------\n").Nl().Reset()
	terminal.Stdout.Color("b").Print("Creating new config now\n").Nl().Reset()
}

func CreateConfig() {

	protocol := askForProtocol()
	host := askUserForData("Enter API host: ", "api.aws.tropo.com")
	route := askUserForData("Enter route: ", "/rest/v1")

	username := askUserForData("Username: ", "")
	password, _ := pass.GetPass("Password: ")

	c := config.NewDefault()

	c.AddSection("DEFAULT")
	c.AddOption("DEFAULT", "host", host)
	c.AddOption("DEFAULT", "route", route)
	c.AddOption("DEFAULT", "protocol", protocol)
	c.AddOption("DEFAULT", "base-url", "%(protocol)s%(host)s%(route)s")

	c.AddSection("hosted")
	c.AddOption("hosted", "url", "%(base-url)s")
	c.AddOption("hosted", "username", username)
	c.AddOption("hosted", "password", password)

	home := fmt.Sprintf("%s/.tropo-api.cfg", UserHomeDir())
	c.WriteFile(home, 0600, "Auto generated config")

}

func askUserForData(msg, default_value string) string {
	var str string
	reader := bufio.NewReader(os.Stdin)

	if default_value == "" {
		fmt.Printf("%s ", msg)
	} else {
		fmt.Printf("%s [ %s ]: ", msg, default_value)
	}

	text, _ := reader.ReadString('\n')

	// If text is we use a default it its provided
	if text != "\n" {
		str, _ = askForConfirmation(strings.Replace(text, "\n", "", -1))
	} else {
		if default_value == "" {
			askUserForData(msg, default_value)
		} else {
			str = default_value
		}
	}
	return str
}

func askForProtocol() string {
	fmt.Printf("https ? (y/n) ")

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}

	if containsString(okayResponses, strings.Replace(text, "\n", "", -1)) {
		return "https://"
	} else if containsString(nokayResponses, strings.Replace(text, "\n", "", -1)) {
		return "http://"
	} else {
		fmt.Println("Please type yes or no and then press enter:")
		return askForProtocol()
	}
}

func askForConfirmation(text string) (string, bool) {
	fmt.Printf("Is %s correct ? (y/n) - ", text)
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response) {
		return text, true
	} else if containsString(nokayResponses, response) {
		return "", false
	} else {
		fmt.Println("Please type yes or no and then press enter:")
		return askForConfirmation("")
	}
}

// You might want to put the following two functions in a separate utility package.

// posString returns the first index of element in slice.
// If slice does not contain element, returns -1.
func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

// containsString returns true iff slice contains element
func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}
