package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No arguments")
		os.Exit(1)
	}

	type Outputs struct {
		Changes   int
		Inserts   int
		Deletions int
	}

	var output Outputs
	var args string

	for i := 1; i < len(os.Args); i++ {
		args += os.Args[i]
	}

	fmt.Println(args)

	changes, err := strconv.Atoi(args[FindFilesChanged(args)-1 : FindFilesChanged(args)])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	insertions, err := strconv.Atoi(args[FindInsertions(args)-1 : FindInsertions(args)])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	deletions, err := strconv.Atoi(args[FindDeletions(args)-1 : FindDeletions(args)])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	output.Changes = changes
	output.Inserts = insertions
	output.Deletions = deletions

	fmt.Print(output)
}

func FindFilesChanged(s string) int {
	fileChangedString := "filechanged"
	filesChangedString := "fileschanged"

	if strings.Contains(s, fileChangedString) {
		return strings.Index(s, fileChangedString)
	} else {
		return strings.Index(s, filesChangedString)
	}
}

func FindInsertions(s string) int {
	insertionString := "insertion"
	insertionsString := "insertions"

	if strings.Contains(s, insertionString) {
		return strings.Index(s, insertionString)
	} else {
		return strings.Index(s, insertionsString)
	}
}

func FindDeletions(s string) int {
	deletionString := "deletion"
	deletionsString := "deletions"

	if strings.Contains(s, deletionString) {
		return strings.Index(s, deletionString)
	} else {
		return strings.Index(s, deletionsString)
	}
}
