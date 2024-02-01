func findAnagrams(s string, p string) []int {
    n := len(s)
    ans := []int{}
    alp := make([]int, 26)
    for _, c := range p {
        alp[c-'a']++
    }
    l, r := 0, 0
    for r < n {
        t := int(s[r] - 'a')
        alp[t]--
        for alp[t] < 0 {
            alp[s[l]-'a']++
            l++
        }
        if r - l + 1 == len(p) {
            ans = append(ans, l)
        }
        r++
    }
    return ans
}