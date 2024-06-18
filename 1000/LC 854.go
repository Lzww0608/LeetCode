func kSimilarity(s1 string, s2 string) int {
    var s, t []byte
    for i := range s1 {
        if s1[i] != s2[i] {
            s = append(s, s1[i])
            t = append(t, s2[i])
        }
    }

    n := len(s)
    if n == 0 {
        return 0
    }

    minSwap := func(i int) int {
        diff := 0
        for j := i; j < n; j++ {
            if s[j] != t[j] {
                diff++
            }
        }
        return (diff + 1) / 2
    }

    ans := n - 1
    var dfs func(int, int) 
    dfs = func(i, cost int) {
        if cost >= ans {
            return
        }
        for i < n && s[i] == t[i] {
            i++
        }
        if i == n {
            ans = min(ans, cost)
            return
        }
        if cost + minSwap(i) >= ans {
            return
        }
        
        for j := i + 1; j < n; j++ {
            if s[j] == t[i] {
                s[i], s[j] = s[j], s[i]
                dfs(i + 1, cost + 1)
                s[i], s[j] = s[j], s[i]
            }
        }
    }
    dfs(0, 0)

    return ans
}