func smallestSubsequence(s string) string {
    alp := [26]int{}
    exist := map[int]bool{}
    for i := range s {
        alp[int(s[i] - 'a')]++
    }

    var ans string
    for i := range s {
        x := int(s[i] - 'a')
        if !exist[x] {
            for len(ans) > 0 && ans[len(ans)-1] > s[i] && alp[int(ans[len(ans)-1] - 'a')] > 0 {
                exist[int(ans[len(ans)-1] - 'a')] = false
                ans = ans[:len(ans)-1]
            }
            ans += string(s[i])
            exist[x] = true
        }
        alp[x]--
    }

    return ans
}
