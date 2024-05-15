package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	filesExpMult     = 50
	insertsExpMult   = 50
	deletionsExpMult = 50
)

// Check if a string is a number
func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// Convert a string number to an integer
func convertToNumeric(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	return num
}

// Find the number of files changed in git commit output
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
		if isNumeric(string(s[strings.Index(s, filesChangedString)-2])) {
			num += string(s[strings.Index(s, filesChangedString)-2])
		}
		num += string(s[strings.Index(s, filesChangedString)-1])
		return convertToNumeric(num), nil
	}
	return 0, errors.New("files changed error")
}

// Find the number of insertions added in git commit output
func findInsertions(s string) (int, error) {
	insertionString := "insertion("
	insertionsString := "insertions"

	var num string

	if strings.Contains(s, insertionString) {
		num += string(s[strings.Index(s, insertionString)-1])
		return convertToNumeric(num), nil
	} else if strings.Contains(s, insertionsString) {
		if isNumeric(string(s[strings.Index(s, insertionsString)-5])) {
			num += string(s[strings.Index(s, insertionsString)-5])
		}
		if isNumeric(string(s[strings.Index(s, insertionsString)-4])) {
			num += string(s[strings.Index(s, insertionsString)-4])
		}
		if isNumeric(string(s[strings.Index(s, insertionsString)-3])) {
			num += string(s[strings.Index(s, insertionsString)-3])
		}
		if isNumeric(string(s[strings.Index(s, insertionsString)-2])) {
			num += string(s[strings.Index(s, insertionsString)-2])
		}
		num += string(s[strings.Index(s, insertionsString)-1])
		return convertToNumeric(num), nil
	}
	return 0, errors.New("no insertion")
}

// Find the number of deletions in git commit output
func findDeletions(s string) (int, error) {
	deletionString := "deletion("
	deletionsString := "deletions"
	var num string

	if strings.Contains(s, deletionString) {
		num += string(s[strings.Index(s, deletionString)-1])
		return convertToNumeric(num), nil
	} else if strings.Contains(s, deletionsString) {
		if isNumeric(string(s[strings.Index(s, deletionsString)-4])) {
			num += string(s[strings.Index(s, deletionsString)-4])
		}
		if isNumeric(string(s[strings.Index(s, deletionsString)-3])) {
			num += string(s[strings.Index(s, deletionsString)-3])
		}
		if isNumeric(string(s[strings.Index(s, deletionsString)-2])) {
			num += string(s[strings.Index(s, deletionsString)-2])
		}
		num += string(s[strings.Index(s, deletionsString)-1])
		return convertToNumeric(num), nil
	}
	return 0, errors.New("no deletions")
}

// Take parsed git commit output numbers and calculate an experience amount
func calculateExp(parsedOutputs ParsedOutputs) int {
	return parsedOutputs.Changes*filesExpMult + parsedOutputs.Inserts*insertsExpMult + parsedOutputs.Deletions*deletionsExpMult
}
