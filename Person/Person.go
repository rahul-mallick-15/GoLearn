package person

import "fmt"

type Person struct {
	Languages, ProgrammingLanguages []string
}

func (p *Person) printPerson() {
	fmt.Print("Languages: ")
	for _, lang := range p.Languages {
		fmt.Printf("%s ", lang)
	}
	fmt.Print("\nProgramming Languages: ")
	for _, plang := range p.ProgrammingLanguages {
		fmt.Printf("%s ", plang)
	}
	fmt.Println()
}
