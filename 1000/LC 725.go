/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) (ans []*ListNode) {
    n := 0
    for p := head; p != nil; p = p.Next {
        n++
    }
    ans = make([]*ListNode, k)

    a, b := n / k, n % k 
    p := head
    for i := 0; i < k; i++ {
        cnt := a 
        if b > 0 {
            cnt++
            b--
        }

        q := p 
        var pre *ListNode
        for j := 0; j < cnt; j++ {
            pre = p
            p = p.Next
        }
        if pre != nil {
            pre.Next = nil
        }
        ans[i] = q
    }

    return ans
}