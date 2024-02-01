func countSubstrings(s string) int {
    ans := 0
    n := len(s)
    var solve func(l, r int) 
    solve = func(l, r int) {
        for l >= 0 && r < n  && s[l] == s[r] {
            l--
            r++
            ans++
        }
    }
    for i := range s {
        solve(i,i)
        solve(i, i + 1)
    }
    return ans
}