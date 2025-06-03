func recoverArray(n int, sums []int) (ans []int) {
    sort.Ints(sums)
    for n > 1 {
        m := 1 << n 
        n--
        l, r, dif := 0, 1, sums[1] - sums[0] 
        f := false
        vis := make(map[int]bool)
        a := make([]int, 0, 1 << n)
        b := make([]int, 0, 1 << n)
        for {
            for l < m && vis[l] {
                l++
            }
            if l == m {
                break
            }
            vis[l] = true 
            for vis[r] || r < m && sums[r] != sums[l] + dif {
                r++
            }
            vis[r] = true 
            if sums[l] == 0 {
                f = true
            }
            a = append(a, sums[l])
            b = append(b, sums[r])
        }

        if f {
            ans = append(ans, dif)
            sums = a 
        } else {
            ans = append(ans, -dif)
            sums = b 
        }
    }
    ans = append(ans, sums[1] + sums[0])
    return
}