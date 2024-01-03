class Solution(object):
    def invertTrees(self, root):
        #check if a head node exists
        if not root:
            return None
        
        tmp = root.left
        root.left = root.right
        root.right = tmp

        self.invertTrees(root.left)
        self.invertTrees(root.right)
        return root
    
    root = [2,1,3]
    inversion = invertTrees(root)
    print(inversion)
