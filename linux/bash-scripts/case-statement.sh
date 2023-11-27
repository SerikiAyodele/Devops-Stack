# Write a script that takes as an argument a month in numerical form (i.e. between 1 and 12),
#  and translates this to the month name and displays the result on standard out (the terminal).

# If no argument is given, or a bad number is given, the script should report the error and exit.

#!/bin/bash

echo "choose a number between 1 and 12"
read response

# set month equal to argument passed for use in the script
case $response in 
    "1") echo "Jan";;
    "2") echo "Feb";;
    "3") echo "March";;
    "4") echo "April";;
    "5") echo "May";;
    "6") echo "June";;
    "7") echo "July";;
    "8") echo "Aug";;
    "9") echo "Sept";;
    "10") echo "Oct";;
    "11") echo "Nov";;
    "12") echo "Dec";;
     *)   echo "you have to give a number between 1 and 12";
          exit 2;;
esac
exit 0 