func minCost(m int, n int, penalty [][]int) int64 {
    dis := make([][2]int, m * n)
    for i := range dis {
        dis[i][0], dis[i][1] = math.MaxInt / 2, math.MaxInt / 2
    }
    dis[0][1] = 1

    h := &hp{}
    heap.Push(h, pair{0, 1, 1})
    for h.Len() > 0 {
        cur := heap.Pop(h).(pair)
        i, j := cur.x / n, cur.x % n
        if cur.c > dis[cur.x][cur.d] {
            continue
        }
        if i == m - 1 && j == n - 1 {
            return int64(cur.c)
        }

        if dis[cur.x][cur.d ^ 1] > cur.c + penalty[i][j] {
            dis[cur.x][cur.d ^ 1] = cur.c + penalty[i][j]
            heap.Push(h, pair{cur.x, dis[cur.x][cur.d ^ 1], cur.d ^ 1})
        } 
        
        for k, dir := range [4][2]int{{0, -1}, {-1, 0}, {1, 0}, {0, 1}} {
            x, y := i + dir[0], j + dir[1]
            if x < 0 || x >= m || y < 0 || y >= n {
                continue
            }
            c := (x + 1) * (y + 1)
            if k / 2 != cur.d {
                c += penalty[i][j]
            } 
            if dis[x * n + y][cur.d ^ 1] > cur.c + c {
                dis[x * n + y][cur.d ^ 1] = cur.c + c
                heap.Push(h, pair{x * n + y, cur.c + c, cur.d ^ 1})
            } 
        }
    } 

    return -1
}

type pair struct {
    x, c, d int 
}
type hp []pair 
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].c < h[j].c}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(pair))
}
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n - 1]
    *h = old[:n - 1]
    return 
}