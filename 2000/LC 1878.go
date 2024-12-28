import (
    "container/heap"
    "sort"
)
func getBiggestThree(grid [][]int) []int {
    m, n := len(grid), len(grid[0])
    exist := make(map[int]bool)
    h := &hp{}
    for d := 1; d <= min((m + 1) / 2, (n + 1) / 2); d++ {
        for i := 0; i + (d - 1) * 2 < m; i++ {
            for j := d - 1; j + d <= n; j++ {
                sum := 0
                for x, y := i, j; x <= i + (d - 1) * 2; x++ {
                    if x == i || x == i + (d - 1) * 2 {
                        sum += grid[x][j]
                    } else {
                        sum += grid[x][y] + grid[x][j-y+j]
                    }

                    if x < i + d -1 {
                        y--
                    } else {
                        y++
                    }
                } 
                if !exist[sum] {
                    exist[sum] = true
                    heap.Push(h, sum)
                }
                if h.Len() > 3 {
                    heap.Pop(h)
                }
            }
        }
    }
    sort.Slice(*h, func(i, j int) bool {
        return (*h)[i] > (*h)[j]
    })
    return *h
}

type hp []int
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i] < h[j]}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(int))
}
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n-1]
    *h = old[:n-1]
    return
}