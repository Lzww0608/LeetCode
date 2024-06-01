/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    m := map[int]int{}
    for i, x := range inorder {
        m[x] = i
    }
    var dfs func(int, int, int) *TreeNode
    dfs = func(post_start, post_end, in_end int) *TreeNode {
        if post_start > post_end {
            return nil 
        }
        root := &TreeNode{postorder[post_end], nil, nil}
        if post_start == post_end {
            return root
        }
        rootIndex := m[postorder[post_end]]
        rightTreeSize := in_end - rootIndex
        root.Left = dfs(post_start, post_end - rightTreeSize - 1, rootIndex - 1)
        root.Right = dfs( post_end - rightTreeSize, post_end - 1, in_end)
        return root
    }
    return dfs(0, len(postorder) - 1, len(inorder) - 1)
}
