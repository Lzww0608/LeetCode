func countPalindromicSubsequence(s string) int {
    pre, suf := make([]int, 26), make([]int, 26)
    for i := range s {
        suf[int(s[i]-'a')]++
    }
    m := map[int]bool{}

    ans := 0
    for i := range s {
        suf[int(s[i]-'a')]--

        for j := 0; j < 26; j++ {
            if _, ok := m[int(s[i]-'a') * 26 + j]; pre[j] > 0 && suf[j] > 0 && !ok {
                ans++
                m[int(s[i]-'a') * 26 + j] = true
            }
        }
        //fmt.Println(ans)

        pre[int(s[i]-'a')]++
        
    }
    return ans
}