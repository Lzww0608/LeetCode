const N int = 26
func partitionLabels(s string) (ans []int) {
    pos := [N]int{}
    for i := range s {
        x := int(s[i] & 31) - 1
        pos[x] = max(pos[x], i)
    }

    n := len(s)
    pre := -1
    for l, r := 0, 0; l < n; l++ {
        x := int(s[l] & 31) - 1
        r = max(r, pos[x])

        if l == r {
            ans = append(ans, l - pre)
            pre = l
        }
    }

    return
}