class Solution(object):
    def isBalanced(self, root):
        curr = root

        while curr:
            if curr.left and curr.right:
                curr = curr.val
                return True
            else:
                return False
        