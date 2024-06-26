/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) (ans []int) {
    var dfs func(*Node)
    dfs = func(root *Node) {
        if root == nil {
            return
        }
        ans = append(ans, root.Val)
        for _, x := range root.Children {
            dfs(x)
        }
    }
    dfs(root)
    return
}
