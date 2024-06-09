package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
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
		fmt.Println(err)
	}

	// Save the number of insertions made into a variable
	insertions, err := findInsertions(args)
	if err != nil {
		fmt.Println(err)
	}

	// Save the number of deletions made into a variable
	deletions, err := findDeletions(args)
	if err != nil {
		fmt.Println(err)
	}

	// Add these variables into Output object
	output.Changes = changes
	output.Inserts = insertions
	output.Deletions = deletions

	fmt.Print(output)

	exp := calculateExp(output)

	fmt.Println(exp)

	cmd := exec.Command("git", "config", "user.email")
	var outBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	email := outBuffer.String()
	email = email[:len(email)-1]

	var newExp Experience
	newExp.Exp = int(exp)
	newExp.UserEmail = email

	jsonBytes, err := json.Marshal(newExp)
	if err != nil {
		panic(err)
	}

	jsonString := string(jsonBytes)
	res, err := http.Post("http://localhost:1234/newcommit", "application/json", strings.NewReader(jsonString))
	if err != nil {
		panic(err)
	}

	if res.StatusCode == 201 {
		fmt.Println("Successful Experience Upload!")
	} else {
		fmt.Println("Failed to save to database!")
	}

	p := tea.NewProgram(model{count: 0})
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v\n", err)
		os.Exit(1)
	}
}
