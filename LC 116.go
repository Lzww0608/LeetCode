/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
        return root
    }
    var dfs func(*Node, *Node)
    dfs = func(l *Node, r *Node) {
        if l == nil || r == nil {
            return 
        }
        l.Next = r 
        dfs(l.Left, l.Right)
        dfs(r.Left, r.Right)
        dfs(l.Right, r.Left)
    }
    dfs(root.Left, root.Right)
    return root
}
