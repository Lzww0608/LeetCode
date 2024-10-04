import "math/bits"
import "math"
func countSubgraphsForEachDiameter(n int, edges [][]int) []int {
    ans := make([]int, n - 1)

    g := make([][]bool, n)
    dis := make([][]int, n)
    for i := range g {
        g[i] = make([]bool, n)
        dis[i] = make([]int, n)
        for j := range dis[i] {
            dis[i][j] = math.MaxInt32 / 2
        }
    }
    for _, e := range edges{
        u, v := e[0] - 1, e[1] - 1
        g[u][v] = true
        g[v][u] = true
        dis[v][u] = 1
        dis[u][v] = 1
    }

    for k := 0; k < n; k++ {
        for i := 0; i < n; i++ {
            if i == k {
                continue
            }
            for j := 0; j < n; j++ {
                if j == i || j == k {
                    continue
                }

                if dis[i][k] + dis[k][j] < dis[i][j] {
                    dis[i][j] = dis[i][k] + dis[k][j]
                }
            }
        }
    }


    m := 1 << n
    for s := 1; s < m; s++ {
        cnt := bits.OnesCount(uint(s))
        if cnt == 1 {
            continue
        }
        a := []int{}
        for i := 0; i < n; i++ {
            if s & (1 << i) != 0 {
                a = append(a, i)
            }
        }
        v, d := 0, 1
        for i := 0; i < len(a); i++ {
            for j := i + 1; j < len(a); j++ {
                d = max(d, dis[a[i]][a[j]])
                if g[a[i]][a[j]] {
                    v++
                }
            }
        }
        if v != len(a) - 1 || d >= math.MaxInt32 / 2 {
            continue
        }

        
        ans[d-1]++
    } 

    return ans
}