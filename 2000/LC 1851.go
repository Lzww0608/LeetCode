func minInterval(intervals [][]int, queries []int) []int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] - intervals[i][0] < intervals[j][1] - intervals[j][0]
    })

    n := len(queries)
    ans := make([]int, n)
    for i := range ans {
        ans[i] = -1
    }
    q := make([][]int, n)
    for i, x := range queries {
        q[i] = []int{x, i}
    }
    sort.Slice(q, func(i, j int) bool {
        return q[i][0] < q[j][0]
    })

    fa := make([]int, n + 1)
    for i := range fa {
        fa[i] = i
    }

    var find func(int) int 
    find = func(x int) int {
        if x != fa[x] {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }

    for _, p := range intervals {
        l, r := p[0], p[1]
        d := r - l + 1
        i := sort.Search(n, func(i int) bool {
            return q[i][0] >= l
        })
        for i = find(i); i < n && q[i][0] <= r; i = find(i + 1) {
            ans[q[i][1]] = d
            fa[i] = i + 1
        }
    }    

    return ans
}