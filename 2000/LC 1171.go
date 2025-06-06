/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
    m := make(map[int]*ListNode)
    sum := 0

    dummy := &ListNode{0, head}
    m[0] = dummy
    for p := head; p != nil; p = p.Next {
        sum += p.Val
        m[sum] = p
    }

    sum = 0
    for p := dummy; p != nil; p = p.Next {
        sum += p.Val
        p.Next = m[sum].Next
    }

    return dummy.Next
}