def validAnagram(s, t):

    """
    Given two strings s and t, return true if t is an anagram of s, 
    and false otherwise.

    An Anagram is a word or phrase formed by rearranging the letters of
    a different word or phrase, typically using all the original letters
    exactly once.
    """
    countS = {}
    countT = {}

    for i in range(len(s)):
        countS[s[i]] = 1 + countS.get(s[i], 0)
        countT[t[i]] = 1 + countT.get(t[i], 0)
    for j in countS:
        if countS[j] != countT.get(j, 0):
            return False
    return True

s = "anagram"
t = "naaram"
result = validAnagram(s, t)
print(result)