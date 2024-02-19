/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    m := map[int]int{}
    for i, x := range inorder {
        m[x] = i
    }
    var dfs func(int, int, int) *TreeNode
    dfs = func(pre_start, pre_end, in_start int) *TreeNode {
        if pre_start > pre_end {
            return nil
        }
        root := &TreeNode{preorder[pre_start], nil, nil}
        if pre_start == pre_end {
            return root
        }
        leftTreeRoot := m[preorder[pre_start]]
        leftTreeSize := leftTreeRoot - in_start
        root.Left = dfs(pre_start + 1, pre_start + leftTreeSize, in_start)
        root.Right = dfs(pre_start + leftTreeSize + 1, pre_end, in_start + leftTreeSize + 1)
        return root
    }
    return dfs(0, len(preorder) - 1, 0)
}