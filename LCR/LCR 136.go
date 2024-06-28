/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
    dummy := &ListNode{-1, head}
    pre, p := dummy, head
    for p.Val != val {
        pre, p = p, p.Next
    }
    pre.Next = p.Next
    return dummy.Next
}
