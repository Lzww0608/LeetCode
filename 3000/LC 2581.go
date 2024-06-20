type pair struct {
    x, y int
}

func rootCount(edges [][]int, guesses [][]int, k int) (ans int) {
    n := len(edges)
    g := make([][]int, n + 1)
    for _, e := range edges {
        a, b := e[0], e[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    s := make(map[pair]int, len(guesses))
    for _, p := range guesses {
        s[pair{p[0], p[1]}] = 1
    }

    cnt0 := 0
    var dfs func(int, int)
    dfs = func(x, fa int) {
        for _, y := range g[x] {
            if y != fa {
                //cnt0为以0为根时正确的猜测
                if s[pair{x, y}] == 1 {
                    cnt0++
                }
                dfs(y, x)
            }
        }
    }
    dfs(0, -1)

    var reroot func(int, int, int)
    reroot = func(x, fa, cnt int) {
        //cnt为以x为根时正确的猜测
        if cnt >= k {
            ans++
        }
        for _, y := range g[x] {
            if y != fa {
                reroot(y, x, cnt - s[pair{x, y}] + s[pair{y, x}])
            }
        }
    }
    reroot(0, -1, cnt0)
    return 
    
}
