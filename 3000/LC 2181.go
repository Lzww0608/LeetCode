/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    dummy := &ListNode{-1, head}
    p := dummy

    k := 0
    for head != nil {
        if head.Val != 0 {
            k += head.Val
        } else if k != 0 {
            p.Next.Val = k
            k = 0
            p = p.Next
        }
        head = head.Next
    }
    p.Next = nil

    return dummy.Next
}