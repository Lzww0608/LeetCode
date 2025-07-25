

import "math/bits"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pseudoPalindromicPaths (root *TreeNode) (ans int) {
    var dfs func(*TreeNode, int) 
    dfs = func(root *TreeNode, mask int) {
        if root == nil {
            return 
        }

        mask ^= 1 << root.Val
        dfs(root.Right, mask)
        dfs(root.Left, mask)
        if root.Left == nil && root.Right == nil {
            cnt := bits.OnesCount(uint(mask))
            if cnt <= 1 {
                ans++
            }
        }
    }
    dfs(root, 0)

    return
}