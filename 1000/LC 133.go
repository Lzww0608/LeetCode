/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }

    m := make(map[int]*Node)
    var dfs func(*Node) *Node
    dfs = func(node *Node) *Node {
        if v, ok := m[node.Val]; !ok {
            m[node.Val] = &Node{
                Val: node.Val,
                Neighbors: []*Node{},
            }
        } else {
            return v
        }
        cur := m[node.Val]
        
        for _, v := range node.Neighbors {
            cur.Neighbors = append(cur.Neighbors, dfs(v))
        }

        return cur
    }

    return dfs(node)
}