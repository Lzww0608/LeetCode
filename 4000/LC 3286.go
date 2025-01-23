
import "container/list"
func findSafeWalk(grid [][]int, health int) bool {
    dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    m, n := len(grid), len(grid[0])
    list := list.New()
    vis := make([][]bool, m * n)
    for i := range vis {
        vis[i] = make([]bool, health + 1)
    }

    if health == grid[0][0] & 1 {
        return false
    }
    list.PushBack([]int{0, 0, health - grid[0][0] & 1})
    for list.Len() > 0 {
        cur := list.Back().Value.([]int)
        list.Remove(list.Back())
        i, j, k := cur[0], cur[1], cur[2]
        if i == m - 1 && j == n - 1 {
            return true
        }
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x < 0 || x >= m || y < 0 || y >= n {
                continue
            }
            if grid[x][y] == 1 {
                if k > 1 && !vis[x * n + y][k-1] {
                    vis[x * n + y][k-1] = true
                    list.PushBack([]int{x, y, k - 1})
                }
            } else {
                if !vis[x * n + y][k] {
                    vis[x * n + y][k] = true
                    list.PushFront([]int{x, y, k})
                }
            }
        }
    }

    return false
}
