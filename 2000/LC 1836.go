/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
    cnt := make(map[int]int)
    for p := head; p != nil; p = p.Next {
        cnt[p.Val]++
    }

    dummy := &ListNode{-1, head}
    pre, p := dummy, head
    for p != nil {
        if cnt[p.Val] == 1 {
            pre, p = p, p.Next
        } else {
            pre.Next, p = p.Next, p.Next
        }
    }
    //pre.Next = nil

    return dummy.Next
}