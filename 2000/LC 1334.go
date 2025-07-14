
import "math"
func findTheCity(n int, edges [][]int, distanceThreshold int) (ans int) {
    dis := make([][]int, n)
    for i := range dis {
        dis[i] = make([]int, n)
        for j := range dis[i] {
            dis[i][j] = math.MaxInt32
        }
    }

    for _, v := range edges {
        dis[v[0]][v[1]] = v[2]
        dis[v[1]][v[0]] = v[2]
    }

    for k := 0; k < n; k++ {
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                if dis[i][j] > dis[i][k] + dis[k][j] {
                    dis[i][j] = dis[i][k] + dis[k][j]
                    dis[j][i] = dis[i][j]
                }
            }
        }
    }

    cnt := n 
    for i := 0; i < n; i++ {
        cur := 0
        for j := 0; j < n; j++ {
            if i != j && dis[i][j] <= distanceThreshold {
                cur++
            }
        }
        if cur <= cnt {
            cnt, ans = cur, i
        }
    }

    return
}