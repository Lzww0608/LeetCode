func processQueries(n int, connections [][]int, queries [][]int) []int {
    id := make(map[int]int)
    k := 0

    g := make([][]int, n + 1)
    for _, v := range connections {
        g[v[0]] = append(g[v[0]], v[1])
        g[v[1]] = append(g[v[1]], v[0])
    }

    var dfs func(int, int, int)
    dfs = func(u, fa, t int) {
        if _, ok := id[u]; ok {
            return
        }
        id[u] = t 
        for _, v := range g[u] {
            if v != fa {
                dfs(v, u, t)
            }
        }
    }
    for i := 1; i <= n; i++ {
        if _, ok := id[i]; ok {
            continue
        }
        dfs(i, -1, k)
        k++
    }

    h := make([]hp, k)
    for i := range k {
        heap.Init(&h[i])
    }

    for i := 1; i <= n; i++ {
        heap.Push(&h[id[i]], i)
    }

    ans := []int{}
    offline := make([]bool, n + 1)
    for _, q := range queries {
        if q[0] == 1 {
            i := id[q[1]]
            if !offline[q[1]] {
                ans = append(ans, q[1])
                continue
            }
            for h[i].Len() > 0 && offline[h[i][0]] {
                heap.Pop(&h[i])
            }
            if h[i].Len() == 0 {
                ans = append(ans, -1)
            } else {
                ans = append(ans, h[i][0])
            }
        } else {
            offline[q[1]] = true
        }
    }

    return ans
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
    x = old[n - 1]
    *h = old[:n - 1]
    return
}