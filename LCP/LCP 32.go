
import "container/heap"
func processTasks(tasks [][]int) (ans int) {
    sort.Slice(tasks, func(i, j int) bool {
        return tasks[i][0] < tasks[j][0]
    })
    tasks = append(tasks, []int{1e9+1, 1e9+1, 1})
    h := &hp{}
    
    for _, v := range tasks {
        for h.Len() > 0 && (*h)[0][0] + ans < v[0] {
            if (*h)[0][0] + ans >= (*h)[0][1] {
                heap.Pop(h)
            } else {
                ans += min((*h)[0][1], v[0]) - ((*h)[0][0] + ans)
            }
        }

        heap.Push(h, []int{v[1] - (v[2] + ans) + 1, v[1] + 1})
    }

    return
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