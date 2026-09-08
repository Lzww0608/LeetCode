func minCost(grid [][]int, k int) int {
    m, n := len(grid), len(grid[0])
    dirs := [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}
    f := make([][]int, m * n)
    for i := range f {
        f[i] = make([]int, (k + 1) * 4)
        for j := range f[i] {
            f[i][j] = math.MaxInt32
        }
    }

    h := &hp{}
    for i := range 4 {
        f[0][i * (k + 1) + k] = grid[0][0]
        heap.Push(h, node{0, 0, i, k, grid[0][0]})
    }

    for h.Len() > 0 {
        cur := heap.Pop(h).(node)
        if cur.c > f[cur.i * n + cur.j][cur.d * (k + 1) + cur.k] {
            continue
        }

        if cur.i == m - 1 && cur.j == n - 1 {
            return cur.c
        }
        for p, dir := range dirs {
            i, j := cur.i + dir[0], cur.j + dir[1]
            if i < 0 || i >= m || j < 0 || j >= n {
                continue
            }
            c := cur.c + grid[i][j]
            tmp_k := cur.k
            if p != cur.d {
                tmp_k--
                if tmp_k < 0 {
                    continue
                }
            }
            if c < f[i * n + j][p * (k + 1) + tmp_k] {
                f[i * n + j][p * (k + 1) + tmp_k] = c 
                heap.Push(h, node{i, j, p, tmp_k, c})
            }
        }
    }

    return -1
}

type node struct {
    i, j, d, k, c int
}
type hp []node 
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].c < h[j].c}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(node))
}
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n - 1]
    *h = old[:n - 1]
    return
}