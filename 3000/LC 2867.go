var f [100005]int
var h [100005]int
func find(x int) int {
    if x != f[x] {
        f[x] = find(f[x])
    }
    return f[x]
}

func merge (x, y int) {
    a, b := find(x), find(y)
    if a != b {
        f[a] = b
        h[b] += h[a]
    }
}

func countPaths(n int, edges [][]int) (ans int64) {
    for i := 1; i <= n; i++ {
        f[i], h[i] = i, 1
    }


    isPrime := make([]int, n + 1)
    isPrime[1] = 1
    for i := 2; i <= n; i++ {
        if isPrime[i] == 0 {
            for j := i * 2; j <= n; j += i {
                isPrime[j] = 1
            }
        }
    }

    g := make([][]int, n + 1)
    for _, e := range edges {
        a, b := e[0], e[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
        if isPrime[a] == 1 && isPrime[b] == 1 {
            merge(a, b)
        }
    }

    for i := 1; i <= n; i++ {
        if isPrime[i] == 1 {
            continue
        }
        tmp := 1
        for _, e := range g[i] {
            if isPrime[e] == 1 {
                cnt := h[find(e)]
                ans += int64(tmp * cnt)
                tmp += cnt
            }
        }
    }
    return
}
