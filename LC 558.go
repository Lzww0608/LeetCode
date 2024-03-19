/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func intersect(quadTree1 *Node, quadTree2 *Node) *Node {
    if quadTree1.IsLeaf {
        if quadTree1.Val {
            return &Node{true, true, nil, nil, nil, nil}
        } 
        return quadTree2
    } else if quadTree2.IsLeaf {
        return intersect(quadTree2, quadTree1)
    }

    a := intersect(quadTree1.TopLeft, quadTree2.TopLeft)
    b := intersect(quadTree1.TopRight, quadTree2.TopRight)
    c := intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
    d := intersect(quadTree1.BottomRight, quadTree2.BottomRight)

    if a.IsLeaf && b.IsLeaf && c.IsLeaf && d.IsLeaf && a.Val == b.Val && b.Val == c.Val && c.Val == d.Val {
        return &Node{a.Val, true, nil, nil, nil, nil}
    }

    return &Node{false, false, a, b, c, d}

}