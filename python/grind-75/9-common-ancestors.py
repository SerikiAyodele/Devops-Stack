def commonAncestors(root, p, q):
    curr = root

    while curr:
        if p.val > curr.val and q.val > curr.val:
            curr = curr.right
        elif p.val < curr.val and q.val < curr.val:
            curr = curr.left
        else:
            return curr

root = [6,2,8,0,4,7,9,3,5]
p = 2
q = 8
result = commonAncestors(root, p, q)
print(result)