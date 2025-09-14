const N int = 26
func maxFreqSum(s string) int {
    cnt := [N]int{}
    a, b := 0, 0
    for i := range s {
        x := int(s[i] - 'a')
        cnt[x]++
        if s[i] == 'a' || s[i] == 'e' || s[i] == 'o' || s[i] == 'i' || s[i] == 'u' {
            a = max(a, cnt[x])
        } else {
            b = max(b, cnt[x])
        }
    }

    return a + b
}