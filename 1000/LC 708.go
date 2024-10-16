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
        return ans
    } else if aNode.Next == aNode {
        aNode.Next = &Node{x, aNode}
        return aNode
    }

    p := aNode
    for {
        if p.Next == aNode || p.Val == x || (p.Val < x && (p.Next.Val >= x || p.Next.Val < p.Val)) || (p.Val > x && p.Next.Val > x && p.Next.Val < p.Val) {
            tmp := &Node{x, p.Next}
            p.Next = tmp
            break
        } else {
            p = p.Next
        }

    }

    return aNode
}
