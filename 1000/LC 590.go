/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) (ans []int) {
    if root == nil {
        return
    }
    q := []*Node{root}
    for len(q) > 0 {
        node := q[len(q)-1]
        q = q[:len(q)-1]
        ans = append(ans, node.Val)
        for _, x := range node.Children {
            q = append(q, x)
        }
    }
    for i, j := 0, len(ans) - 1; i < j; i, j = i + 1, j - 1 {
        ans[i], ans[j] = ans[j], ans[i]
    }
    return
}
