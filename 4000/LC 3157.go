/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minimumLevel(root *TreeNode) (ans int) {
    mx := math.MaxInt
    
    q := []*TreeNode{root}
    d := 1
    for len(q) > 0 {
        sum := 0
        for sz := len(q); sz > 0; sz-- {
            cur := q[0]
            q = q[1:]
            sum += cur.Val
            if cur.Left != nil {
                q = append(q, cur.Left)
            } 
            if cur.Right != nil {
                q = append(q, cur.Right)
            }
        }
        if sum < mx {
            mx, ans = sum, d
        }
        d++
    }

    return
}