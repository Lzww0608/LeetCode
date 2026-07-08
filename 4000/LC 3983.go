func canMakeSubsequence(s string, t string) bool {
    n, m := len(s), len(t)
    if n > m {
        return false
    }

    if n == 1 {
        return true
    }

    pre := make([]int, n + 1)
    suf := make([]int, n + 1)
    for i := range pre {
        pre[i] = m
        suf[i] = -1
    }
    for i, j := 0, 0; j < m && i < n; j++ {
        if t[j] == s[i] {
            pre[i] = j
            i++
        }
    }

    for i, j := n - 1, m - 1; j >= 0 && i >= 0; j-- {
        if t[j] == s[i] {
            suf[i] = j 
            i--
        }
    }
    //fmt.Println(pre, suf)
    if pre[n - 1] < m || suf[0] >= 0 {
        return true
    }
    
    for i := range n {
        if i == 0 {
            if suf[1] > 0 {
                return true
            }
            continue
        } 

        if i == n - 1 {
            if pre[n - 2] < m - 1 {
                return true
            }
            continue
        } 
            
        if pre[i - 1] < suf[i + 1] - 1  {   
            return true
        }
    }

    return false
}