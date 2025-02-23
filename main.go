package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	// printBasicDataTypes()
	// arrayOfData()
	// iterateArray()
	userInput()
	// menu()
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

func userInput() {
	var number = int(5)
	fmt.Printf(`Enter integer for address:%%d: %d
			  %%v: %v
			  %%p: %p
 			  %%x: %x
			 `, &number, &number, &number, &number)
	fmt.Scanf("%d", &number)
	fmt.Printf("User input is: %d", number)

	var reader = bufio.NewReader(os.Stdin)
	fmt.Printf("\nType of reader: %T", reader)
	reader.ReadString('\n')

	char := 'A'
	fmt.Printf("\nEnter char input: ")
	b, _ := reader.ReadByte()
	char = rune(b)
	fmt.Printf("Entered char is %c, %v, %v", char, char, &char)

	var str string = "myString value"
	fmt.Printf("\nEnter input for string: ")
	reader.ReadString('\n')
	str, _ = reader.ReadString('\n')
	fmt.Printf("Entered string: %s", str)

	array := []int{1, 2, 3, 4, 5}
	fmt.Println("Enter elements for array")
	for i := range array {
		fmt.Scanf("%d", &array[i])
	}

	fmt.Println("Enter elements:")
	for _, val := range array {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}

func clearScreen() {
	// fmt.Print("\033[H\033[2J")
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func menu() {
	var check bool = true
	var optionEntered rune
	reader := bufio.NewReader(os.Stdin)
	for check {
		fmt.Printf("1.Enter 'y' to clear screen\n" +
			"2.Enter 'h' to print 'Hello World'\n" +
			"3.Enter 'q' to quit\n")
		b, _ := reader.ReadByte()
		reader.ReadString('\n')
		optionEntered = rune(b)
		switch optionEntered {
		case 'y':
			fmt.Printf("\033[H\033[2J")
		case 'h':
			fmt.Printf("Hello World\n")
		case 'q':
			check = false
		default:
			fmt.Println("Invalid option")
		}
	}
}
