
import (
	"container/heap"
	"math"
)
func minimumCost(start []int, target []int, specialRoads [][]int) int {
    h := &hp{{start[0], start[1], 0}}
    heap.Init(h)
    n := len(specialRoads)
    dis := make([]int, n + 1)
    for i := 0; i < len(dis); i++ {
        dis[i] = math.MaxInt32
    }

    for h.Len() > 0 {
        cur := heap.Pop(h).(pair)
        i, j, d := cur.x, cur.y, cur.weight
        dis[n] = min(dis[n], d + abs(target[0] - i) + abs(target[1] - j))
        for k, v := range specialRoads {
            w := d + v[4] + abs(i - v[0]) + abs(j - v[1])
            if w < dis[k] {
                dis[k] = w
                heap.Push(h, pair{v[2], v[3], w}) 
            }
        }
    }

    return dis[n]
}

type pair struct {
    x, y, weight int
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

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}