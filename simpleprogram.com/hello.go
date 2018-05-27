package main

import "simpleprogram.com/greeting"

func RenameToFrog(r greeting.Renamable) {
	r.Rename("Frog")
}

func main() {
	//var s = greeting.Salutation{"Mary", "Hi!"}
	//s := []int{1, 10, 500, 25}

	salutations := greeting.Salutations{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "Wassup"},
	}

	// greeting.Salutation{"Frank", "Hi"}
	//slice = append(slice[:1], slice[2:]...)

	//slice = slice[1:]

	salutations[0].Rename("John")
	RenameToFrog(&salutations[0])
	salutations.Greet(greeting.CreatePrintFunction("!!!"), true, 6)

	//greeting.Greet(slice, greeting.CreatePrintFunction("!!!"), true, 6)
}
