func swimInWater(grid [][]int) (ans int) {
    n := len(grid)
    dirs := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
    vis := make([]bool, n * n)
    h := IntHeap{}
    heap.Init(&h)

    heap.Push(&h, []int{0, 0, grid[0][0]})

    for h.Len() > 0 {
        i, j, k := h[0][0], h[0][1], h[0][2]
        heap.Pop(&h)
        ans = max(ans, k)
        if i == n - 1 && j == n - 1 {
            break
        }

        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x >= 0 && x < n && y >= 0 && y < n && !vis[x * n + y] {
                vis[x * n + y] = true
                heap.Push(&h, []int{x, y, grid[x][y]})
            }
        }
    }

    return
}

type IntHeap [][]int
func (h IntHeap) Len() int {return len(h)}
func (h IntHeap) Less(i, j int) bool {return h[i][2] < h[j][2]}
func (h IntHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *IntHeap) Push(x any) {
    *h = append(*h, x.([]int))
}
func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}
