/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
    m := make(map[int]bool)
    for _, x := range nums {
        m[x] = true
    }

    dummy := &ListNode{-1, head}
    pre, p := dummy, head
    for p != nil {
        next := p.Next
        if m[p.Val] {
            pre.Next = next
        } else {
            pre = p 
        }
        p = next
    }

    return dummy.Next
}