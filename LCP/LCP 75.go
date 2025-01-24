
import "math"
func challengeOfTheKeeper(maze []string) int {
    dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    m, n := len(maze), len(maze[0])
    var sx, sy, tx, ty int
    for i, row := range maze {
        for j, c := range row {
            if c == 'S' {
                sx, sy = i, j
            } else if c == 'T' {
                tx, ty = i, j
            }
        }
    }

    dis := make([][]int, m)
    for i := range dis {
        dis[i] = make([]int, n)
        for j := range dis[i] {
            dis[i][j] = math.MaxInt32
        }
    }

    dis[tx][ty] = 0
    q := [][]int{{tx, ty}}
    d := 1
    for len(q) > 0 {
        for sz := len(q); sz > 0; sz-- {
            cur := q[0]
            q = q[1:]
            i, j := cur[0], cur[1]
            for _, dir := range dirs {
                x, y := i + dir[0], j + dir[1]
                if x >= 0 && x < m && y >= 0 && y < n && 
                    maze[x][y] != '#' && dis[x][y] == math.MaxInt32 {
                    dis[x][y] = d
                    q = append(q, []int{x, y})
                }
            }
        }
        d++
    }

    if dis[sx][sy] == math.MaxInt32 {
        return -1
    }

    vis := make([]int, m * n)
    for i := range vis {
        vis[i] = -1
    }

    var dfs func(int, int, int) bool
    dfs = func(mid, x, y int) bool {
        if x < 0 || x >= m || y < 0 || y >= n || vis[x * n + y] == mid || maze[x][y] == '#' {
            return false
        }
        if maze[x][y] == 'T' {
            return true
        }
        vis[x * n + y] = mid
        if maze[x][y] == '.' && 
        (maze[m-x-1][y] != '#' && dis[m-1-x][y] > mid || 
        maze[x][n-1-y] != '#' && dis[x][n-1-y] > mid) {
            return false
        }
        for _, dir := range dirs {
            if dfs(mid, x + dir[0], y + dir[1]) {
                return true
            }
        }
        return false
    }

    l, r := -1, m * n + 1
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if dfs(mid, sx, sy) {
            r = mid
        } else {
            l = mid
        }
    }
    if r > m * n {
        return -1
    }
    return r
}