# A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters,
# it reads the same forward and backward. Alphanumeric characters include letters and numbers.

# Given a string s, return true if it is a palindrome, or false otherwise.

# Example 1:

# Input: s = "A man, a plan, a canal: Panama"
# Output: true
# Explanation: "amanaplanacanalpanama" is a palindrome.

class Solution(object):
    def isPalindrome(self, s):
        l,r = 0,len(s)-1

        while l<r:
            while l<r and not isalnum(s[l]):
                l += 1
            while r>l and not isalnum(s[r]):
                r -= 1
            if s[l].lower() != s[r].lower():
                return False
            l,r = l + 1, r - 1
        return True


    def isalnum(self, j):
        return (ord("A") <= ord(j) <= ord("Z"),
                ord("a") <= ord(j) <= ord("z"),
                ord("0") <= ord(j) <= ord("9"))