def containsDuplicates(nums):
    """
    Given an integer array nums, return true if any value appears at 
    least twice in the array, and return false if every element is distinct.
    """
    hashmap = []
    for i in nums:
        if i in hashmap:
            return True
        else:
            hashmap.append(i)
    return False

nums = [1,4,6,7]
result = containsDuplicates(nums)
print(result)



