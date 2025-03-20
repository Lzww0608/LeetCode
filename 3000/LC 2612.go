func minReverseOperations(n int, p int, banned []int, k int) []int {
    fa := make([]int, n + 2)
    for i := range fa {
        fa[i] = i
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
        if rx != ry {
            fa[rx] = ry
        }
        return
    }

    merge(p, p + 2)
    for _, x := range banned {
        merge(x, x + 2)
    }

    q := []int{p}
    ans := make([]int, n)
    for i := range ans {
        ans[i] = -1
    }
    ans[p] = 0

    for len(q) > 0 {
        x := q[0]
        q = q[1:]
        mn := max(x - k + 1, k - x - 1)
        mx := min(x + k - 1, n * 2 - k - x - 1)
        for y := find(mn); y <= mx; y = find(y + 2) {
            ans[y] = ans[x] + 1
            q = append(q, y)
            merge(y, mx + 2)
        }
    }

    return ans
}