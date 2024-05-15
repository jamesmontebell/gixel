package main

import (
	"fmt"
	"os"
)

func main() {
	// Program needs a git commmit output
	if len(os.Args) < 2 {
		fmt.Println("No arguments")
		os.Exit(1)
	}

	var output ParsedOutputs
	var args string

	// Save arguements into a variable
	args += os.Args[1]

	// Save the number of changes made into a variable
	changes, err := findFilesChanged(args)
	if err != nil {
		fmt.Println("Error occurred:", err)
	}

	// Save the number of insertions made into a variable
	insertions, err := findInsertions(args)
	if err != nil {
		fmt.Println("Error occurred:", err)
	}

	// Save the number of deletions made into a variable
	deletions, err := findDeletions(args)
	if err != nil {
		fmt.Println("Error occurred:", err)
	}

	// Add these variables into Output object
	output.Changes = changes
	output.Inserts = insertions
	output.Deletions = deletions

	fmt.Print(output)

	res := calculateExp(output)

	fmt.Println(res)
}
