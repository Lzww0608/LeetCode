/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

func lowestCommonAncestor(p *Node, q *Node) *Node {
    a, b := p, q

    for p != q {
        p = p.Parent
        if p == nil {
            p = b 
        }

        q = q.Parent
        if q == nil {
            q = a
        }
    }

    return p
}