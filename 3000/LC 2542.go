import (
    "sort"
    "container/heap"
)
func maxScore(a []int, b []int, k int) (ans int64) {
    n := len(a)
    id := make([]int, n)
    for i := 0; i < n; i++ {
        id[i] = i 
    }
    sort.Slice(id, func(i, j int) bool {
        return b[id[i]] > b[id[j]]
    })

    h := &hp{}
    sum := 0
    for _, i := range id {
        if h.Len() == k {
            cur := heap.Pop(h).([2]int)
            sum -= cur[1]
        }
        heap.Push(h, [2]int{i, a[i]})
        sum += a[i]


        if h.Len() == k {
            ans = max(ans, int64(b[i] * sum))
        }
    }

    return
}


type hp [][2]int 
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i][1] < h[j][1]}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.([2]int))
}
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n-1]
    *h = old[:n-1]
    return
}