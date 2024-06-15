from collections import defaultdict


def groupAnagram(strs):
    hashmap = defaultdict(list)

    for i in strs:
        count = ord[0] * 26
        for c in i:
            count[ord[c] - ord["a"]] += 1
        
        hashmap(count).append(i)

    return hashmap.values()


    # store = {}

    # for st in strs:
        
    #     temp = sorted(st) 
    #     print(temp)
    #     temp_ = "".join(temp)
        
        
    #     if temp_ in store:
    #         store[temp_].append(st)
    #     else:
    #         store[temp_] = [st]  
    
strs = ["eat","tea","tan","ate","nat","bat"]
res = groupAnagram(strs)
print(strs)