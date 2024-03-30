/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) (ans int) {
    var dfs func(*Node) int
    dfs = func(root *Node) int {
        if root == nil {
            return 0
        }
        depth := 0
        for _, node := range root.Children {
            depth = max(depth, dfs(node))
        }
        ans = max(ans, depth + 1)
        return depth + 1
    }
    dfs(root)
    return 
}