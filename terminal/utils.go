package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	baseExp          = 64
	filesExpMult     = 2
	insertsExpMult   = 4
	deletionsExpMult = 3
	scaler           = 7
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
func calculateExp(parsedOutputs ParsedOutputs) float32 {
	combinedOutputs := parsedOutputs.Changes*filesExpMult + parsedOutputs.Deletions*deletionsExpMult + parsedOutputs.Inserts*insertsExpMult
	num := baseExp * combinedOutputs
	res := float32(num / scaler)
	return res
}

// Define the model
type model struct {
	count int
}

// Init function initializes the model
func (m model) Init() tea.Cmd {
	return nil
}

// Update function to handle messages and update the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up":
			m.count++
		case "down":
			m.count--
		}
	}
	return m, nil
}

// View function to render the UI
func (m model) View() string {
	return fmt.Sprintf("Count: %d\n\n↑: up\n↓: down\nq: quit\n", m.count)
}
