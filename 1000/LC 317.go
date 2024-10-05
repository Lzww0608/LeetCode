
import "math"
func shortestDistance(grid [][]int) int {
    dirs := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
    m, n := len(grid), len(grid[0])
    cnt := 0
    reach := make([]int, m * n)
    dist := make([]int, m * n)

    bfs := func(start_x, start_y int) {
        vis := make([]bool, m * n)
        q := [][]int{{start_x, start_y}}
        vis[start_x * n + start_y] = true
        d := 1
        for len(q) > 0 {
            for sz := len(q); sz > 0; sz-- {
                i, j := q[0][0], q[0][1]
                q = q[1:]
                for _, dir := range dirs {
                    x, y := i + dir[0], j + dir[1]
                    if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != 0 {
                        continue
                    }
                    if !vis[x * n + y] {
                        vis[x * n + y] = true
                        reach[x * n + y]++
                        dist[x * n + y] += d
                        q = append(q, []int{x, y})
                    }
                }
            }
            d++
        }
    }

    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 1 {
                cnt++
                bfs(i, j)
            }
        }
    }

    ans := math.MaxInt32
    for i := range grid {
        for j := range grid[i] {
            k := i * n + j
            if grid[i][j] == 0 && reach[k] == cnt {
                ans = min(ans, dist[k])
            }
        }
    }

    if ans == math.MaxInt32 {
        return -1
    }

    return ans
}