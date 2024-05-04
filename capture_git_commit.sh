#!/bin/bash

# Check if the commit message is provided as an argument
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <commit_message>"
    exit 1
fi

commit_message="$1"

# Use echo to pipe the commit message directly into git commit
commit_output=$(echo "$commit_message" | git commit -F - 2>&1)

# Pass the commit output directly to the Go program as an argument
# /path/to/your/go/program "$commit_output"

echo $commit_output
