#!/bin/bash

# Check if exactly two arguments are provided
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 -m <message>"
    exit 1
fi

# Parse the arguments
while getopts ":m:" opt; do
    case ${opt} in
        m )
            # Verify if the first argument is '-m'
            if [ "$OPTARG" != "-m" ]; then
                echo "Invalid option. Please use '-m' followed by your message."
                exit 1
            fi
            ;;
        \? )
            echo "Invalid option: $OPTARG"
            exit 1
            ;;
    esac
done
shift $((OPTIND -1))

# Now, $1 will be the message
message="$1"

# Print the message
echo "$message"

