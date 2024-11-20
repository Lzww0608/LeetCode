func maxTwoEvents(a [][]int) (ans int) {
    sort.Slice(a, func(i, j int) bool {
        return a[i][0] < a[j][0] || (a[i][0] == a[j][0] && a[i][1] < a[j][1])
    })

    h := &hp{}
    mx := 0
    for _, v := range a {
        l, r, c := v[0], v[1], v[2]
        for h.Len() > 0 && (*h)[0].end < l {
            t := heap.Pop(h).(pair)
            mx = max(mx, t.cost)
        }
        ans = max(ans, c + mx)
        heap.Push(h, pair{r, c})
    }

    return 
}

type pair struct {
    end, cost int
}

type hp []pair
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].end < h[j].end}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(pair))
}
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n-1]
    *h = old[:n-1]
    return 
}
