package main

import "fmt"

func main() {
	// printBasicDataTypes()
	// arrayOfData()
	iterateArray()
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

func iterateArray() {
	maxLen := 5

	var charArray1 []rune = make([]rune, 0, maxLen)
	for i := 0; i < maxLen; i++ {
		charArray1 = append(charArray1, rune(i+97))
		fmt.Printf("rune[%d] = %c\n", i, charArray1[i])
	}

	charArray2 := []rune{'9', 'a', '*'}

	fmt.Println()
	for i, char := range charArray2 {
		fmt.Printf("rune[%d] = %c, %d\n", i, char, char)
	}

}
