/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func spiralMatrix(m int, n int, head *ListNode) [][]int {
    ans := make([][]int, m)
    for i := range ans {
        ans[i] = make([]int, n)
        for j := range ans[i] {
            ans[i][j] = -1
        }
    }

    for k := 0; head != nil; k++ {
        i, j := k, k
        for head != nil && j < n - k {
            ans[i][j] = head.Val
            head = head.Next
            j++
        }
        i, j = i + 1, j - 1
        for head != nil && i < m - k {
            ans[i][j] = head.Val
            head = head.Next
            i++
        }
        i, j = i - 1, j - 1
        for head != nil && j >= k {
            ans[i][j] = head.Val
            head = head.Next
            j--
        }
        i, j = i - 1, j + 1
        for head != nil && i > k {
            ans[i][j] = head.Val
            head = head.Next
            i--
        }
    }

    return ans
}