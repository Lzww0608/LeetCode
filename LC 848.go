func shiftingLetters(s string, shifts []int) string {
    n := len(s)
    ans := make([]rune, n)
    for i := n - 2; i >= 0; i-- {
        shifts[i] += shifts[i+1] % 26
    }
    for i, x := range shifts {
        ans[i] = 'a' + (rune(s[i])-'a' + rune(x)) % 26;
    }
    return string(ans)
}