class Codec:
    def encode(strs):
        newStr = ""
        for i in strs:
            newStr += str(len(i)) + "$" + i
        return newStr
    
    def decode(str):
        newList, i = [], 0

        while i<len(str):
            j = i
            if j != "$":
                j += 1
            length = int(str[i:j])
            newList.append(str[j + 1 : j + 1 + length])
            i = j + 1 + length
        return newList