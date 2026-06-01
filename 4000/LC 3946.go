const N int = 1501

func maximumSaleItems(items [][]int, budget int) int {
    cnt := [N]int{}
    mn := math.MaxInt
    vis := [N]bool{}
    for _, v := range items {
        mn = min(v[1], mn)
        if vis[v[0]] {
            continue
        }
        vis[v[0]] = true
        for _, u := range items {
            if u[0] % v[0] == 0 {
                cnt[v[0]]++
            }
        }
    }

    f := make([]int, budget + 1)
    for _, v := range items {
        for i := budget; i >= v[1]; i-- {
            f[i] = max(f[i], f[i - v[1]] + cnt[v[0]])
        }
    }

    for i := mn; i <= budget; i++ {
        f[i] = max(f[i], f[i - mn] + 1)
    }

    return slices.Max(f)
}