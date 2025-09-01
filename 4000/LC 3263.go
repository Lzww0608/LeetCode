/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Prev *Node
 * }
 */

func toArray(head *Node) (ans []int) {
    for head.Prev != nil {
        head = head.Prev
    }

    for head != nil {
        ans = append(ans, head.Val)
        head = head.Next
    }

    return
}