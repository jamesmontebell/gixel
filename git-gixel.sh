#!/bin/bash

# Check if there are any command line arguments
if [ $# -eq 0 ]; then
    echo "Please provide a commit message"
    exit 1
fi

# Read the input provided
commit_message="$1"

git commit -m "$commit_message"
exit 