/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, nums []int) (ans int) {
    pos := make(map[int]bool)
    for _, x := range nums {
        pos[x] = true
    }

    f := false
    for p := head; p != nil; p = p.Next {
        if !pos[p.Val] {
            f = false
        } else if !f {
            ans++
            f = true
        }
    }

    return 
}