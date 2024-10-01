
import (
	"container/list"
	"math"
)
func minimumObstacles(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    q := list.New()
    q.PushFront([]int{0, 0, 0})
    vis := make([]int, m * n)
    for i := range vis {
        vis[i] = math.MaxInt32
    }

    dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
    for q.Len() > 0 {
        cur := q.Remove(q.Front()).([]int)
        i, j, d := cur[0], cur[1], cur[2]
        if i == m - 1 && j == n - 1 {
            return d
        }
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x >= 0 && x < m && y >= 0 && y < n {
                newD := d
                if grid[x][y] == 1 {
                    newD += 1
                }
                if newD < vis[x * n + y] {
                    vis[x * n + y] = newD
                    if grid[x][y] == 1 {
                        q.PushBack([]int{x, y, newD})
                    } else {
                        q.PushFront([]int{x, y, newD})
                    }
                }
            }
        }

    }

    return -1
}