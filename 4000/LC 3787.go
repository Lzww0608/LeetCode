func findSpecialNodes(n int, edges [][]int) string {
    g := make([][]int, n)
    for _, v := range edges {
        a, b := v[0], v[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    dis := make([]int, n)
    var dfs func(int, int, int)
    dfs = func(u, fa, d int) {
        dis[u] = d 
        for _, v := range g[u] {
            if v != fa {
                dfs(v, u, d + 1)
            }
        }
    }
    dfs(0, -1, 0)

    ans := []int{}
    mx := 0
    for i, x := range dis {
        if x > mx {
            mx = x 
            ans = ans[:0]
            ans = append(ans, i)
        } else if x == mx {
            ans = append(ans, i)
        }
    }

    s := make([]byte, n)
    for _, x := range ans {
        s[x] = '1'
    }
    dis = make([]int, n)
    dfs(ans[0], -1, 0)
    for i, x := range dis {
        if x > mx {
            mx = x  
            ans = ans[:0]
            ans = append(ans, i)
        } else if x == mx {
            ans = append(ans, i)
        }
    }
    for _, x := range ans {
        s[x] = '1'
    }
    for i := range s {
        if s[i] != '1' {
            s[i] = '0'
        }
    }

    return string(s)
}