/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func cloneTree(root *Node) *Node {
	
    var dfs func(*Node) *Node 
    dfs = func(root *Node) *Node {
        if root == nil {
            return nil
        }
        ans := &Node{Val: root.Val}
        for _, v := range root.Children {
            ans.Children = append(ans.Children, dfs(v))
        }
        return ans
    }

    return dfs(root)
}