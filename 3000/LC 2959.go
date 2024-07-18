const N int = 10
func numberOfSets(n int, maxDistance int, roads [][]int) (ans int) {
    dis := [N][N]int{}

    for s := 1; s < (1 << n); s++ {
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                dis[i][j] = math.MaxInt / 2
            }
            dis[i][i] = 0
        }

        for _, e := range roads {
            a, b, c := e[0], e[1], e[2]
            if ((s >> a) & 1) == 1 && ((s >> b) & 1) == 1{
                dis[a][b] = min(dis[a][b], c)
                dis[b][a] = min(dis[b][a], c)
            }
        }

        for k := 0; k < n; k++ {
            for j := 0; j < n; j++ {
                for i := 0; i < n; i++ {
                    dis[i][j] = min(dis[i][j], dis[i][k] + dis[k][j])
                }
            }
        }
        
        f := true
        for i := 0; f && i < n; i++ {
            for j := 0; j < n; j++ {
                if ((s >> i) & 1) == 1 && ((s >> j) & 1) == 1 {
                    if dis[i][j] > maxDistance {
                        f = false
                        break
                    }
                }
            }
        }
        if f {
            ans++
        }
    }

    ans++
    return
}