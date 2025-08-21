/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    n := len(nums)
    i := 0
    var dfs func(int) *TreeNode
    dfs = func(mx int) *TreeNode {
        var root *TreeNode = nil 
        for i < n && nums[i] <= mx {
            left := root
            root = &TreeNode{nums[i], left, nil}
            i++
            root.Right = dfs(root.Val)
        }
        return root
    }

    return dfs(math.MaxInt32)
}