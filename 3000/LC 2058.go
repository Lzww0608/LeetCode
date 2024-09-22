

import "math"
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {
    ans := []int{math.MaxInt32, -1}
    pre := -1
    firstPos, lastPos := math.MaxInt32, -1

    for p, k := head, 0; p.Next != nil; p = p.Next {
        cur := p.Val
        if pre != -1 {
            next := p.Next.Val
            if cur > pre && cur > next || cur < pre && cur < next {
                if firstPos != math.MaxInt32 {
                    ans[1] = max(ans[1], k - firstPos)
                    ans[0] = min(ans[0], k - lastPos)
                }
                lastPos = k
                firstPos = min(firstPos, k)
            }
        }
        pre = cur
        k++
    }
    if ans[1] == -1 {
        return []int{-1, -1}
    }

    return ans
}