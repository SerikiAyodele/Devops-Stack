# Lab 16.3: Using Random Numbers

# Write a script which:
# - Takes a word as an argument.
# - Appends a random number to it.
# - Displays the answer.

#!/bin/bash

echo "type in a word"
read response

echo "$response-$RANDOM"
exit 0