class Solution(object):
    def lowestCommonAncestor(self, root, p, q):
        if root == p or root == q:
            return root.val

        while root:
            if root.left < p.val and root.right > q.val:
                 root.val
            elif root.left > p.val:
                 root.right.val
            elif root.right < q.val:
                 root.left.val
            else:
                break