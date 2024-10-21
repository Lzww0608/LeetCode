func maximumSegmentSum(nums []int, removeQueries []int) []int64 {
    n := len(nums)
    fa := make([]int, n)
    sum := make([]int, n)
    for i, x := range nums {
        fa[i] = i
        sum[i] = x
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
            sum[ry] += sum[rx]
        }
    }

    ans := make([]int64, n)
    pre := sum[removeQueries[n-1]]
    vis := make([]bool, n)
    vis[removeQueries[n-1]] = true
    for i := n - 2; i >= 0; i-- {
        ans[i] = int64(pre)
        x := removeQueries[i]
        if x + 1 < n && vis[x+1] {
            merge(x, x + 1)
        }
        if x - 1 >= 0 && vis[x-1] {
            merge(x, x - 1)
        }
        vis[x] = true
        pre = max(pre, sum[find(x)])
    }

    return ans
}