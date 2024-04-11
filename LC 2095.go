/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    dummy := &ListNode{-1, head}
    //pre := dummy
    slow, fast := dummy, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    if slow.Next != nil {
        slow.Next = slow.Next.Next
    }
    return dummy.Next
}