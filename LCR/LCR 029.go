/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

func insert(aNode *Node, x int) *Node {
    if aNode == nil {
        ans := &Node{x, nil}
        ans.Next = ans
        return  ans
    } else if aNode.Next == aNode {
        node := &Node{x, aNode}
        aNode.Next = node
        return aNode
    }

    for p := aNode; ; p = p.Next {
        if p.Val > p.Next.Val && (x <= p.Next.Val || x >= p.Val)||
        p.Val <= x && p.Next.Val >= x || p.Next == aNode {
            node := &Node{x, p.Next}
            p.Next = node
            break
        }
    }

    return aNode
}