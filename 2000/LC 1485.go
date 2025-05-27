/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Random *Node
 * }
 */

func copyRandomBinaryTree(root *Node) *NodeCopy {
    m := make(map[*Node]*NodeCopy)

    var dfs func(*Node) *NodeCopy
    dfs = func(root *Node) *NodeCopy {
        if root == nil {
            return nil 
        } else {
            if v, ok := m[root]; ok {
                return v
            }
            ans := &NodeCopy{Val: root.Val}
            m[root] = ans
            ans.Left = dfs(root.Left)
            ans.Right = dfs(root.Right)
            ans.Random = dfs(root.Random)
            return ans
        }
    }

    return dfs(root)
}