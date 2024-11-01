func maxPalindromes(s string, k int) int {
    n := len(s)
    p := make([]int, n)
    for i := range p {
        p[i] = -1
    }
    
    solve := func(i, j int) {
        for i >= 0 && j < n && s[i] == s[j] {
            if j - i + 1 >= k {
                p[j] = max(p[j], i)
            }
            i--
            j++
        }
    }

    for i := 0; i < n; i++ {
        solve(i, i)
        solve(i, i + 1)
    }


    f := make([]int, n + 1)
    for i := range s {
        if p[i] != -1 {
            f[i+1] = max(f[i], f[p[i]] + 1)
        } else {
            f[i+1] = f[i]
        }
    }
    //fmt.Println(f)

    return f[n]
}