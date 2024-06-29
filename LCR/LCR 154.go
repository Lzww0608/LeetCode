/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }
    for node := head; node != nil; node = node.Next {
        newNode := &Node{node.Val, nil, nil}
        newNode.Next = node.Next
        node.Next = newNode
        node = node.Next
    }

    for node := head; node != nil; node = node.Next.Next {
        newNode := node.Next
        if node.Random == nil {
            newNode.Random = nil
        } else {
            newNode.Random = node.Random.Next
        }
    }

    newHead := head.Next
    for node := head; node != nil; node = node.Next {
        newNode := node.Next
        node.Next = node.Next.Next
        if newNode.Next != nil {
            newNode.Next = newNode.Next.Next
        }
    }

    return newHead
}
