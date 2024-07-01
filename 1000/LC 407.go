func trapRainWater(heightMap [][]int) (ans int) {
    m, n := len(heightMap), len(heightMap[0])
    if m <= 2 || n <= 2 {
        return
    }

    vis := make([]bool, m * n)
    h := &hp{}
    heap.Init(h)

    for i, row := range heightMap {
        for j, v := range row {
            if i == 0 || i == m - 1 || j == 0 || j == n - 1 {
                heap.Push(h, cell{v, i, j})
                vis[i * n + j] = true
            }
        }
    }

    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    for h.Len() > 0 {
        cur := heap.Pop(h).(cell)
        for _, dir := range dirs {
            x, y := cur.x + dir[0], cur.y + dir[1]
            if x >= 0 && x < m && y >= 0 && y < n && !vis[x * n + y] {
                if heightMap[x][y] < cur.h {
                    ans += cur.h - heightMap[x][y]
                }
                vis[x * n + y] = true
                heap.Push(h, cell{max(heightMap[x][y], cur.h), x, y})
            }
        }
    }

    return
}

type cell struct {h, x, y int}
type hp []cell
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].h < h[j].h}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(cell))
}

func (h *hp) Pop() any {
    old := *h 
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}
