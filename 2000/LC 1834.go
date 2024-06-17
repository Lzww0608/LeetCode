func getOrder(tasks [][]int) (ans []int) {
    n := len(tasks)
    time := make([]int, n)
    pos := make([]int, n)
    for i, v := range tasks {
        time[i] = v[0]
        pos[i] = i
    }
    sort.Slice(pos, func(i, j int) bool {
        return time[pos[i]] < time[pos[j]] || (time[pos[i]] == time[pos[j]] && tasks[pos[i]][1] < tasks[pos[j]][1])
    })

    cur := 0
    h := hp{}
    heap.Init(&h)
    i := 0

    for i < n || h.Len() > 0 {
        for i < n && cur >= time[pos[i]] {
            x := pos[i]
            heap.Push(&h, []int{tasks[x][1], x})
            i++
        }

        if h.Len() > 0 {
            t := heap.Pop(&h).([]int)
            ans = append(ans, t[1])
            cur += t[0]
        } else {
            if i < n {
                cur = time[pos[i]]
            }
        }
    }

    return
}

type hp [][]int
func (h hp) Len() int {
    return len(h)
}
func (h hp) Less(i, j int) bool {
    return h[i][0] < h[j][0] || (h[i][0] == h[j][0] && h[i][1] < h[j][1])
}
func (h hp) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
func (h *hp) Push(x any) {
    *h = append(*h, x.([]int))
}
func (h *hp) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}
