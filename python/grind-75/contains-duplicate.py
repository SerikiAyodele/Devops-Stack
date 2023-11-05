"""
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
Example 1:

Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true """

#USING SET
def duplicates(arr):
    setValue = set()
    
    for i in arr:
        if i in setValue:
            return True
        else:
            setValue.add(i)
    return False

arr = [1,2,3,4,5,6,7,8,9,0]
result = duplicates(arr)
print(result)

# BRUTE FORCE SOLUTION
    for i in arr:
        for j in arr:
            if i == j:
                return True
            else:
                return False
    return

arr = [1,1,1,3,3,4,3,2,4,2]
result = duplicates(arr)
print(result)
    
