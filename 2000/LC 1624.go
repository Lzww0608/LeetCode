const N int = 26
func maxLengthBetweenEqualCharacters(s string) (ans int) {
    pos := make([]int, N)
    ans = -1
    for i := range s {
        x := int(s[i] & 31) - 1
        if pos[x] != 0 {
            ans = max(ans, i - pos[x])
        } else {
            pos[x] = i + 1
        }
    }

    return
}