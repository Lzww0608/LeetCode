/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargestLevelSum(root *TreeNode, k int) int64 {
    ans := []int64{}
    q := []*TreeNode{root}
    for len(q) > 0 {
        var sum int64
        for k := len(q); k > 0; k-- {
            node := q[0]
            q = q[1:]
            sum += int64(node.Val)
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node. Right != nil {
                q = append(q, node.Right)
            }
        }
        ans = append(ans, sum)
    }
    if k > len(ans) {
        return -1
    }
    sort.Slice(ans, func(i, j int) bool {
        return ans[i] > ans[j]
    })
    return ans[k-1]
}