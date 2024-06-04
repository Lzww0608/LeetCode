/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
    if root == nil {
        return root
    }
    dfs(root)
    return root
}

func dfs(root *Node) *Node {
    cur, last := root, &Node{}
    for cur != nil {
        next := cur.Next
        if cur.Child != nil {
            child := dfs(cur.Child)
            next := cur.Next
            cur.Next = cur.Child
            cur.Child.Prev = cur
            cur.Child = nil
            if next != nil {
                child.Next = next
                next.Prev = child
            }
            last = child
        } else {
            last = cur
        }
        cur = next
    }
    return last
}
