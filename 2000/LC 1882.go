func assignTasks(servers []int, tasks []int) (ans []int) {
    h1 := hp{}
    h2 := hp{}
    for i, x := range servers {
        heap.Push(&h1, []int{x, i})
    }

    for i, x := range tasks {
        for h2.Len() > 0 && h2[0][0] <= i {
            tmp := heap.Pop(&h2).([]int)
            heap.Push(&h1, []int{tmp[1], tmp[2]})
        }

        if h1.Len() > 0 {
            tmp := heap.Pop(&h1).([]int)
            ans = append(ans, tmp[1])
            heap.Push(&h2, []int{i + x, tmp[0], tmp[1]})
        } else {
            tmp := heap.Pop(&h2).([]int)
            ans = append(ans, tmp[2])
            tmp[0] += x
            heap.Push(&h2, tmp)
        }
    }

    return 
}

type hp [][]int
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i][0] < h[j][0] || (h[i][0] == h[j][0] && (h[i][1] < h[j][1] || h[i][1] == h[j][1] && h[i][2] < h[j][2])) }
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.([]int))
}
func (h *hp) Pop() (x any) {
    old := *h
    n := len(old)
    x = old[n-1]
    *h = old[:n-1]
    return
}