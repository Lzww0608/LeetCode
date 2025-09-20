const N int = 100_001;

var adj [N]bool
var vis [N]bool

func maxWeight(edges [][]int, value []int) (ans int) {
    n, m := len(value), len(edges)
    
    sort.Slice(edges, func(i, j int) bool {
        x := value[edges[i][0]] + value[edges[i][1]]
        y := value[edges[j][0]] + value[edges[j][1]]
        return x > y
    })

    g := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
    }

    for i := 0; i < n; i++ {
        for _, v := range g[i] {
            adj[v] = true
        }

        var v []int 
        for j := range m {
            if adj[edges[j][0]] && adj[edges[j][1]] {
                v = append(v, j)
            }
        }

        for j := 0; j < min(3, len(v)); j++ {
            vis[edges[v[j]][0]] = true
            vis[edges[v[j]][1]] = true
            sum := value[i] + value[edges[v[j]][0]] + value[edges[v[j]][1]]
            for _, k := range v {
                cur := sum 
                if !vis[edges[k][0]] {
                    cur += value[edges[k][0]]
                }
                if !vis[edges[k][1]] {
                    cur += value[edges[k][1]]
                }
                
                if cur > ans {
                    ans = cur
                }
            }

            vis[edges[v[j]][0]] = false
            vis[edges[v[j]][1]] = false
        }

        for _, v := range g[i] {
            adj[v] = false
        }

    }

    return 
}