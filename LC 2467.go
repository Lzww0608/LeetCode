func mostProfitablePath(edges [][]int, bob int, amount []int) int {
    n := len(amount)
    g := make([][]int, n)
    for _, e := range edges {
        a, b := e[0], e[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    path := []int{}
    dis := map[int]int{}
    var cal func(int, int) bool
    cal = func(root, fa int) bool {
        path = append(path, root)
        if root == bob {
            d := len(path)
            for i, x := range path {
                dis[x] = d - 1 - i
            }
            return true
        }
        for _, x := range g[root] {
            if x != fa {
               if cal(x, root) {
                    return true
               } 
            }
        }
        path = path[:len(path)-1]
        return false
    }
    cal(0, -1)

    ans := math.MinInt32
    var dfs func(int, int, int, int) 
    dfs = func(root, fa, d, sum int) {
        if v, ok := dis[root]; ok {
            if v == d {
                sum += amount[root] / 2
            } else if v > d {
                sum += amount[root]
            }
        } else {
            sum += amount[root]
        }
        if len(g[root]) == 1 && g[root][0] == fa {
            ans = max(ans, sum)
        }
        for _, x := range g[root] {
            if x != fa {
                dfs(x, root, d + 1, sum)
            }
        }
    }
    dfs(0, -1, 0, 0)

    return ans
}