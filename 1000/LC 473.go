func makesquare(matchsticks []int) bool {
    n := len(matchsticks)
    sum := 0
    for _, x := range matchsticks {
        sum += x
    }

    if sum % 4 != 0 {
        return false
    }
    target := sum / 4
    sort.Ints(matchsticks)
    if matchsticks[n-1] > target {
        return false
    }

    f := make([]bool, 1 << n)
    for i := range f {
        f[i] = true
    }

    var dfs func(int, int) bool
    dfs = func(s, t int) bool {
        if s == 0 {
            return true
        }

        if !f[s] {
            return f[s]
        }

        f[s] = false
        for i := range matchsticks {
            if (s >> i) & 1 == 1 {
                if t + matchsticks[i] > target {
                    break
                }
                if dfs(s ^ (1 << i), (t + matchsticks[i]) % target) {
                    return true
                }
            }
        }

        return false
    }

    return dfs(1 << n - 1, 0)
}