
import "math"
func minTimeToReach(moveTime [][]int) int {
    dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    n, m := len(moveTime), len(moveTime[0])
    dis := make([]int, m * n)
    for i := 0; i < m * n; i++ {
        dis[i] = math.MaxInt32
    }
    dis[0] = 0
    h := &hp{{0, 0}}
    heap.Init(h)

    for h.Len() > 0 {
        cur := heap.Pop(h).(pair)
        i, j, d := cur.to / m, cur.to % m, cur.time
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x < 0 || x >= n || y < 0 || y >= m {
                continue
            }
            t := max(d + 1, moveTime[x][y] + 1)
            if t < dis[x * m + y] {
                dis[x * m + y] = t
                heap.Push(h, pair{x * m + y, t})
            }
        }
    }

    return dis[m * n - 1]
}

type pair struct {
    to, time int
}
type hp []pair
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].time < h[j].time}
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