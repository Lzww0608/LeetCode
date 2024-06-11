/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) (ans int) {
    type pair struct {
        node *TreeNode
        x int
    }
    q := []pair{pair{root, 0}}
    ans = 1
    for len(q) > 0 {
        mn, mx := math.MaxInt64, math.MinInt64
        for k := len(q); k > 0; k-- {
            node, x := q[0].node, q[0].x
            mn = min(mn, x)
            mx = max(mx, x)
            q = q[1:]
            if node.Left != nil {
                q = append(q, pair{node.Left, x * 2 + 1})
            }
            if node.Right != nil {
                q = append(q, pair{node.Right, x * 2 + 2})
            }
            
        }
        if mn == math.MaxInt32 || mx == math.MinInt64 {
            continue
        }
        ans = max(ans, int(mx - mn + 1))
    } 

    return 
}
