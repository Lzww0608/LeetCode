/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
    dummy := &ListNode{-1, head}
    m := map[int]bool{}
    pre := dummy
    for head != nil {
        if m[head.Val] {
            pre.Next = head.Next
            head = head.Next
        } else {
            m[head.Val] = true
            pre, head = head, head.Next
        }
    }

    return dummy.Next
}