/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    dummy := &ListNode{0,head}
    for head != nil && head.Next != nil {
        if (head.Val == head.Next.Val) {
            head.Next = head.Next.Next
        } else {
            head = head.Next
        }
    }
    return dummy.Next
}