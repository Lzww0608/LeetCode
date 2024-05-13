/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    a, b := []byte{}, []byte{}
    for l1 != nil {
        a = append(a, byte('0' + l1.Val))
        l1 = l1.Next
    }
    for l2 != nil {
        b = append(b, byte('0' + l2.Val))
        l2 = l2.Next
    }

    dummy := &ListNode{-1, nil}
    m, n := len(a) - 1, len(b) - 1
    add := 0
    for m >= 0 || n >= 0 || add > 0 {
        x := add
        if m >= 0 {
            x += int(a[m] - '0')
            m--
        }
        if n >= 0 {
            x += int(b[n] - '0')
            n--
        }
        add = x / 10
        x %= 10
        node := &ListNode{x, dummy.Next}
        dummy.Next = node
    }

    return dummy.Next
}