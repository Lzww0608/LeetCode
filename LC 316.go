func removeDuplicateLetters(s string) string {
    alp := [26]int{}
    exists := map[int]bool{}

    for _, c := range s {
        x := int(c - 'a')
        alp[x]++
    }

    var ans string
    for i := range s {
        x := int(s[i] - 'a')
        if !exists[x] {
            for len(ans) > 0 && (ans[len(ans)-1] > s[i] && alp[int(ans[len(ans)-1] - 'a')] > 0) {
                exists[int(ans[len(ans)-1] - 'a')] = false
                ans = ans[:len(ans)-1]
            }
            ans += string(s[i])
            exists[x] = true
        }
        alp[x]--
    }

    return ans
}