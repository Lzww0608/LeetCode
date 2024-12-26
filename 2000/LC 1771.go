func longestPalindrome(word1 string, word2 string) (ans int) {
    s := word1 + word2
    n := len(s)
    
    f := make([][]int, n)
    for i := range f {
        f[i] = make([]int, n)
        f[i][i] = 1
    }
    for d := 2; d <= n; d++ {
        for i := 0; i + d <= n; i++ {
            j := i + d - 1
            if s[i] == s[j] {   
                f[i][j] = max(f[i][j-1], f[i+1][j], f[i+1][j-1] + 2)
                if i < len(word1) && j >= len(word1) {
                    ans = max(ans, f[i][j])
                }
                
            } else {
                f[i][j] = max(f[i+1][j], f[i][j-1])
            }
        }
    } 

    return 
}


