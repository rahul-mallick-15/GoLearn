package person

func PersonExample() {
	dummyPerson := Person{Languages: []string{"English", "Spanish", "Mandarin"}, ProgrammingLanguages: []string{"Cobalt", "Carbon", "Solana"}}
	dummyPerson.printPerson()
}
