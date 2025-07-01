func makePalindrome(s string) bool {
    n := len(s)
    cnt := 0
    for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1 {
        if s[i] != s[j] {
            cnt++
        }
    }
    
    return cnt <= 2
}