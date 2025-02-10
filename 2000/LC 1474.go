/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNodes(head *ListNode, m int, n int) *ListNode {
    dummy := &ListNode{-1, head}

    pre, p := dummy, head
    for p != nil {
        for i := 0; i < m && p != nil; i++ {
            pre, p = p, p.Next
        }
        for i := 0; i < n && p != nil; i++ {
            p = p.Next
        }
        pre.Next = p 
    }

    return dummy.Next
}