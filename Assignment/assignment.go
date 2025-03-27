package Assignment

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/* Objective:
Build a command-line Contact Book where users can add, search, and delete contacts.

Requirements:
Add a Contact: Prompt the user for a name and a phone number. Store contacts in a map (map[string]string where name -> phoneNumber).
Search for a Contact :Allow users to enter a name and find the corresponding phone number. If the name doesn’t exist, display a message saying "Contact not found".
Delete a Contact: Allow users to remove a contact by name.
List All Contacts: Print all saved contacts in a readable format.
Exit the Program: Provide an option to exit the program.
*/

var contacts map[string]string

func Task() {
	contacts = make(map[string]string)
	for {
		fmt.Printf("1.Search\t2.Add\t3.Delete\t4.List\t5.Exit\n")
		fmt.Printf("Enter choice: ")

		option := takeOptionInput()

		switch option {
		case 1:
			searchContact()
		case 2:
			addNewContact()
		case 3:
			deleteContact()
		case 4:
			listAllContacts()
		case 5:
			fmt.Printf("Exiting Program.")
			os.Exit(0)
		default:
			fmt.Printf("Please enter valid option!\n")
		}
	}
}

func searchContact() {
	if len(contacts) == 0 {
		fmt.Printf("No contacts to search\n")
		return
	}
	fmt.Printf("Enter contact name: ")
	name := takeNameInput()
	if _, exists := contacts[name]; !exists {
		fmt.Printf("No contacts found with name='%s'\n", name)
		return
	}
	fmt.Printf("Found contact {name: '%s', number:'%s'}\n", name, contacts[name])
}

func addNewContact() {
	fmt.Printf("Enter contact name: ")
	name := takeNameInput()
	fmt.Printf("Enter contact number: ")
	number := takeNumberInput()
	contacts[name] = number
	fmt.Printf("Added new contact {name: '%s', number: '%s'}\n", name, number)
}

func deleteContact() {
	if len(contacts) == 0 {
		fmt.Printf("No contacts to delete\n")
		return
	}
	fmt.Printf("Enter contact name to delete: ")
	name := takeNameInput()
	if _, exists := contacts[name]; !exists {
		fmt.Printf("Contact '%s' not found\n", name)
		return
	}
	delete(contacts, name)
	fmt.Printf("Contact '%s' successfully deleted\n", name)
}

func listAllContacts() {
	i := 0
	if len(contacts) == 0 {
		fmt.Printf("Contacts list is empty\n")
		return
	}
	for name, number := range contacts {
		i += 1
		fmt.Printf("[%d] name=%s\tnumber=%s\n", i, name, number)
	}
}

func takeNumberInput() string {
	for {
		input := getInputFromUser()
		if isValidContactNumber(input) {
			return input
		}
		fmt.Println("Please enter a valid contact number (exactly 10 digits)")
	}
}

func isValidContactNumber(number string) bool {
	if len(number) != 10 {
		return false
	}

	for _, char := range number {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func takeNameInput() string {
	for {
		input := getInputFromUser()
		if isValidName(input) {
			return input
		}
		fmt.Println("Please enter a valid name (letters only)")
	}
}

func isValidName(name string) bool {
	if len(name) == 0 {
		return false
	}

	for _, char := range name {
		if !unicode.IsLetter(char) && char != ' ' && char != '-' {
			return false
		}
	}
	return true
}

func takeOptionInput() int {
	input := getInputFromUser()

	option, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}
	return option
}

func getInputFromUser() string {
	reader := bufio.NewReader(os.Stdin)
	inputLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input")
		reader.Reset(os.Stdin)
		return ""
	}
	return strings.TrimSpace(inputLine)
}
