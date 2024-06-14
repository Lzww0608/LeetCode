var f = [55]int{}
var g = [55]int{}
var h = [55]int{}

func find(x int) int {
    if x != f[x] {
        f[x] = find(f[x])
    }
    return f[x]
}

func merge(x, y int) {
    rx, ry := find(x), find(y)
    if rx == ry {
        return
    }

    f[rx] = ry
    g[ry] += g[rx]
}

func countCompleteComponents(n int, edges [][]int) (ans int) {
    for i := 0; i < n; i++ {
        f[i], g[i], h[i] = i, 1, 0
    }

    for _, e := range edges {
        merge(e[0], e[1])
    }

    for _, e := range edges {
        h[find(e[0])]++
    }

    for i := 0; i < n; i++ {
        if i == find(i) && h[i] * 2 == g[i] * (g[i] - 1) {
            ans++
        }
    }

    return
}
