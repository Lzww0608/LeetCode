/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func trainningPlan(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{-1, nil}
    p := dummy
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            p.Next = l1
            l1, p = l1.Next, p.Next
        } else {
            p.Next = l2
            l2, p = l2.Next, p.Next
        }
    }

    if l1 != nil {
        p.Next = l1
    } else if l2 != nil {
        p.Next = l2

    }

    return dummy.Next
}
