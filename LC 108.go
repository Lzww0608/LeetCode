/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    
    var construct func(l, r int) *TreeNode 
    construct = func(l, r int) *TreeNode {
        if l > r {
            return nil
        }
        mid := l + ((r - l) >> 1)
        node := &TreeNode{nums[mid], nil, nil}
        node.Left = construct(l, mid - 1)
        node.Right = construct(mid + 1, r)
        return node
    }
    return construct(0, len(nums) - 1)
}