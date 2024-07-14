/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    n := len(nums)

    var dfs func(int, int) *TreeNode
    dfs = func(l, r int) *TreeNode {
        if l > r {
            return nil
        } else if l == r {
            return &TreeNode{nums[l], nil, nil}
        }
        mid := l + ((r - l) >> 1)
        node := &TreeNode{nums[mid], nil, nil}
        node.Left = dfs(l, mid - 1)
        node.Right = dfs(mid + 1, r)
        return node
    }

    return dfs(0, n - 1)
}