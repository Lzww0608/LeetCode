/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func checkEqualTree(root *TreeNode) (ans bool) {
    sum := dfs0(root)

    var dfs func(*TreeNode) (cur int) 
    dfs = func(root *TreeNode) (cur int) {
        if ans || root == nil {
            return 
        }
        
        cur += root.Val + dfs(root.Left) + dfs(root.Right)
        if cur * 2 == sum {
            ans = true 
        }
        return 
    }
    dfs(root.Left)
    dfs(root.Right)

    return
}

func dfs0(root *TreeNode) (ans int) {
    if root == nil {
        return
    }

    ans += root.Val + dfs0(root.Left) + dfs0(root.Right)
    return
}