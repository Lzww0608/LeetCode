func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
    m := len(queries)
    ans := make([]bool, m)

    fa := make([]int, n)
    sz := make([]int, n)
    for i := range n {
        fa[i] = i 
        sz[i] = 1
    }

    var find func(int) int 
    find = func(x int) int {
        if fa[x] != x {
            fa[x] = find(fa[x])
        }

        return fa[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx == ry {
            return
        }
        
        if sz[rx] < sz[ry] {
            rx, ry = ry, rx
        }
        fa[ry] = rx 
        sz[rx] += sz[ry]
    }

    for i := 1; i < n; i++ {
        if nums[i] - nums[i - 1] <= maxDiff {
            merge(i, i - 1)
        }
    }

    for i, q := range queries {
        x, y := q[0], q[1]
        if find(x) == find(y) {
            ans[i] = true
        }
    }

    return ans
}