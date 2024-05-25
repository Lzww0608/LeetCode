/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deduceTree(preorder []int, inorder []int) *TreeNode {
    m := map[int]int{}
    for i, x := range inorder {
        m[x] = i
    }

    var construct func(int, int, int) *TreeNode
    construct = func(pre_start, pre_end, in_start int) *TreeNode {
        if pre_start > pre_end {
            return nil
        }
        root := &TreeNode{preorder[pre_start], nil, nil}
        if pre_start == pre_end {
            return root
        }
        
        rootIdx := m[preorder[pre_start]]
        left := rootIdx - in_start
        root.Left = construct(pre_start + 1, pre_start + left, in_start)
        root.Right = construct(pre_start + left + 1, pre_end, rootIdx + 1)
        return root
    }

    return construct(0, len(preorder) - 1, 0)
}