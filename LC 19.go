/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, k int) *ListNode {
    dummy := &ListNode{-1, head}
    l, r := dummy, dummy
    for ; k > 0; k-- {
        r = r.Next
    }

    for r.Next != nil {
        r = r.Next
        l = l.Next
    }

    l.Next = l.Next.Next

    return dummy.Next
}