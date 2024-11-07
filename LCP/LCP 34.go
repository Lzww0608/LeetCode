/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxValue(root *TreeNode, k int) int {
    ans := f(root, k)
    return slices.Max(ans)
}

func f(root *TreeNode, k int) []int {
    ans := make([]int, k + 1)
    if root == nil {
        return ans
    }

    l := f(root.Left, k)
    r := f(root.Right, k)
    
    for i := 0; i < k; i++ {
        for j := 0; i + j < k; j++ {
            ans[i+j+1] = max(ans[i+j+1], l[i] + r[j] + root.Val)
        }
    }

    for i := 0; i <= k; i++ {
        for j := 0; j <= k; j++ {
            ans[0] = max(ans[0], l[i] + r[j])
        }
    }

    return ans
}