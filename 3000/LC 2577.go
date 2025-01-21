const N int = 100_005
func minimumTime(grid [][]int) int {
    dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    m, n := len(grid), len(grid[0])
    if (m == 1 || m > 1 && grid[1][0] > 1) && (n == 1 || n > 1 && grid[0][1] > 1) {
        return -1
    }

    check := func(mid int) bool {
        vis := make([]int, m * n)
        for i := range vis {
            vis[i] = -1
        }
        vis[m * n - 1] = mid
        q := [][]int{{m - 1, n - 1, mid}}
        for len(q) > 0 {
            cur := q[0]
            i, j, cnt := cur[0], cur[1], cur[2]
            if i == 0 && j == 0 {
                return true
            }
            q = q[1:]
            cnt--
            for _, dir := range dirs {
                x, y := i + dir[0], j + dir[1]
                if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] <= cnt && vis[x * n + y] < cnt {
                    vis[x * n + y] = cnt
                    q = append(q, []int{x, y, cnt})
                } 
            }
        }
        return false
    }

    l, r := max(grid[m-1][n-1], m + n - 2) - 1, N + m + n 
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid
        } else {
            l = mid
        }
    }

    return r + (r + m + n) & 1
}