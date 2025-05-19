
import "math"
const N int = 26
func minMoves(matrix []string) int {
    dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    m, n := len(matrix), len(matrix[0])
    if matrix[m-1][n-1] == '#' {
        return -1
    }
    pos := make([][]int, N)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] >= byte('A') && matrix[i][j] <= byte('Z') {
                x := int(matrix[i][j] - byte('A'))
                pos[x] = append(pos[x], i * n + j)
            }
        }
    }

    dis := make([]int, m * n)
    for i := range dis {
        dis[i] = math.MaxInt32
    }
    dis[0] = 0
    q := [][]int{}
    q1 := [][]int{}
    q = append(q, []int{0, 0, 0})
    for len(q) > 0 || len(q1) > 0 {
        var cur []int
        if len(q1) > 0 {
            cur = q1[0]
            q1 = q1[1:]
        } else {
            cur = q[0]
            q = q[1:]
        }
        
        i, j, d := cur[0] / n, cur[0] % n, cur[1]
        if dis[i * n + j] < d {
            continue
        }
        if i == m - 1 && j == n - 1 {
            return d
        }

        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x < 0 || x >= m || y < 0 || y >= n || matrix[x][y] == '#' {
                continue
            }
            if dis[x * n + y] > d + 1 {
                dis[x * n + y] = d + 1
                q = append(q, []int{x * n + y, d + 1, cur[2]})
            }
        }
        
        if matrix[i][j] >= byte('A') && matrix[i][j] <= byte('Z') {
            t := matrix[i][j] - byte('A')
            if cur[2] & (1 << t) != 0 {
                continue
            }
            for _, v := range pos[t] {
                x, y := v / n, v % n 
                if x != i || y != j {
                    if dis[x * n + y] > d {
                        dis[x * n + y] = d 
                        q1 = append(q1, []int{x * n + y, d, cur[2] | (1 << t)})
                    }
                }
            }
        }
    }

    return -1
}