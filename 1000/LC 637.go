/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) (ans []float64) {
    q := []*TreeNode{root}

    for len(q) > 0 {
        var sum, n float64
        n = float64(len(q))
        for k := len(q); k > 0; k-- {
            node := q[0]
            q = q[1:]
            sum += float64(node.Val)
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        ans = append(ans, sum / n)
    }

    return 
}
