package main

import "fmt"

func main() {
	// printBasicDataTypes()
	arrayOfData()
}

func printBasicDataTypes() {
	var number int = 5
	fmt.Println("int: ", number)

	var boolean bool = true
	fmt.Println("bool: ", !boolean)

	var character rune = 'a'
	fmt.Printf("char: %c", character)

	var string string = "a string value"
	fmt.Printf("\nstring: %s\n", string)

	fmt.Println("%T, %T, %T, %T", number, boolean, character, string)
	fmt.Println(character)
	fmt.Printf("%T, %T, %T, %T", number, boolean, character, string)
}

func arrayOfData() {
	var charArray1 []rune = []rune{}
	fmt.Println("Uninitialized Rune: ", charArray1, "len: ", len(charArray1), "capacity: ", cap(charArray1))

	var charArray2 []rune = make([]rune, 2, 3)
	fmt.Println("Rune Array of size : ", charArray2, "len: ", len(charArray2), "capacity: ", cap(charArray2))

	var charArray3 []rune = []rune{'a', 98}
	fmt.Println("Initialized Rune: ", charArray3, "len: ", len(charArray3), "capacity: ", cap(charArray3))
	fmt.Printf("rune[%d] = %c, rune[%d] = %c\n", 0, charArray3[0], 1, charArray3[1])
	fmt.Printf("rune[%d] = %d, rune[%d] = %d\n", 0, charArray3[0], 1, charArray3[1])
	fmt.Printf("rune = %s\n", charArray3)
	fmt.Printf("rune = %s", string(charArray3))
}
