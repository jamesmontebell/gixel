#!/bin/bash

# Check if the commit message is provided as an argument
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <commit_message>"
    exit 1
fi

commit_message="$1"

# Execute git commit with the specified message and capture the output into a variable
commit_output=$(git commit -m "$commit_message" 2>&1)

echo $commit_output

# Pass the commit output directly to the Go program as an argument
# /path/to/your/go/program "$commit_output"

exit
