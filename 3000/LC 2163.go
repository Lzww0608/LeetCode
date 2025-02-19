
import (
	"container/heap"
	"math"
)
func minimumDifference(nums []int) int64 {
    sum := 0
    n := len(nums) / 3
    pre := make([]int, n + 1)
    suf := make([]int, n + 1)
    h := &hp{} 
    for i := 0; i < n; i++ {
        sum += nums[i]
        heap.Push(h, nums[i])
    }

    pre[0] = sum
    for i := n; i < n * 2; i++ {
        if nums[i] >= (*h)[0] {
            pre[i-n+1] = pre[i-n]
            continue
        }
        cur := heap.Pop(h).(int)
        heap.Push(h, nums[i])
        sum += nums[i] - cur
        pre[i-n+1] = sum
    }

    sum = 0 
    h = &hp{}
    for i := n * 2; i < n * 3; i++ {
        sum += nums[i]
        heap.Push(h, -nums[i])
    }
    suf[0] = sum
    for i := n * 2 - 1; i >= n; i-- {
        if nums[i] <= -(*h)[0] {
            suf[n*2-i] = suf[n*2-i-1]
            continue
        }
        cur := -heap.Pop(h).(int)
        heap.Push(h, -nums[i])
        sum += nums[i] - cur
        suf[n*2-i] = sum
    }

    ans := math.MaxInt
    for i := 0; i <= n; i++ {
        ans = min(ans, pre[i] - suf[n-i])
    }

    return int64(ans)
}

type hp []int
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i] > h[j]}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(int))
}
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n-1]
    *h = old[:n-1]
    return 
}