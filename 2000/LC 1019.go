/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
    ans := []int{}
    st := []int{}
    for p := head; p != nil; p = p.Next {
        for len(st) > 0 && ans[st[len(st)-1]] < p.Val {
            i := st[len(st)-1]
            st = st[:len(st)-1]
            ans[i] = p.Val
        }
        st = append(st, len(ans))
        ans = append(ans, p.Val)
    }
    
    for _, i := range st {
        ans[i] = 0
    }
    return ans
}