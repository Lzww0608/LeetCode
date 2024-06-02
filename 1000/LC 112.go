/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) (f bool) { 
    var dfs func(*TreeNode, int) 
    dfs = func(root *TreeNode, sum int) {
        if root == nil || f {
            return
        }

        sum += root.Val
        if root.Left == nil && root.Right == nil && sum == targetSum {
            f = true 
            return
        } else {
            dfs(root.Left, sum)
            dfs(root.Right, sum)
        }
        
    }

    if root != nil {
        dfs(root, 0)
    }
    
    return
}
