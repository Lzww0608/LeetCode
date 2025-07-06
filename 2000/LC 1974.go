const N int = 26
func minTimeToType(word string) int {
    ans := len(word)
    pre := 'a'
    for _, c := range word {
        t := int(int(c - pre) + N) % N
        ans += min(t, N - t)
        pre = c
    }

    return ans
}