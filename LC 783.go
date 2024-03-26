/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
    dif := math.MaxInt32
    x := -1

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        if x == -1 {
            x = root.Val
        } else {
            dif = min(dif, root.Val - x)
            x = root.Val
        }
        dfs(root.Right)
    }
    dfs(root)


    return dif
}