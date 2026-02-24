/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelMedian(root *TreeNode, level int) int {
    a := []int{}
    q := []*TreeNode{root}
    for t := level + 1; t > 0 && len(q) > 0; t-- {
        for sz := len(q); sz > 0; sz-- {
            node := q[0]
            q = q[1:]
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }

            if t == 1 {
                a = append(a, node.Val)
            }
        }
    }
    if len(a) == 0 {
        return -1
    }
    sort.Ints(a)
    n := len(a)
    return a[n / 2]
}