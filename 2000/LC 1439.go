
func find(a, b []int, k int) []int {
    m, n := len(a), len(b)
    ans := make([]int, 0, min(k, n * m))
    h := &hp{{a[0] + b[0], 0, 0}}
    heap.Init(h)
    for h.Len() > 0 && len(ans) < k {
        p := heap.Pop(h).([]int)
        i, j := p[1], p[2]
        ans = append(ans, a[i] + b[j])
        if j == 0 && i + 1 < m {
            heap.Push(h, []int{a[i+1] + b[0], i + 1, 0})
        }
        if j + 1 < n {
            heap.Push(h, []int{a[i] + b[j+1], i, j + 1})
        }
    }

    return ans
}

func kthSmallest(mat [][]int, k int) int {
    a := []int{0}
    for _, row := range mat {
        a = find(a, row, k)
    }

    return a[k-1]
}

type hp [][]int
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i][0] < h[j][0]}
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