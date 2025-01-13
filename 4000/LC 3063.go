/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func frequenciesOfElements(head *ListNode) *ListNode {
    cnt := make(map[int]int)
    for head != nil {
        cnt[head.Val]++
        head = head.Next
    }

    dummy := &ListNode{}
    p := dummy
    for _, v := range cnt {
        p.Next = &ListNode{}
        p = p.Next
        p.Val = v
    }
    return dummy.Next
}