package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Outputs struct {
	Changes   int
	Inserts   int
	Deletions int
}

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

func findFilesChanged(s string) (int, error) {
	fileChangedString := "filechanged"
	filesChangedString := "fileschanged"

	var num string

	if strings.Contains(s, fileChangedString) {
		num += string(s[strings.Index(s, fileChangedString)-1])
		return convertToNumeric(num), nil
	} else if strings.Contains(s, filesChangedString) {
		if isNumeric(string(s[strings.Index(s, filesChangedString)-3])) {
			num += string(s[strings.Index(s, filesChangedString)-3])
		}
		num += string(s[strings.Index(s, filesChangedString)-2])
		num += string(s[strings.Index(s, filesChangedString)-1])
		return convertToNumeric(num), nil
	}
	return 0, errors.New("files changed error")
}

func findInsertions(s string) (int, error) {
	insertionString := "insertion"
	insertionsString := "insertions"

	var num string

	if strings.Contains(s, insertionString) {
		num += string(s[strings.Index(s, insertionString)-1])
		return convertToNumeric(num), nil
	} else if strings.Contains(s, insertionsString) {
		if isNumeric(string(s[strings.Index(s, insertionsString)-3])) {
			num += string(s[strings.Index(s, insertionsString)-3])
		}
		num += string(s[strings.Index(s, insertionsString)-2])
		num += string(s[strings.Index(s, insertionsString)-1])
		return convertToNumeric(num), nil
	}
	return 0, errors.New("no insertion")
}

func findDeletions(s string) (int, error) {
	deletionString := "deletion"
	deletionsString := "deletions"
	var num string

	if strings.Contains(s, deletionString) {
		num += string(s[strings.Index(s, deletionString)-1])
		return convertToNumeric(num), nil
	} else if strings.Contains(s, deletionsString) {
		if isNumeric(string(s[strings.Index(s, deletionsString)-3])) {
			num += string(s[strings.Index(s, deletionsString)-3])
		}
		num += string(s[strings.Index(s, deletionsString)-2])
		num += string(s[strings.Index(s, deletionsString)-1])
		return convertToNumeric(num), nil
	}
	return 0, errors.New("no deletions")
}
