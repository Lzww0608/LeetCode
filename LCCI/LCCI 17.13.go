const P int = 131313

func respace(dictionary []string, sentence string) int {
    n := len(sentence)
    h := make([]int, n+1)
    p := make([]int, n+1)
    
    p[0] = 1
    for i := 0; i < n; i++ {
        p[i+1] = p[i] * P
        h[i+1] = h[i] * P + int(sentence[i])
    }

    m := map[int]bool{}
    for _, word := range dictionary {
        wordHash := 0
        for _, c := range word {
            wordHash = wordHash * P + int(c)
        }
        m[wordHash] = true
    }

    dp := make([]int, n+1)
    for i := range dp {
        dp[i] = n 
    }
    dp[0] = 0

    for i := 1; i <= n; i++ {
        dp[i] = dp[i-1] + 1 

        for j := 0; j < i; j++ {
            currentHash := h[i] - h[j]*p[i-j]
            if m[currentHash] {
                dp[i] = min(dp[i], dp[j])
            }
        }
    }

    return dp[n]
}
