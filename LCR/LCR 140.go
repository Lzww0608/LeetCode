/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func trainingPlan(head *ListNode, cnt int) *ListNode {
    p, n := head, 0
    for p != nil {
        p = p.Next
        n++
    }
    for i := 0; i < n - cnt; i++ {
        head = head.Next
    }

    return head
}
