#!/bin/bash

# Extract the commit message from the arguments passed to the script
commit_message="${@: -1}"  # Get the last argument, which is assumed to be the commit message

# Call the actual git commit with the provided message
output=$(git commit -m "$commit_message")

echo $output
