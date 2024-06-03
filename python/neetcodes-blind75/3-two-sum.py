# pseudo code
# hashmap for getting the Index
# iterate through the hashmap
# see what two values in the hashmap will add up to the target

def twoSum(nums, target):
    hashmap = {}

    for i, n in enumerate(nums):
        diff = target - n
        if diff in hashmap:
            return [hashmap[diff], i]
        hashmap[n] = i
    return 

nums = [2,7,11,15]
target = 9
result = twoSum(nums, target)
print(result)