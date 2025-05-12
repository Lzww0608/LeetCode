func cheapestJump(coins []int, maxJump int) (ans []int) {
    n := len(coins)
    if coins[n-1] == -1 {
        return nil
    }

    f := make([]int, n)
    f[n-1] = coins[n-1]
    for i := n - 2; i >= 0; i-- {
        f[i] = math.MaxInt32 
        if coins[i] == -1 {
            continue
        }
        for j := 1; j <= maxJump && i + j < n; j++ {
            f[i] = min(f[i], coins[i] + f[i+j])
        }
    }

    if f[0] == math.MaxInt32 {
        return nil
    }

    pre := 0 
    ans = append(ans, 1)
    for i := 1; i < n; i++ {
        if f[i] == f[pre] - coins[pre] {
            ans = append(ans, i + 1)
            pre = i 
        }
    }

    return 
}