
import "slices"
func maximumMinimumPath(grid [][]int) (ans int) {
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    m, n := len(grid), len(grid[0])
    vis := make([]bool, m * n)

    bfs := func(limit int) bool {
        if grid[0][0] < limit {
            return false
        }
        q := [][]int{{0, 0}}
        for len(q) > 0 {
            cur := q[0]
            i, j := cur[0], cur[1]
            if i == m - 1 && j == n - 1 {
                return true
            }
            q = q[1:]
            if vis[i * n + j] {
                continue
            }
            vis[i * n + j] = true
            for _, dir := range dirs {
                x, y := i + dir[0], j + dir[1]
                if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] >= limit {
                    q = append(q, []int{x, y})
                }
            }
        }

        return false
    }

    mn, mx := slices.Min(grid[0]), slices.Max(grid[0])
    for i := 1; i < m; i++ {    
        mn = min(mn, slices.Min(grid[i]))
        mx = max(mx, slices.Max(grid[i]))
    }

    l, r := mn, mx + 1
    for l < r {
        mid := l + ((r - l) >> 1)
        clear(vis)
        if bfs(mid) {
            ans = mid
            l = mid + 1
        } else {
            r = mid
        }
    }

    return 
}