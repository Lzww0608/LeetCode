/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) (ans []string) {
    var dfs func(*TreeNode, string)
    dfs = func(root *TreeNode, s string) {
        s = s + strconv.Itoa(root.Val) + "->"
        if root.Left == nil && root.Right == nil {
            s = s[:len(s)-2]
            ans = append(ans, s)
            return
        }

        if root.Left != nil {
            dfs(root.Left, s)
        }
        if root.Right != nil {
            dfs(root.Right, s)
        }
    }
    dfs(root, "")

    return 
}