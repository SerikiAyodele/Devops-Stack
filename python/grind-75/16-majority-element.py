def majorityElement(nums):
    count = {}
    res, maxCount = 0, 0

    for n in nums:
        count[n] = 1 + count.get(n, 0)
        res = n if count[n]