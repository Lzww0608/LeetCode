/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    leaf1, leaf2 := []int{}, []int{}
    var dfs func(leaf *[]int, root *TreeNode) 
    dfs = func(leaf *[]int, root *TreeNode) {
        if root == nil {
            return 
        } else if root.Left == nil && root.Right == nil {
            *leaf = append(*leaf, root.Val)
            return 
        }
        dfs(leaf, root.Left)
        dfs(leaf, root.Right)
    }
    dfs(&leaf1, root1)
    dfs(&leaf2, root2)
    if len(leaf1) != len(leaf2) {
        return false
    }
    for i := 0; i < len(leaf1); i++ {
        if leaf1[i] != leaf2[i] {
            return false
        }
    }
    return true
}
