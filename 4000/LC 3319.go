/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargestPerfectSubtree(root *TreeNode, k int) int {
    a := []int{}
    var dfs func(*TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        l := dfs(root.Left)
        r := dfs(root.Right)
        if l < 0 || r < 0 || l != r {
            return -1
        }
        
        a = append(a, l + r + 1)
        return l + r + 1
    }
    dfs(root)

    sort.Slice(a, func(i, j int) bool {
        return a[i] > a[j]
    })
    if k > len(a) {
        return -1
    }

    return a[k - 1]
}