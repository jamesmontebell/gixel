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

type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

// Init implements tea.Model.
func (m model) Init() tea.Cmd {
	panic("unimplemented")
}

func initialModel() model {
	return model{
		// Our to-do list is a grocery list
		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are selected. We're using
		// the map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
