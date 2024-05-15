#!/bin/bash

# Check if there are any command line arguments
if [ $# -eq 0 ]; then
    echo "Please provide a commit message."
    exit 1
fi

# Check if commit message ends in a number
last_char="${1: -1}"
if [[ "$last_char" =~ [0-9] ]]; then
    echo "Please don't end commit message with a number."
    exit 1
fi

# Read the input provided
commit_message="$1"

# Set the output of commit to a variable
commit_output=$(git commit -m "$commit_message" 2>&1)

echo "$commit_output"

# Check the exit status of the git commit command
if [ $? -eq 0 ]; then
    echo "Commit successful"
else
    echo "Commit failed"
fi


# Delete spaces and whitespace from string and last two characters because
# interprets .go at the end of some commits as input to go run
commit_output="${commit_output//[[:space:]]/}"
dropped=$(echo "$commit_output" | sed 's/..$//')

# Run go TUI program with git commit output
cd /Users/jamesmontebell/Github/gixel/terminal
go run terminal.go types.go utils.go $dropped

exit 0