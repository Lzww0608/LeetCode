/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumEvenGrandparent(root *TreeNode) int {
    ans := 0
    var dfs func(*TreeNode)
    dfs = func(node * TreeNode) {
        if node == nil {
            return 
        }
        if node.Val & 1 == 0 {
            if node.Left != nil {
                if node.Left.Left != nil {
                    ans += node.Left.Left.Val
                }
                if node.Left.Right != nil {
                    ans += node.Left.Right.Val
                }
            }
            if node.Right != nil {
                if node.Right.Left != nil {
                    ans += node.Right.Left.Val
                }
                if node.Right.Right != nil {
                    ans += node.Right.Right.Val
                }
            }
        }
        dfs(node.Left)
        dfs(node.Right)
    }
    dfs(root)
    return ans
}