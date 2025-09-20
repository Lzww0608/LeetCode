/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func navigation(root *TreeNode) (ans int) {
    add := 1

    var dfs func(root *TreeNode) int 
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }

        l := dfs(root.Left)
        r := dfs(root.Right)
        if root.Left != nil && root.Right != nil {
            if l == 0 && r == 0 {
                ans++
            }
            if l == 0 || r == 0 {
                add = 1
            } else {
                add = 0
            }

            return 1
        }

        if l == 1 || r == 1 {
            return 1
        }

        return 0
    }
    
    l := dfs(root.Left)
    r := dfs(root.Right)
    if l == 0 || r == 0 {
        ans += add
    }

    return
}