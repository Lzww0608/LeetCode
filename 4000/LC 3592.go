func findCoins(numWays []int) (ans []int) {
    n := len(numWays)
    cur := 0
    for i, x := range numWays {
        if x == 0 {
            continue
        } else if x != 1 {
            return nil
        } else {
            ans = append(ans, i + 1)
            cur = i + 1
            break
        }
    }

    pre := make([]int, n + 1)
    pre[0] = 1
    for {
        f := make([]int, n + 1)
        copy(f, pre)
        for i := 1; i <= n; i++ {
            if i - cur >= 0 {
                f[i] += f[i - cur]
            }
        }

        b := false
        for i := range n {
            if numWays[i] == f[i + 1] {
                continue
            } else if numWays[i] == f[i + 1] + 1 {
                cur = i + 1 
                ans = append(ans, i + 1)
                b = true
                break
            } else {
                return nil
            }
        }

        if !b {
            break
        }

        pre = f
    }
    
    return 
}