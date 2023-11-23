# Write a script which reads two strings as arguments and then:

# Tests to see if the first string is of zero length, and if the other is of non-zero length, telling the user of both results.
# Determines the length of each string, and reports on which one is longer or if they are of equal length.
# Compares the strings to see if they are the same, and reports on the result.

#!/bin/bash
echo "give two strings"
read string1 string2
echo string1=$string1, string2=$string2

# checking if they are of zero length
echo "Is string 1 zero length? Value of 1 means FALSE"
[[ -z $string1 ]]
echo $? 
echo "Is string 2 nonzero length? Value of 0 means TRUE"
[[ -n $string2 ]]
echo $?

# comparing the lengths
if [[ $string1 -gt $string2 ]] ; then
    echo "$string1 is longer than $string2"
else
    echo "$string2 is longer than $string1"
fi

# checking if they are the same string
if [[ $string1 == $string2 ]] ; then
    echo "The strings are the same"
else
    echo "The strings are not the same"
fi