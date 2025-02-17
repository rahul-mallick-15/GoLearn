package main

import "fmt"

func main() {
	printBasicDataTypes()
}

func printBasicDataTypes() {
	var number int = 5
	fmt.Println("int: ", number)

	var boolean bool = true
	fmt.Println("bool: ", !boolean)

	var character rune = 'a'
	fmt.Printf("char: %c", character)

	var string string = "a string value"
	fmt.Printf("\nstring: %s", string)
}
