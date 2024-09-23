
import (
	"math"
)

func shortestDistance(maze [][]int, start []int, destination []int) int {
    dirs := [][]int{{1, 0},{-1, 0}, {0, 1}, {0, -1}}
    m, n := len(maze), len(maze[0])
    vis := make([]int, m * n)
    for i := range vis {
        vis[i] = math.MaxInt32
    }

    q := [][]int{start}
    vis[start[0] * n + start[1]] = 0
    for len(q) > 0 {
        node := q[0]
        q = q[1:]
        i, j := node[0], node[1]
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            for x >= 0 && x < m && y >= 0 && y < n && maze[x][y] == 0 {
                x, y = x + dir[0], y + dir[1]
            }
            x, y = x - dir[0], y - dir[1]
            dis := vis[i * n + j] + max(x - i, i - x, j - y, y - j)
            if vis[x * n + y] > dis {
                vis[x * n + y] = dis
                if x == destination[0] && y == destination[1] {
                    break
                }
                q = append(q, []int{x, y})
            }
        }
    } 
    
    dis := vis[destination[0] * n + destination[1]]
    if dis == math.MaxInt32 {
        return -1
    }

    return dis
}