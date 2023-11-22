# **Lab 15.6: Arithmetic and Functions**
# Write a script that will act as a simple calculator for add, subtract, multiply and divide.

# Each operation should have its own function.
# Any of the three methods for bash arithmetic, ($((..)), let, or expr) may be used.
# The user should give 3 arguments when executing the script:
# - The first should be one of the letters a, s, m, or d to specify which math operation.
# - The second and third arguments should be the numbers that are being operated on.
#  The script should detect for bad or missing input values and display the results when done.


#!/bin/bash
add (){
    answer1 = $(($1 + $2))
}

subtract (){
    answer1 = $(($1 - $2))
}

multiply (){
    answer1 = $(($1 * $2))
}

divide(){
    answer1 = $(($1 / $2))
}

op1=$1 ; arg1=$2 ; arg2=$3

[[ $# -lt 3 ]] && \
    echo "Usage: Provide an operation (a,s,m,d) and two numbers" && exit 1
    
[[ $op != a ]] && [[ $op != s ]] && [[ $op != m ]] && [[ $op != d ]] && \
    echo "operator must be a, s, m, or d not $op as supplied"

if [[ $op == a ]]; then add $arg1 $arg2
elif [[ $op == s ]]; then sub $arg1 $arg2
elif [[ $op == m ]]; then mul $arg1 $arg2
elif [[ $op == d ]]; then div $arg1 $arg2
else
echo "$op is not a,s,m, or d, aborting" ; exit 2
fi

#show answers
echo $arg1,$op1,$arg2:
echo "$((..))" Answer is $answer1
