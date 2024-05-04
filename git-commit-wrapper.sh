#!/bin/bash

# Path to the output file
OUTPUT_FILE="./git_commit_output.txt"

# Execute the git commit command with the provided message
git commit "$@" > "$OUTPUT_FILE" 2>&1

# Optionally, you can add additional actions here after the commit
echo "Git commit output saved to: $OUTPUT_FILE"

