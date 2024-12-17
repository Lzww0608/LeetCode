func minValidStrings(words []string, target string) (ans int) {
    n := len(target)
    f := make([]int, n)
    
    for _, word := range words {
        s := word + "#" + target
        z := Z(s) 
        m := len(word) + 1
        for i := 0; i < n; i++ {
            f[i] = max(f[i], z[i+m])
        }
    }

    cur, nxt := 0, 0
    for i, x := range f {
        nxt = max(nxt, i + x)
        if i == cur {
            if i == nxt {
                return -1
            }
            ans++
            cur = nxt
        }
    }

    return
}


func Z(s string) []int {
    n := len(s)
    z := make([]int, n)
    
    for i, l, r := 0, 0, 0; i < n; i++ {
        z[i] = max(min(z[i-l], r - i + 1), 0)
        for i + z[i] < n && s[z[i]] == s[i+z[i]] {
            l, r = i, i + z[i]
            z[i]++
        }
    }

    return z
}