func minimumEffortPath(heights [][]int) int {
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    type pair struct {
        p, d int
    }
    m, n := len(heights), len(heights[0])
    dis := make([][]int, m)
    for i := range dis {
        dis[i] = make([]int, n)
        for j := range dis[i] {
            dis[i][j] = math.MaxInt32
        }
    }
    dis[0][0] = 0

    q := []pair{{0, 0}}
    for len(q) > 0 {
        cur := q[0]
        i, j := cur.p / n, cur.p % n 
        q = q[1:]
        if cur.d > dis[i][j] {
            continue
        }
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x < 0 || x >= m || y < 0 || y >= n || dis[x][y] <= max(cur.d, abs(heights[i][j] - heights[x][y])) {
                continue
            }
            dis[x][y] = max(cur.d, abs(heights[i][j] - heights[x][y]))
            q = append(q, pair{x * n + y, dis[x][y]})
        }
    }

    return dis[m - 1][n - 1]
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    
    return x
}