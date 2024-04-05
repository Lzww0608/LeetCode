func highestPeak(isWater [][]int) [][]int {
    type pair struct {
        x, y int
    }
    q := []pair{}
    for i := range isWater {
        for j := range isWater[i] {
            if isWater[i][j] == 1 {
                q = append(q, pair{i, j})
                isWater[i][j] = 0
            } else {
                isWater[i][j] = math.MaxInt32
            }
        }
    }

    k := 1
    m, n := len(isWater), len(isWater[0])
    dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    for len(q) > 0 {
        for t := len(q); t > 0; t-- {
            i, j := q[0].x, q[0].y
            q = q[1:]
            for _, dir := range dirs {
                x, y := i + dir[0], j + dir[1]
                if x >= 0 && x < m && y >= 0 && y < n {
                    if isWater[x][y] > k {
                        q = append(q, pair{x, y})
                    }
                    isWater[x][y] = min(isWater[x][y], k)
                }
            }
        }
        k++
    }

    return isWater
} 