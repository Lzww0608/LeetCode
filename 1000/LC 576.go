var mod int = 1e9 + 7
var dir [][]int = [][]int{[]int{0,1}, []int{0,-1}, []int{1,0}, []int{-1,0}}
func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
    g := make([][]int, m)
    for i := range g {
        g[i] = make([]int, n)
    }
    g[startRow][startColumn] = 1
    ans := 0

    for move := 0; move < maxMove; move++ {
        newG := make([][]int, m)
        for i := range newG {
            newG[i] = make([]int, n)
        }

        for i := 0; i < m; i++ {
            for j := 0; j < n; j++ {
                if g[i][j] > 0 {
                    for _, d := range dir {
                        x, y := i + d[0], j + d[1]
                        if x < 0 || x >= m || y < 0 || y >= n {
                            ans = (ans + g[i][j]) % mod
                        } else {
                            newG[x][y] = (newG[x][y] + g[i][j]) % mod
                        }
                    }
                }
            }
        }
        g = newG
    }

    return ans
}
