func maximumSaleItems(items [][]int, budget int) int {
    n := len(items)
    cnt := make([]int, n + 1)
    mn, mx := math.MaxInt, 0 
    for _, v := range items {
        mn = min(v[1], mn)
        mx = max(v[0], mx)
        cnt[v[0]]++
    }

    i := 0
    ans := budget / mn 
    for _, v := range items {
        if v[1] < mn * 2 {
            items[i] = v
            i++
        }
    }

    sort.Slice(items[:i], func(p, q int) bool {
        return items[p][1] < items[q][1]
    })

    cost, sum := 0, 0
    for _, v := range items[:i] {
        factor, price := v[0], v[1]
        for x := factor; x <= mx; x += factor {
            if cnt[x] > 0 {
                t := min(cnt[x], (budget - cost) / price)
                if x == factor {
                    t = min(t, cnt[x] - 1)
                }
                if t == 0 {
                    if x != factor {
                        break
                    }
                    continue
                }
                cost += price * t
                sum += 2 * t
                cur := sum + (budget - cost) / mn
                ans = max(ans, cur)
            }
        }
    }

    return ans
}