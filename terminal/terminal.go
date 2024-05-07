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

	changes := findFilesChanged(args)
	insertions, err := strconv.Atoi(args[findInsertions(args)-1 : findInsertions(args)])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	deletions, err := strconv.Atoi(args[findDeletions(args)-1 : findDeletions(args)])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	output.Changes = changes
	output.Inserts = insertions
	output.Deletions = deletions

	fmt.Print(output)
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func convertToNumeric(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	return num
}

func findFilesChanged(s string) int {
	fileChangedString := "filechanged"
	filesChangedString := "fileschanged"

	var num string

	if strings.Contains(s, fileChangedString) {
		num += string(s[strings.Index(s, fileChangedString)-1])
		return convertToNumeric(num)
	} else {
		if isNumeric(string(s[strings.Index(s, filesChangedString)-3])) {
			num += string(s[strings.Index(s, filesChangedString)-3])
		}
		num += string(s[strings.Index(s, filesChangedString)-2])
		num += string(s[strings.Index(s, filesChangedString)-1])
		return convertToNumeric(num)
	}
}

func findInsertions(s string) int {
	insertionString := "insertion"
	insertionsString := "insertions"

	if strings.Contains(s, insertionString) {
		return strings.Index(s, insertionString)
	} else {
		return strings.Index(s, insertionsString)
	}
}

func findDeletions(s string) int {
	deletionString := "deletion"
	deletionsString := "deletions"

	if strings.Contains(s, deletionString) {
		return strings.Index(s, deletionString)
	} else {
		return strings.Index(s, deletionsString)
	}
}
