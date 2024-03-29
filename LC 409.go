func longestPalindrome(s string) (ans int) {
    m := map[byte]int{}
    for i := range s {
        m[s[i]]++
    }
    mx := 0
    for _, x := range m {
        if x & 1 == 0 {
            ans += x
        } else {
            mx = max(mx, x)
            ans += x - 1
        }
    }
    if mx > 0 {
        ans++
    }

    return ans
}