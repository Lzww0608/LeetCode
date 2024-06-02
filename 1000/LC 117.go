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
    q := []*Node{root}
    for len(q) > 0 {
        var preNode *Node 
        for k := len(q); k > 0; k-- {
            node := q[0]
            q = q[1:]
            if preNode != nil {
                preNode.Next = node
            }
            preNode = node
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
    }
    return root
}
