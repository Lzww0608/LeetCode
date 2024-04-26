type pair struct{right, height int}
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].height > h[j].height } //maxheap
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any) {
    *h = append(*h, x.(pair))
}
func (h *hp) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}


func getSkyline(buildings [][]int) (ans [][]int) {
    n := len(buildings)
    a := make([]int, 0, n * 2)
    for _, v := range buildings {
        a = append(a, v[0], v[1])
    }
    sort.Ints(a)

    idx := 0
    h := hp{}
    for _, x := range a {
        for idx < n && buildings[idx][0] <= x {
            heap.Push(&h, pair{buildings[idx][1], buildings[idx][2]})
            idx++
        }
        for h.Len() > 0 && h[0].right <= x {
            heap.Pop(&h)
        }

        mx := 0
        if h.Len() > 0 {
            mx = h[0].height
        }
        if len(ans) == 0 || mx != ans[len(ans)-1][1] {
            ans = append(ans, []int{x, mx})
        }
    }

    return
}