#!/bin/bash

echo "Enter commit message: "
read message

git commit -m "$message"

current_directory=$(pwd)

exit