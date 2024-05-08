package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No arguments")
		os.Exit(1)
	}

	var output Outputs
	var args string

	for i := 1; i < len(os.Args); i++ {
		args += os.Args[i]
	}

	changes, err := findFilesChanged(args)
	if err != nil {
		fmt.Println("Error occurred:", err)
	}

	insertions, err := findInsertions(args)
	if err != nil {
		fmt.Println("Error occurred:", err)
	}

	deletions, err := findDeletions(args)
	if err != nil {
		fmt.Println("Error occurred:", err)
	}

	output.Changes = changes
	output.Inserts = insertions
	output.Deletions = deletions

	fmt.Print(output)
}
