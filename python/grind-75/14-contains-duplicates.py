def containsDuplicate(nums):
    hashset = []

    for i in range(len(nums)):
        if nums[i] in hashset:
            return True
        hashset.append(i)
    return

nums = [1,1,1,3,3,4,3,2,4,2]
result = containsDuplicate(nums)
print(result)