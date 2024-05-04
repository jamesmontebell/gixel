#!/bin/bash

# Path to the output file
OUTPUT_FILE="./git_commit_output.txt"

# Check if the command is 'git commit' with the '-m' flag
if [ "$1" = "commit" ] && [ "$2" = "-m" ]; then
    # Extract the commit message (all arguments after '-m')
    shift 2
    commit_message="$@"
    
    # Execute the git commit command with the provided message
    git commit -m "$commit_message" > "$OUTPUT_FILE" 2>&1
    
    # Display a message indicating where the output was saved
    echo "Git commit output saved to: $OUTPUT_FILE"
else
    # For other git commands, execute them directly
    git "$@"
fi

