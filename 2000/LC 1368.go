import (
    "container/list"
)

func minCost(grid [][]int) int {
    dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    m := len(grid)
    n := len(grid[0])
    vis := make([]bool, m * n)

    dq := list.New()
    dq.PushFront([]int{0, 0})  

    for dq.Len() > 0 {
        node := dq.Remove(dq.Front()).([]int)
        d, i, j := node[0], node[1] / n, node[1] % n
        if i == m - 1 && j == n - 1 {
            return d
        }
        if vis[node[1]] {
            continue
        }
        vis[node[1]] = true
        for k, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x < 0 || x >= m || y < 0 || y >= n {
                continue
            }
            if k + 1 == grid[i][j] {
                dq.PushFront([]int{d, x * n + y})  
            } else {
                dq.PushBack([]int{d + 1, x * n + y})  
            }
        }
    }

    return 0
}
