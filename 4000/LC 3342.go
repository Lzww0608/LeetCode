
import (
	"container/heap"
	"math"
)
func minTimeToReach(moveTime [][]int) int {
    dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    m, n := len(moveTime), len(moveTime[0])
    dis := make([]int, m * n)
    for i := range dis {
        dis[i] = math.MaxInt32
    }
    dis[0] = 0

    h := &hp{{0, 0, 0, 0}}
    heap.Init(h)
    for h.Len() > 0 {
        cur := heap.Pop(h).(pair)
        i, j := cur.x, cur.y
        if i == m - 1 && j == n - 1 {
            return cur.weight
        }
        if cur.weight > dis[i * n + j] {
            continue
        }
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x < 0 || x >= m || y < 0 || y >= n {
                continue
            }
            w := max(moveTime[x][y], cur.weight) + cur.p + 1
            p := cur.p ^ 1
            if w < dis[x * n + y] {
                dis[x * n + y] = w 
                heap.Push(h, pair{x, y, w, p})
            }
        }
    }

    return dis[m * n - 1]
}

type pair struct {
    x, y, weight, p int
}
type hp []pair
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].weight < h[j].weight}
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