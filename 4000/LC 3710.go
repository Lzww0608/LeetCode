func maxPartitionFactor(points [][]int) int {
    n := len(points)
    if n == 2 {
        return 0
    }

    check := func(mid int) bool {
        col := make([]int, n)

        var dfs func(int, int) bool 
        dfs = func(x, c int) bool {
            col[x] = c 
            for y := range n {
                if x == y || abs(points[x][0] - points[y][0]) + abs(points[x][1] - points[y][1]) >= mid {
                    continue
                }
                if col[y] == c || col[y] == 0 && !dfs(y, -c) {
                    return false
                }
            }

            return true
        }

        for i := range n {
            if col[i] == 0 && !dfs(i, 1) {
                return false
            }
        }

        return true
    }


    l, r := 0, int(1e9)
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            l = mid
        } else {
            r = mid
        }
    }

    return l
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}