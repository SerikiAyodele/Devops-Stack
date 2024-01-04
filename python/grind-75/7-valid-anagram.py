class Solution(object):
    def validAnagram(self, s: str, t: str):
        if len(s) != len(t):
            return False
        
        countS, countT = {}, {}

        for i in range(len(s)):
            countS[s[i]] = 1 + countS.get(s[i], 0)
            countS[t[i]] = 1 + countT.get(t[i], 0)
        for c in countS:
            if countS[c] != countT.get(c,0):
                return False
        return True
    
solution = Solution()  

s , t = "anagram", "nagaram"
result = solution.validAnagram(s, t)
print(result)