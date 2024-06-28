func orangesRotting(grid [][]int) int {
    type pair struct {
        x, y int
    }
    m, n := len(grid), len(grid[0])
    ans, cnt := -1, 0
    q := []pair{}
    for i := range grid {
        for j, x := range grid[i] {
            if x == 1 {
                cnt++
            } else if x == 2 {
                q = append(q, pair{i, j})
            }
        }
    }
    if cnt == 0 {
        return 0
    }
    
    dirs := [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, -1}, []int{0, 1}}

    for len(q) > 0 {
        ans++
        for k := len(q); k > 0; k-- {
            t := q[0]
            q = q[1:]
            i, j := t.x, t.y
            for _, dir := range dirs {
                x, y := i + dir[0], j + dir[1]
                if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
                    cnt--
                    grid[x][y] = 2
                    q = append(q, pair{x, y})
                }
            }
        }
    }
    if cnt > 0 {
        return -1
    }
    return ans
}
