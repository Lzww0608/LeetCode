import (
    "sort"
    "container/heap"
)
func smallestChair(times [][]int, targetFriend int) (ans int) {
    n := len(times)
    id := make([]int, n)
    for i := range id {
        id[i] = i
    }
    sort.Slice(id, func(i, j int) bool {
        return times[id[i]][0] < times[id[j]][0]
    })

    h1 := &hp{}
    h2 := &hp1{}
    for i := 0; i <= n; i++ {
        heap.Push(h1, i)
    }
    heap.Init(h1)
    heap.Init(h2)

    for _, i := range id {
        a, b := times[i][0], times[i][1]
        for h2.Len() > 0 && (*h2)[0][0] <= a {
            t := heap.Pop(h2).([]int)
            heap.Push(h1, t[1])
        } 
        ans = heap.Pop(h1).(int)
        if i == targetFriend {
            break
        }
        heap.Push(h2, []int{b, ans})
    }

    return
}

type hp []int
func (h hp) Less(i, j int) bool {return h[i] < h[j]}
func (h hp) Len() int {return len(h)}
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

type hp1 [][]int
func (h hp1) Less(i, j int) bool {return h[i][0] < h[j][0]}
func (h hp1) Len() int {return len(h)}
func (h hp1) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp1) Push(x any) {
    *h = append(*h, x.([]int))
}
func (h *hp1) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n-1]
    *h = old[:n-1]
    return
}