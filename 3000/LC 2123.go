func minimumOperations(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    N := 0

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                N++
                grid[i][j] = N
            }
        }
    }

    G := make([][][2]int, N + 5)
    var rem []int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] != 0 {
                if i != 0 && grid[i-1][j] != 0 {
                    if ((i + j) & 1 != 0) {
                        G[grid[i][j]] = append(G[grid[i][j]], [2]int{grid[i-1][j], len(rem)})
                        rem = append(rem, 1)
                        G[grid[i-1][j]] = append(G[grid[i-1][j]], [2]int{grid[i][j], len(rem)})
                        rem = append(rem, 0)
                    } else {
                        G[grid[i][j]] = append(G[grid[i][j]], [2]int{grid[i-1][j], len(rem)})
                        rem = append(rem, 0)
                        G[grid[i-1][j]] = append(G[grid[i-1][j]], [2]int{grid[i][j], len(rem)})
                        rem = append(rem, 1)
                    }
                }

                if j != 0 && grid[i][j-1] != 0 {
                    if (i + j) & 1 != 0 {
                        G[grid[i][j]] = append(G[grid[i][j]], [2]int{grid[i][j-1], len(rem)})
                        rem = append(rem, 1)
                        G[grid[i][j-1]] = append(G[grid[i][j-1]], [2]int{grid[i][j], len(rem)})
                        rem = append(rem, 0)
                    } else {
                        G[grid[i][j]] = append(G[grid[i][j]], [2]int{grid[i][j-1], len(rem)})
                        rem = append(rem, 0)
                        G[grid[i][j-1]] = append(G[grid[i][j-1]], [2]int{grid[i][j], len(rem)})
                        rem = append(rem, 1)
                    }
                }

                if (i + j) & 1 != 0 {
                    G[0] = append(G[0], [2]int{grid[i][j], len(rem)})
                    rem = append(rem, 1)
                    G[grid[i][j]] = append(G[grid[i][j]], [2]int{0, len(rem)})
                    rem = append(rem, 0)
                } else {
                    G[grid[i][j]] = append(G[grid[i][j]], [2]int{N + 1, len(rem)})
                    rem = append(rem, 1)
                    G[N+1] = append(G[N+1], [2]int{grid[i][j], len(rem)})
                    rem = append(rem, 0)
                }
            }
        }
    }

    res := 0
    dis := make([]int, N + 5)
    for {
        for i := range dis {
            dis[i] = 0
        }
        dis[0] = 1
        q := []int{0}
        for len(q) > 0 {
            cur := q[0]
            q = q[1:]
            for _, v := range G[cur] {
                ne, p := v[0], v[1]
                if dis[ne] == 0 && rem[p] != 0 {
                    dis[ne] = dis[cur] + 1
                    q = append(q, ne)
                }
            }
        }

        if dis[N + 1] == 0 {
            break
        }

        var dfs func(int, int) int 
        dfs = func(i, x int) int {
            if i == N + 1 || x == 0 {
                return x
            }
            ret := 0
            for ne := 0; ne < len(G[i]) && ret < x; ne++ {
                if dis[G[i][ne][0]] == dis[i] + 1 && rem[G[i][ne][1]] != 0 {
                    cost := dfs(G[i][ne][0], min(x, rem[G[i][ne][1]]))
                    rem[G[i][ne][1]] -= cost
                    rem[G[i][ne][1] ^ 1] += cost
                    ret += cost
                }

            }
            return ret
        }
        res += dfs(0, 1_000_000_000)
    }

    return res
}