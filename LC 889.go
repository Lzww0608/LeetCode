/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var m map[int]int
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
    m = make(map[int]int)
    for i, x := range postorder {
        m[x] = i
    } 
    return solve(preorder , postorder, 0, len(postorder) - 1, 0, len(postorder) - 1 )
}

func solve(preorder []int, postorder []int, pre_start int, pre_end int, post_start int, post_end int) *TreeNode {
    if pre_start > pre_end {
        return nil
    }
    root := &TreeNode{preorder[pre_start], nil, nil}
    if pre_start == pre_end {
        return root
    }
    left_tree_root := m[preorder[pre_start + 1]]
    sizeof_left_tree := m[preorder[pre_start + 1]] - post_start + 1
    root.Left = solve(preorder, postorder, pre_start + 1, pre_start + sizeof_left_tree, post_start, left_tree_root) 
    root.Right = solve(preorder, postorder,  pre_start + sizeof_left_tree + 1, pre_end, left_tree_root + 1 , post_end - 1)
    return root
}