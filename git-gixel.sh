#!/bin/bash

# Check if there are any command line arguments
if [ $# -eq 0 ]; then
    echo "Please provide a commit message"
    exit 1
fi

# Read the input provided
commit_message="$1"

# Set the output of commit to a variable
commit_output=$(git commit -m "$commit_message" 2>&1)

# Check the exit status of the git commit command
if [ $? -eq 0 ]; then
    echo "Commit successful"
else
    echo "Commit failed"
fi

echo "$commit_output"
exit 0