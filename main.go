package main

import (
	"bufio"
	"fmt"
	person "golearn/Person"
	"log"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	clearScreen()
	// printBasicDataTypes()
	// arrayOfData()
	// iterateArray()
	// userInput()
	// menu()
	// structPerson()
	// someStringFunctions()
	binaryRepresentation()
}

func binaryRepresentation() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal("Error reading input:", err)
		}
		log.Fatal("No input provided")
		return
	}

	str := scanner.Text()
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("Invalid number\nerror:", err)
	}

	fmt.Printf("MinInt = %d, MaxInt = %d\n", math.MinInt, math.MaxInt)
	fmt.Printf("MinInt8 = %d, MaxInt8 = %d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("Binary: %b\n", num)
	fmt.Printf("Binary (8-bit): %08b\n", num)
}

func someStringFunctions() {
	str := "a random stRing ǳ 😶"
	fmt.Printf("Type: %T, size: %d, runeCount: %d", str, len(str), utf8.RuneCountInString(str))
	fmt.Printf("\nString: %s", str)

	fmt.Printf("\nuppercase: %s", strings.ToUpper(str))
	fmt.Printf("\nlowercase: %s", strings.ToLower(str))
	fmt.Printf("\ntotitle: %s", strings.ToTitle(str))

	newString := strings.Map(func(r rune) rune {
		fmt.Printf("\nRune '%c' uses %d bytes", r, utf8.RuneLen(r))
		return 'ʬ'
	}, str)
	fmt.Printf("\nModified String: %s", newString)

	for _, r := range newString {
		fmt.Printf("\n%c, %c", r, newString[0])
		fmt.Print("\n", []byte(string(r)), []byte(string(newString[0])))
		fmt.Printf("\n []byte(string[rune]) = %v, []byte(string[byte]) = %v", []byte(string(r)), []byte(string(newString[0])))
		bytes := []byte(string(r))
		fmt.Printf("\nr = %c, bytes = %d\n", r, len(bytes))
		for _, b := range bytes {
			fmt.Printf("%08b ", b)
		}
		fmt.Println()

		for _, b := range []byte(string(newString[0])) {
			fmt.Printf("%08b ", b)
		}
		fmt.Println()
		break
	}
}

func structPerson() {
	person.PersonExample()
	// go build -gcflags="-m" main.go
	fmt.Printf("%T\n", (&person.Person{}).Languages)
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
