func wordsTyping(s []string, rows int, cols int) (ans int) {
    n := len(s)
    f := make([]int, n)
    next := make([]int, n)
    for i := range n {
        ptr := i
        rem := cols 
        for rem >= len(s[ptr % n]) {
            rem -= len(s[ptr % n]) + 1
            ptr++
        }
        f[i] = ptr / n 
        next[i] = ptr % n
    }
    
    cur := 0
    for i := 0; i < rows; i++ {
        ans += f[cur]
        cur = next[cur]
    }

    return
}